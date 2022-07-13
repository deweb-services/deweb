package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrDomainAlreadyExists    = sdkerrors.Register(ModuleName, 5, "domain already exists")
	ErrUnknownDomain          = sdkerrors.Register(ModuleName, 6, "unknown domain")
	ErrUnauthorized           = sdkerrors.Register(ModuleName, 8, "unauthorized address")
	ErrInvalidDenom           = sdkerrors.Register(ModuleName, 9, "invalid denom")
	ErrInvalidTokenID         = sdkerrors.Register(ModuleName, 10, "invalid domain name")
	ErrDomainPermissionDenied = sdkerrors.Register(ModuleName, 12, "ownership check for domain failed")
)
