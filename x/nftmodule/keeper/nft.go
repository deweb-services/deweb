package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/deweb-services/deweb/x/nftmodule/exported"
	"github.com/deweb-services/deweb/x/nftmodule/types"
)

// GetNFT gets the the specified NFT
func (k Keeper) GetNFT(ctx sdk.Context, tokenID string) (nft exported.NFT, err error) {
	store := ctx.KVStore(k.storeKey)
	recKey := types.KeyNFT(k.dnsDenomName, tokenID)
	bz := store.Get(recKey)
	if bz == nil {
		return nil, sdkerrors.Wrapf(types.ErrUnknownNFT, "not found NFT: %s", tokenID)
	}

	var baseNFT types.BaseNFT
	k.cdc.MustUnmarshal(bz, &baseNFT)

	return baseNFT, nil
}

// GetNFTs returns all NFTs by the specified denom ID
func (k Keeper) GetNFTs(ctx sdk.Context, denom string) (nfts []exported.NFT) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.KeyNFT(denom, ""))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var baseNFT types.BaseNFT
		k.cdc.MustUnmarshal(iterator.Value(), &baseNFT)
		nfts = append(nfts, baseNFT)
	}

	return nfts
}

// Authorize checks if the sender is the owner of the given NFT
// Return the NFT if true, an error otherwise
func (k Keeper) Authorize(ctx sdk.Context, denomID, tokenID string, owner sdk.AccAddress) (types.BaseNFT, error) {
	nft, err := k.GetNFT(ctx, tokenID)
	if err != nil {
		return types.BaseNFT{}, err
	}

	if !owner.Equals(nft.GetOwner()) {
		return types.BaseNFT{}, sdkerrors.Wrap(types.ErrUnauthorized, owner.String())
	}

	return nft.(types.BaseNFT), nil
}

// HasNFT checks if the specified NFT exists
func (k Keeper) HasNFT(ctx sdk.Context, denomID, tokenID string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.KeyNFT(denomID, tokenID))
}

func (k Keeper) registerDomain(ctx sdk.Context, nft types.BaseNFT) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(&nft)
	recKey := types.KeyNFT(k.dnsDenomName, nft.GetID())
	store.Set(recKey, bz)
}

// deleteNFT deletes an existing NFT from store
func (k Keeper) deleteNFT(ctx sdk.Context, denomID string, nft exported.NFT) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyNFT(denomID, nft.GetID()))
}
