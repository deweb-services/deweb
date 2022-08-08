package keeper

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/types/errors"
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
