package keeper

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"regexp"
	"time"
)

type DNSTypeRecord struct {
	RecordType   string   `json:"type"`
	RecordValues []string `json:"values"`
}

type DNSTransferOffer struct {
	Price                uint64 `json:"price"` // in udws
	ExpectedOwnerAddress string `json:"expected_owner_address"`
}

type DNSNameRecord struct {
	Issued                  time.Time         `json:"issued"`
	ValidTill               time.Time         `json:"valid_till"`
	TransferOffer           *DNSTransferOffer `json:"transfer_offer"`
	SubDomainsOnSale        bool              `json:"sub_domains_sale"`
	SubDomainsSalePrice     uint64            `json:"sub_domains_sale_price"`
	DomainProlongationPrice uint64            `json:"domain_prolongation_price"` // price to pay parent domain owner on prolongation
	Records                 []DNSTypeRecord   `json:"records"`
}

func ParseDomainData(data []byte) (DNSNameRecord, error) {
	resRecord := DNSNameRecord{}
	err := json.Unmarshal(data, &resRecord)
	if err != nil {
		return resRecord, errors.Wrap(err, "error parsing NFT DNS record")
	}
	return resRecord, nil
}

func validateARecord(values []string) error {
	ipRe := regexp.MustCompile(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}$`)
	for _, val := range values {
		valid := ipRe.MatchString(val)
		if !valid {
			return fmt.Errorf("record %s is invalid IP address", val)
		}
	}
	return nil
}

func (d *DNSNameRecord) validateRecords() error {
	for _, rec := range d.Records {
		if rec.RecordType == "A" {
			err := validateARecord(rec.RecordValues)
			if err != nil {
				return fmt.Errorf("record type A is invalid: %w", err)
			}
		}
	}
	return nil
}
