package dns_server

import (
	"context"
	"fmt"
	"github.com/deweb-services/deweb/x/dns_module/types"
	"strings"
	"time"
)

type resolvedRecord struct {
	recType string
	value   string
}

func (srv *DNSResolverService) resolveDNSRecord(domain string, recordType uint16) ([]resolvedRecord, error) {
	if strings.HasSuffix(domain, ".") {
		domain = domain[:len(domain)-1]
	}
	queryClient := types.NewQueryClient(srv.cliCtx)
	resp, err := queryClient.Domain(
		context.Background(),
		&types.QueryDomainRequest{
			DomainName: domain,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("cannot perform domain search: %w", err)
	}

	var storedRecords []*types.DNSRecords
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
		srv.cachedRecords[domain] = CacheRecord{
			Values:     resp.Domain.Records,
			CreateTime: time.Now(),
		}
		storedRecords = resp.Domain.Records
	}

	recordTypeName, ok := recordTypesMapping[recordType]
	if !ok {
		return nil, fmt.Errorf("unsupported type %d", recordType)
	}
	var cnameRecords []string
	var resRecords []string
	for _, rec := range storedRecords {
		if rec.Type == recordTypeName {
			resRecords = rec.Values
			break
		}
		if rec.Type == "CNAME" {
			cnameRecords = rec.Values
		}
	}
	result := make([]resolvedRecord, 0)
	if resRecords != nil {
		for _, resRec := range resRecords {
			respItem := resolvedRecord{
				recType: recordTypeName,
				value:   resRec,
			}
			result = append(result, respItem)
		}
		return result, nil
	}
	if cnameRecords != nil {
		for _, resRec := range cnameRecords {
			respItem := resolvedRecord{
				recType: "CNAME",
				value:   resRec,
			}
			result = append(result, respItem)
		}
		return result, nil
	}
	return nil, fmt.Errorf("record for %s not found", domain)
}
