package dns_server

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/deweb-services/deweb/x/dns_module/keeper"
	"github.com/deweb-services/deweb/x/dns_module/types"
	"golang.org/x/net/dns/dnsmessage"
	"net"
	"os"
	"strings"
	"time"
)

// DNSHeader describes the request/response DNS header
type DNSHeader struct {
	TransactionID  uint16
	Flags          uint16
	NumQuestions   uint16
	NumAnswers     uint16
	NumAuthorities uint16
	NumAdditionals uint16
}

// DNSResourceRecord describes individual records in the request and response of the DNS payload body
type DNSResourceRecord struct {
	DomainName         string
	Type               uint16
	Class              uint16
	TimeToLive         uint32
	ResourceDataLength uint16
	ResourceData       []byte
}

// Type and Class values for DNSResourceRecord
const (
	TypeA                  uint16 = 1
	TypeMX                 uint16 = 15
	ClassINET              uint16 = 1   // the Internet
	UDPMaxMessageSizeBytes uint   = 512 // RFC1035
)

type CacheRecord struct {
	Values     []keeper.DNSTypeRecord
	CreateTime time.Time
}

type DNSResolverService struct {
	cliCtx        client.Context
	cachedRecords map[string]CacheRecord
}

func NewDNSResolverService(cliCtx client.Context) *DNSResolverService {
	return &DNSResolverService{
		cliCtx:        cliCtx,
		cachedRecords: make(map[string]CacheRecord),
	}
}

// Pretend to look up values in a database
func (srv *DNSResolverService) dbLookup(queryResourceRecord DNSResourceRecord) []DNSResourceRecord {
	var answerResourceRecords = make([]DNSResourceRecord, 0)

	resolvedRecords, err := srv.resolveDNSRecord(queryResourceRecord.DomainName, queryResourceRecord.Type)
	if err != nil {
		fmt.Printf("Cannot resolve record for %s (%d): %v", queryResourceRecord.DomainName, queryResourceRecord.Type, err)
		return answerResourceRecords
	}
	for _, rec := range resolvedRecords {
		var resData []byte
		if queryResourceRecord.Type == 1 {
			ipAddr := net.ParseIP(rec)
			resData = ipAddr[12:16]
		} else {
			resData = []byte(rec)
		}

		answerResourceRecords = append(answerResourceRecords, DNSResourceRecord{
			DomainName:         queryResourceRecord.DomainName,
			Type:               queryResourceRecord.Type,
			Class:              ClassINET,
			TimeToLive:         31337,
			ResourceData:       resData, // ipv4 address
			ResourceDataLength: uint16(len(resData)),
		})
	}

	return answerResourceRecords
}

