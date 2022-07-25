package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/deweb-services/deweb/x/dns_module/exported"
	"github.com/deweb-services/deweb/x/dns_module/types"
)

// GetDomain gets the the specified Domain
func (k Keeper) GetDomain(ctx sdk.Context, tokenID string) (nft exported.Domain, err error) {
	store := ctx.KVStore(k.storeKey)
	recKey := types.KeyDomain(k.dnsDenomName, tokenID)
	bz := store.Get(recKey)
	if bz == nil {
		return nil, sdkerrors.Wrapf(types.ErrUnknownDomain, "not found Domain: %s", tokenID)
	}

	var baseDomain types.BaseDomain
	k.cdc.MustUnmarshal(bz, &baseDomain)

	return baseDomain, nil
}

// GetDomains returns all Domains by the specified denom ID
func (k Keeper) GetDomains(ctx sdk.Context, denom string) (nfts []exported.Domain) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.KeyDomain(denom, ""))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var baseDomain types.BaseDomain
		k.cdc.MustUnmarshal(iterator.Value(), &baseDomain)
		nfts = append(nfts, baseDomain)
	}

	return nfts
}

// Authorize checks if the sender is the owner of the given Domain
// Return the Domain if true, an error otherwise
func (k Keeper) Authorize(ctx sdk.Context, tokenID string, owner sdk.AccAddress) (types.BaseDomain, error) {
	domain, err := k.GetDomain(ctx, tokenID)
	if err != nil {
		return types.BaseDomain{}, err
	}

	if !owner.Equals(domain.GetOwner()) {
		return types.BaseDomain{}, sdkerrors.Wrap(types.ErrUnauthorized, owner.String())
	}

	return domain.(types.BaseDomain), nil
}

// HasDomain checks if the specified Domain exists
func (k Keeper) HasDomain(ctx sdk.Context, denomID, tokenID string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.KeyDomain(denomID, tokenID))
}

func (k Keeper) registerDomain(ctx sdk.Context, domain types.BaseDomain) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(&domain)
	recKey := types.KeyDomain(k.dnsDenomName, domain.GetID())
	store.Set(recKey, bz)
}

// deleteNFT deletes an existing Domain from store
func (k Keeper) deleteNFT(ctx sdk.Context, denomID string, nft exported.Domain) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyDomain(denomID, nft.GetID()))
}
