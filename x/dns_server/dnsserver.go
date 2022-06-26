package dns_server

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/deweb-services/deweb/x/nftmodule/keeper"
	"github.com/deweb-services/deweb/x/nftmodule/types"
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
	TypeA                  uint16 = 1  // a host address
	TypeMX                 uint16 = 15 // a host address
	ClassINET              uint16 = 1  // the Internet
	FlagResponse           uint16 = 1 << 15
	UDPMaxMessageSizeBytes uint   = 512 // RFC1035
)

var recordTypesMapping = map[uint16]string{
	TypeA:  "A",
	TypeMX: "MX",
}

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
func (srv *DNSResolverService) dbLookup(queryResourceRecord DNSResourceRecord) ([]DNSResourceRecord, []DNSResourceRecord, []DNSResourceRecord) {
	var answerResourceRecords = make([]DNSResourceRecord, 0)
	var authorityResourceRecords = make([]DNSResourceRecord, 0)
	var additionalResourceRecords = make([]DNSResourceRecord, 0)

	resolvedRecords, err := srv.ResolveDNSRecord(queryResourceRecord.DomainName, queryResourceRecord.Type)
	if err != nil {
		fmt.Printf("Cannot resolve record for %s (%d): %v", queryResourceRecord.DomainName, queryResourceRecord.Type, err)
		return answerResourceRecords, authorityResourceRecords, additionalResourceRecords
	}
	for _, rec := range resolvedRecords {
		ipAddr := net.ParseIP(rec)
		answerResourceRecords = append(answerResourceRecords, DNSResourceRecord{
			DomainName:         queryResourceRecord.DomainName,
			Type:               queryResourceRecord.Type,
			Class:              ClassINET,
			TimeToLive:         31337,
			ResourceData:       ipAddr[12:16], // ipv4 address
			ResourceDataLength: 4,
		})
	}

	return answerResourceRecords, authorityResourceRecords, additionalResourceRecords
}

func (srv *DNSResolverService) ResolveDNSRecord(domain string, recordType uint16) ([]string, error) {
	queryClient := types.NewQueryClient(srv.cliCtx)
	resp, err := queryClient.NFT(
		context.Background(),
		&types.QueryNFTRequest{
			DenomId: "domains",
			TokenId: domain,
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
		domainRecords, err := keeper.ParseNFTData([]byte(resp.NFT.Data))
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
	var authorityResourceRecords = make([]DNSResourceRecord, 0)
	var additionalResourceRecords = make([]DNSResourceRecord, 0)

	for _, queryResourceRecord := range queryResourceRecords {
		newAnswerRR, newAuthorityRR, newAdditionalRR := srv.dbLookup(queryResourceRecord)

		answerResourceRecords = append(answerResourceRecords, newAnswerRR...) // three dots cause the two lists to be concatenated
		authorityResourceRecords = append(authorityResourceRecords, newAuthorityRR...)
		additionalResourceRecords = append(additionalResourceRecords, newAdditionalRR...)
	}

	/**
	 * write response
	 */
	var responseBuffer = new(bytes.Buffer)
	var responseHeader DNSHeader

	responseHeader = DNSHeader{
		TransactionID:  queryHeader.TransactionID,
		Flags:          FlagResponse,
		NumQuestions:   queryHeader.NumQuestions,
		NumAnswers:     uint16(len(answerResourceRecords)),
		NumAuthorities: uint16(len(authorityResourceRecords)),
		NumAdditionals: uint16(len(additionalResourceRecords)),
	}

	err = Write(responseBuffer, &responseHeader)

	if err != nil {
		fmt.Println("Error writing to buffer: ", err.Error())
	}

	for _, queryResourceRecord := range queryResourceRecords {
		err = srv.writeDomainName(responseBuffer, queryResourceRecord.DomainName)

		if err != nil {
			fmt.Println("Error writing to buffer: ", err.Error())
		}

		Write(responseBuffer, queryResourceRecord.Type)
		Write(responseBuffer, queryResourceRecord.Class)
	}

	for _, answerResourceRecord := range answerResourceRecords {
		err = srv.writeDomainName(responseBuffer, answerResourceRecord.DomainName)

		if err != nil {
			fmt.Println("Error writing to buffer: ", err.Error())
		}

		Write(responseBuffer, answerResourceRecord.Type)
		Write(responseBuffer, answerResourceRecord.Class)
		Write(responseBuffer, answerResourceRecord.TimeToLive)
		Write(responseBuffer, answerResourceRecord.ResourceDataLength)
		Write(responseBuffer, answerResourceRecord.ResourceData)
	}

	for _, authorityResourceRecord := range authorityResourceRecords {
		err = srv.writeDomainName(responseBuffer, authorityResourceRecord.DomainName)

		if err != nil {
			fmt.Println("Error writing to buffer: ", err.Error())
		}

		Write(responseBuffer, authorityResourceRecord.Type)
		Write(responseBuffer, authorityResourceRecord.Class)
		Write(responseBuffer, authorityResourceRecord.TimeToLive)
		Write(responseBuffer, authorityResourceRecord.ResourceDataLength)
		Write(responseBuffer, authorityResourceRecord.ResourceData)
	}

	for _, additionalResourceRecord := range additionalResourceRecords {
		err = srv.writeDomainName(responseBuffer, additionalResourceRecord.DomainName)

		if err != nil {
			fmt.Println("Error writing to buffer: ", err.Error())
		}

		Write(responseBuffer, additionalResourceRecord.Type)
		Write(responseBuffer, additionalResourceRecord.Class)
		Write(responseBuffer, additionalResourceRecord.TimeToLive)
		Write(responseBuffer, additionalResourceRecord.ResourceDataLength)
		Write(responseBuffer, additionalResourceRecord.ResourceData)
	}

	serverConn.WriteToUDP(responseBuffer.Bytes(), clientAddr)
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

	fmt.Println("Listening at: ", serverAddr)

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