func (srv *DNSResolverService) resolveDNSRecord(domain string, recordType uint16) ([]string, error) {
	queryClient := types.NewQueryClient(srv.cliCtx)
	resp, err := queryClient.Domain(
		context.Background(),
		&types.QueryDomainRequest{
			DomainName: domain,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("cannot perform NFT search: %w", err)
	}

	var storedRecords []keeper.DNSTypeRecord
	recFromCache, ok := srv.cachedRecords[domain]
	if ok {
		cacheExpireTime := recFromCache.CreateTime.Add(10 * time.Second)
		if cacheExpireTime.After(time.Now()) {
			storedRecords = recFromCache.Values
		} else {
			delete(srv.cachedRecords, domain)
		}
	}
	if storedRecords == nil {
		domainRecords, err := keeper.ParseDomainData([]byte(resp.Domain.Data))
		if err != nil {
			return nil, fmt.Errorf("cannot parse stored record for domain %s: %w", domain, err)
		}
		storedRecords = domainRecords.Records
		srv.cachedRecords[domain] = CacheRecord{
			Values:     domainRecords.Records,
			CreateTime: time.Now(),
		}
	}

	recordTypeName, ok := recordTypesMapping[recordType]
	if !ok {
		return nil, fmt.Errorf("unsupported type %d", recordType)
	}
	for _, rec := range storedRecords {
		if rec.RecordType == recordTypeName {
			return rec.RecordValues, nil
		}
	}
	return nil, fmt.Errorf("record for %s not found", domain)
}

// RFC1035: "Domain names in messages are expressed in terms of a sequence
// of labels. Each label is represented as a one octet length field followed
// by that number of octets.  Since every domain name ends with the null label
// of the root, a domain name is terminated by a length byte of zero."
func (srv *DNSResolverService) readDomainName(requestBuffer *bytes.Buffer) (string, error) {
	var domainName string

	b, err := requestBuffer.ReadByte()

	for ; b != 0 && err == nil; b, err = requestBuffer.ReadByte() {
		labelLength := int(b)
		labelBytes := requestBuffer.Next(labelLength)
		labelName := string(labelBytes)

		if len(domainName) == 0 {
			domainName = labelName
		} else {
			domainName += "." + labelName
		}
	}

	return domainName, err
}

// RFC1035: "Domain names in messages are expressed in terms of a sequence
// of labels. Each label is represented as a one octet length field followed
// by that number of octets.  Since every domain name ends with the null label
// of the root, a domain name is terminated by a length byte of zero."
func (srv *DNSResolverService) writeDomainName(responseBuffer *bytes.Buffer, domainName string) error {
	labels := strings.Split(domainName, ".")

	for _, label := range labels {
		labelLength := len(label)
		labelBytes := []byte(label)

		responseBuffer.WriteByte(byte(labelLength))
		responseBuffer.Write(labelBytes)
	}

	err := responseBuffer.WriteByte(byte(0))

	return err
}

func (srv *DNSResolverService) prepareDNSAnswer(reqID uint16, dnsQuestions []DNSResourceRecord, dnsAnswers []DNSResourceRecord) ([]byte, error) {
	msg := dnsmessage.Message{
		Header: dnsmessage.Header{
			Response:      true,
			Authoritative: true,
			ID:            reqID,
		},
		Questions: make([]dnsmessage.Question, 0, len(dnsQuestions)),
		Answers:   make([]dnsmessage.Resource, 0, len(dnsAnswers)),
	}
	for _, qRec := range dnsQuestions {
		domainName := qRec.DomainName + "."
		name, err := dnsmessage.NewName(domainName)
		if err != nil {
			continue
		}
		qMessage := dnsmessage.Question{
			Name:  name,
			Type:  dnsmessage.Type(qRec.Type),
			Class: dnsmessage.ClassINET,
		}
		msg.Questions = append(msg.Questions, qMessage)
	}

	for _, answer := range dnsAnswers {
		domainName := answer.DomainName + "."
		name, err := dnsmessage.NewName(domainName)
		if err != nil {
			continue
		}
		var body dnsmessage.ResourceBody
		switch answer.Type {
		case TypeA:
			if len(answer.ResourceData) != 4 {
				fmt.Printf("Invalid response with length %d: %v\n", len(answer.ResourceData), answer.ResourceData)
			}
			var resContent [4]byte
			for i := 0; i < 4; i++ {
				resContent[i] = answer.ResourceData[i]
			}
			body = &dnsmessage.AResource{A: resContent}
		case TypeMX:
			resName, err := dnsmessage.NewName(string(answer.ResourceData))
			if err != nil {
				fmt.Printf("cannot process MX response for %s. result=%s: %v",
					answer.DomainName, string(answer.ResourceData), err)
			}
			body = &dnsmessage.MXResource{
				Pref: 0,
				MX:   resName,
			}
		}

		answerRec := dnsmessage.Resource{
			Header: dnsmessage.ResourceHeader{
				Name:  name,
				Type:  dnsmessage.Type(answer.Type),
				Class: dnsmessage.ClassINET,
			},
			Body: body,
		}
		msg.Answers = append(msg.Answers, answerRec)
	}

	buf, err := msg.Pack()
	if err != nil {
		return nil, fmt.Errorf("cannot pack message: %w", err)
	}
	return buf, nil
}

func (srv *DNSResolverService) handleDNSClient(requestBytes []byte, serverConn *net.UDPConn, clientAddr *net.UDPAddr) {
	/**
	 * read request
	 */
	var requestBuffer = bytes.NewBuffer(requestBytes)
	var queryHeader DNSHeader
	var queryResourceRecords []DNSResourceRecord

	err := binary.Read(requestBuffer, binary.BigEndian, &queryHeader) // network byte order is big endian

	if err != nil {
		fmt.Println("Error decoding header: ", err.Error())
	}

	queryResourceRecords = make([]DNSResourceRecord, queryHeader.NumQuestions)

	for idx, _ := range queryResourceRecords {
		queryResourceRecords[idx].DomainName, err = srv.readDomainName(requestBuffer)

		if err != nil {
			fmt.Println("Error decoding label: ", err.Error())
		}

		queryResourceRecords[idx].Type = binary.BigEndian.Uint16(requestBuffer.Next(2))
		queryResourceRecords[idx].Class = binary.BigEndian.Uint16(requestBuffer.Next(2))
	}

	/**
	 * lookup values
	 */
	var answerResourceRecords = make([]DNSResourceRecord, 0)

	for _, queryResourceRecord := range queryResourceRecords {
		newAnswerRR := srv.dbLookup(queryResourceRecord)

		answerResourceRecords = append(answerResourceRecords, newAnswerRR...) // three dots cause the two lists to be concatenated
	}

	/**
	 * write response
	 */

	resBytes, err := srv.prepareDNSAnswer(queryHeader.TransactionID, queryResourceRecords, answerResourceRecords)
	if err != nil {
		fmt.Printf("Error making response for %v: %v", queryResourceRecords, err)
		resBytes = make([]byte, 0)
	}
	_, _ = serverConn.WriteToUDP(resBytes, clientAddr)
}

func (srv *DNSResolverService) RunServer(port int) {
	serverAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))

	if err != nil {
		fmt.Println("Error resolving UDP address: ", err.Error())
		os.Exit(1)
	}

	serverConn, err := net.ListenUDP("udp", serverAddr)

	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		os.Exit(1)
	}

	fmt.Println("DNS Listening at: ", serverAddr)

	defer serverConn.Close()

	for {
		requestBytes := make([]byte, UDPMaxMessageSizeBytes)

		_, clientAddr, err := serverConn.ReadFromUDP(requestBytes)

		if err != nil {
			fmt.Println("Error receiving: ", err.Error())
		} else {
			fmt.Println("Received request from ", clientAddr)
			go srv.handleDNSClient(requestBytes, serverConn, clientAddr) // array is value type (call-by-value), i.e. copied
		}
	}
}
