package dns_server

import (
	"fmt"
	"github.com/miekg/dns"
)

func (srv *DNSResolverService) getResponse(requestMsg *dns.Msg) (*dns.Msg, error) {
	responseMsg := new(dns.Msg)
	if len(requestMsg.Question) > 0 {
		question := requestMsg.Question[0]
		_, ok := recordTypesMapping[question.Qtype]
		if !ok {
			fmt.Printf("unsupported type %d for domain %s \n", question.Qtype, question.Name)
			return responseMsg, nil
		}
		resAddresses, err := srv.resolveDNSRecord(question.Name, question.Qtype)
		if err == nil {
			storedVal := resAddresses[0]
			fmt.Printf("record for %s: %s found in chain\n", question.Name, storedVal.value)
			answer, err := dns.NewRR(fmt.Sprintf("%s %s %s", question.Name, storedVal.recType, storedVal.value))
			if err != nil {
				return responseMsg, err
			}
			responseMsg.Answer = append(responseMsg.Answer, answer)
		} else {
			answer, err := srv.processOtherTypes(srv.proxyDNSServer, &question, requestMsg)
			if err != nil {
				return responseMsg, err
			}
			responseMsg.Answer = append(responseMsg.Answer, *answer)
		}
	}

	return responseMsg, nil
}

func (srv *DNSResolverService) processOtherTypes(dnsServer string, q *dns.Question, requestMsg *dns.Msg) (*dns.RR, error) {
	queryMsg := new(dns.Msg)
	requestMsg.CopyTo(queryMsg)
	queryMsg.Question = []dns.Question{*q}

	msg, err := lookup(dnsServer, queryMsg)
	if err != nil {
		return nil, err
	}

	if len(msg.Answer) > 0 {
		return &msg.Answer[0], nil
	}
	return nil, fmt.Errorf("not found")
}

func lookup(server string, m *dns.Msg) (*dns.Msg, error) {
	dnsClient := new(dns.Client)
	dnsClient.Net = "udp"
	response, _, err := dnsClient.Exchange(m, server)
	if err != nil {
		return nil, err
	}

	return response, nil
}
