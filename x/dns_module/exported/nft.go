package exported

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Domain non fungible token interface
type Domain interface {
	GetID() string
	GetName() string
	GetOwner() sdk.AccAddress
	GetData() string
}
