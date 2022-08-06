package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/deweb-services/deweb/x/dns_module/types"
)

// SetCollection saves all Domai and returns an error if there already exists
func (k Keeper) SetCollection(ctx sdk.Context, collection types.Collection) error {
	for _, nft := range collection.NFTs {
		if err := k.RegisterDomain(ctx, nft.GetID(), nft.GetData(), nft.GetOwner(), nil); err != nil {
			return err
		}
	}
	return nil
}

// GetCollection returns the collection by the specified denom ID
func (k Keeper) GetCollection(ctx sdk.Context, denomID string) (types.Collection, error) {
	denom, found := k.GetDenom(ctx, denomID)
	if !found {
		return types.Collection{}, sdkerrors.Wrapf(types.ErrInvalidDenom, "denomID %s not existed ", denomID)
	}

	nfts := k.GetDomains(ctx, denomID)
	return types.NewCollection(denom, nfts), nil
}

// GetCollections returns all the collections
func (k Keeper) GetCollections(ctx sdk.Context) (cs []types.Collection) {
	for _, denom := range k.GetDenoms(ctx) {
		nfts := k.GetDomains(ctx, denom.Id)
		cs = append(cs, types.NewCollection(denom, nfts))
	}
	return cs
}

// GetDenomSupply returns the number of Domains by the specified denom ID
func (k Keeper) GetTotalSupply(ctx sdk.Context, denomID string) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyCollection(denomID))
	if len(bz) == 0 {
		return 0
	}
	return types.MustUnMarshalSupply(k.cdc, bz)
}

func (k Keeper) increaseSupply(ctx sdk.Context, denomID string) {
	supply := k.GetTotalSupply(ctx, denomID)
	supply++

	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalSupply(k.cdc, supply)
	store.Set(types.KeyCollection(denomID), bz)
}

func (k Keeper) decreaseSupply(ctx sdk.Context, denomID string) {
	supply := k.GetTotalSupply(ctx, denomID)
	supply--

	store := ctx.KVStore(k.storeKey)
	if supply == 0 {
		store.Delete(types.KeyCollection(denomID))
		return
	}

	bz := types.MustMarshalSupply(k.cdc, supply)
	store.Set(types.KeyCollection(denomID), bz)
}
