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
	Issued        time.Time         `json:"issued"`
	ValidTill     time.Time         `json:"valid_till"`
	TransferOffer *DNSTransferOffer `json:"transfer_offer"`
	Records       []DNSTypeRecord   `json:"records"`
}

func ParseNFTData(data []byte) (DNSNameRecord, error) {
	resRecord := DNSNameRecord{}
	err := json.Unmarshal(data, &resRecord)
	if err != nil {
		return resRecord, errors.Wrap(err, "error parsing NFT DNS record")
	}
	return resRecord, nil
}
