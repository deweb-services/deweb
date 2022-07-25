package dns_module

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/deweb-services/deweb/x/dns_module/keeper"
	"github.com/deweb-services/deweb/x/dns_module/types"
)

// InitGenesis stores the NFT genesis.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {
	if err := types.ValidateGenesis(data); err != nil {
		panic(err.Error())
	}

	k.SetParams(ctx, data.Params)

	for _, c := range data.Collections {
		if err := k.SetDenom(ctx, c.Denom); err != nil {
			panic(err)
		}
		if err := k.SetCollection(ctx, c); err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	params := k.GetParams(ctx)
	return types.NewGenesisState(k.GetCollections(ctx), params)
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() *types.GenesisState {
	dnsDenom := types.Denom{
		Id:     "domains",
		Name:   "dnsregistry",
		Symbol: "DwDNS",
	}
	return types.NewGenesisState([]types.Collection{{
		Denom: dnsDenom,
	}}, types.DefaultParams())
}
