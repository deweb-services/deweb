package dns_server

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/deweb-services/deweb/x/dns_module/types"
	"github.com/miekg/dns"
	"net"
	"os"
	"regexp"
	"strings"
	"time"
)

type CacheRecord struct {
	Values     []*types.DNSRecords
	CreateTime time.Time
}

type DNSResolverService struct {
	cliCtx         client.Context
	proxyDNSServer string
	cachedRecords  map[string]CacheRecord
}

func NewDNSResolverService(cliCtx client.Context, proxyServer string) *DNSResolverService {
	if len(proxyServer) == 0 {
		proxyServer = "1.1.1.1:53"
	}
	return &DNSResolverService{
		cliCtx:         cliCtx,
		proxyDNSServer: proxyServer,
		cachedRecords:  make(map[string]CacheRecord),
	}
}

func (srv *DNSResolverService) RunServer(port int) {
	serverAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))

	if err != nil {
		fmt.Println("Error resolving UDP address: ", err.Error())
		os.Exit(1)
	}

	dns.HandleFunc(".", func(w dns.ResponseWriter, r *dns.Msg) {
		switch r.Opcode {
		case dns.OpcodeQuery:
			qStrings := make([]string, 0, len(r.Question))
			for _, q := range r.Question {
				qStr := strings.ReplaceAll(q.String(), "\t", " | ")
				qStrings = append(qStrings, qStr)
			}

			m, err := srv.getResponse(r)
			if err != nil {
				logStr := fmt.Sprintf("Failed lookup for [%s] with error: %s\n", strings.Join(qStrings, ";"), err.Error())
				fmt.Printf(logStr)
				m.SetReply(r)
				w.WriteMsg(m)
				return
			}
			ipPattern := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)
			for _, answer := range m.Answer {
				ipAddress := ipPattern.FindAllString(m.Answer[0].String(), -1)
				if len(ipAddress) > 0 {
					fmt.Printf("Lookup for [%s] answer for %s with ip %s\n", strings.Join(qStrings, ";"), answer.Header().Name, ipAddress[0])
				} else {
					fmt.Printf("Lookup for [%s] answer %s with response %s\n", strings.Join(qStrings, ";"), answer.Header().Name, answer.String())
				}
			}
			m.SetReply(r)
			w.WriteMsg(m)
		}
	})

	server := &dns.Server{Addr: serverAddr.String(), Net: "udp"}
	fmt.Printf("Starting at %s\n", serverAddr.String())
	err = server.ListenAndServe()
	if err != nil {
		fmt.Printf("Failed to start server: %s\n ", err.Error())
	}

}
