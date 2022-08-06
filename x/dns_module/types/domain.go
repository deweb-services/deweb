package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/deweb-services/deweb/x/dns_module/exported"
)

var _ exported.Domain = BaseDomain{}

// NewBaseNFT creates a new Domain instance
func NewBaseNFT(id string, owner sdk.AccAddress, data string) BaseDomain {
	return BaseDomain{
		Id:    id,
		Owner: owner.String(),
		Data:  data,
	}
}

// GetID return the id of BaseNFT
func (bnft BaseDomain) GetID() string {
	return bnft.Id
}

// GetOwner return the owner of BaseNFT
func (bnft BaseDomain) GetOwner() sdk.AccAddress {
	owner, _ := sdk.AccAddressFromBech32(bnft.Owner)
	return owner
}

// GetData return the Data of BaseNFT
func (bnft BaseDomain) GetData() string {
	return bnft.Data
}
