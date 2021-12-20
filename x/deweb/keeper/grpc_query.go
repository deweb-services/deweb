package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/deweb-services/deweb/x/deweb/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) getUserKeyRecord(ctx sdk.Context, address string) (userRecord types.UserWalletRec, err error) {
	store := ctx.KVStore(k.storeKey)
	postStore := prefix.NewStore(store, []byte(types.RecordsKey))
	postKey := []byte(address)
	value := postStore.Get(postKey)
	if value == nil {
		return userRecord, fmt.Errorf("record not found")
	}

	err = k.cdc.Unmarshal(value, &userRecord)
	return
}

func (k Keeper) getUserRecordsIDs(ctx sdk.Context, address string) ([]string, error) {
	k.Logger(ctx).Error(fmt.Sprintf("[getUserRecordsIDs] Process"))

	// Get the key-value module store using the store key (in our case store key is "chain")
	store := ctx.KVStore(k.storeKey)
	// Get the part of the store that keeps posts (using post key, which is "Post-value-")
	userRecordsMapStore := prefix.NewStore(store, []byte(types.UsersRecords))
	recKey := []byte(address)
	value := userRecordsMapStore.Get(recKey)
	if value == nil {
		return nil, nil
	}
	var userRecordsIDs types.RecordsToUser
	if err := k.cdc.Unmarshal(value, &userRecordsIDs); err != nil {
		return nil, err
	}

	return userRecordsIDs.Records, nil
}
