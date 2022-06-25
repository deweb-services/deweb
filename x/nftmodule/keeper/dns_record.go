package keeper

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

type DNSTypeRecord struct {
	RecordType   string   `json:"type"`
	RecordValues []string `json:"values"`
}

type DNSNameRecord struct {
	Records []DNSTypeRecord `json:"records"`
}

func ParseNFTData(data []byte) (DNSNameRecord, error) {
	resRecord := DNSNameRecord{}
	err := json.Unmarshal(data, &resRecord)
	if err != nil {
		return resRecord, errors.Wrap(err, "error parsing NFT DNS record")
	}
	return resRecord, nil
}
