package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/deweb-services/deweb/x/nftmodule/types"
)

func (k Keeper) DomainPriceUDWS(ctx sdk.Context) (res uint64) {
	k.paramStore.Get(ctx, types.KeyDomainPrice, &res)
	return res
}

func (k Keeper) DomainExpirationMinutes(ctx sdk.Context) (res int64) {
	k.paramStore.Get(ctx, types.KeyDomainExpirationMinutes, &res)
	return res
}

func (k Keeper) DomainOwnerProlongationMinutes(ctx sdk.Context) (res int64) {
	k.paramStore.Get(ctx, types.KeyDomainOwnerProlongationMinutes, &res)
	return res
}

func (k Keeper) BlockedTLDs(ctx sdk.Context) (res []string) {
	k.paramStore.Get(ctx, types.KeyBlockedTLDs, &res)
	return res
}

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.DomainPriceUDWS(ctx),
		k.DomainExpirationMinutes(ctx),
		k.DomainOwnerProlongationMinutes(ctx),
		k.BlockedTLDs(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramStore.SetParamSet(ctx, &params)
}
