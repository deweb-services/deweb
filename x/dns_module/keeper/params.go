package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/deweb-services/deweb/x/dns_module/types"
)

func (k Keeper) DomainPriceUDWS(ctx sdk.Context) (res uint64) {
	k.paramStore.Get(ctx, types.KeyDomainPrice, &res)
	return res
}

func (k Keeper) DomainExpirationHours(ctx sdk.Context) (res int64) {
	k.paramStore.Get(ctx, types.KeyDomainExpirationHours, &res)
	return res
}

func (k Keeper) DomainOwnerProlongationHours(ctx sdk.Context) (res int64) {
	k.paramStore.Get(ctx, types.KeyDomainOwnerProlongationHours, &res)
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
		k.DomainExpirationHours(ctx),
		k.DomainOwnerProlongationHours(ctx),
		k.BlockedTLDs(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramStore.SetParamSet(ctx, &params)
}
