package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/deweb-services/deweb/x/dns_module/types"
)

// GetByOwner gets all the domains owned by an address
func (k Keeper) GetByOwner(ctx sdk.Context, address sdk.AccAddress) types.Owner {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyOwner(address, ""))
	defer iterator.Close()

	owner := types.Owner{
		Address:       address.String(),
		IDCollections: types.IDCollections{},
	}
	domainsCollections := types.IDCollections{}
	resDomains := make([]string, 0)

	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		_, domainName, _ := types.SplitKeyOwner(key)
		resDomains = append(resDomains, domainName)
		domainsCollections = domainsCollections.Add(k.dnsDenomName, domainName)
	}

	owner.IDCollections = domainsCollections

	return owner
}

func (k Keeper) deleteOwner(ctx sdk.Context, tokenID string, owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyOwner(owner, tokenID))
}

func (k Keeper) setOwner(ctx sdk.Context, tokenID string, owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)

	bz := types.MustMarshalTokenID(k.cdc, tokenID)
	store.Set(types.KeyOwner(owner, tokenID), bz)
}

func (k Keeper) swapOwner(ctx sdk.Context, tokenID string, srcOwner, dstOwner sdk.AccAddress) {

	// delete old owner key
	k.deleteOwner(ctx, tokenID, srcOwner)

	// set new owner key
	k.setOwner(ctx, tokenID, dstOwner)
}
