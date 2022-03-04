package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/deweb-services/deweb/x/deweb/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) getUserKeyRecord(ctx sdk.Context, recordID string) (userRecord types.UserWalletRec, err error) {
	store := ctx.KVStore(k.storeKey)
	postStore := prefix.NewStore(store, []byte(types.RecordsKey))
	postKey := []byte(recordID)
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
	records, err := k.readUserMappings(ctx, types.UsersRecords, address)

	return records, err
}

func (k Keeper) getChainMappingRecord(ctx sdk.Context, recordID string) (userRecord types.ChainAddressMapping, err error) {
	store := ctx.KVStore(k.storeKey)
	postStore := prefix.NewStore(store, []byte(types.ConnectChainRecords))
	postKey := []byte(recordID)
	value := postStore.Get(postKey)
	if value == nil {
		return userRecord, fmt.Errorf("record not found")
	}

	err = k.cdc.Unmarshal(value, &userRecord)
	return
}

func (k Keeper) getChainMappingRecordsIDs(ctx sdk.Context, address string) ([]string, error) {
	k.Logger(ctx).Error(fmt.Sprintf("[getChainMappingRecordsIDs] Process"))

	// Get the key-value module store using the store key (in our case store key is "chain")
	records, err := k.readUserMappings(ctx, types.UserConnectChainRecords, address)

	return records, err
}

// readUserMappings - retrieve address-to-ids mappings from storage
func (k Keeper) readUserMappings(ctx sdk.Context, storePrefix string, address string) ([]string, error) {
	// Get the key-value module store using the store key (in our case store key is "chain")
	store := ctx.KVStore(k.storeKey)
	// Get the part of the store that keeps posts (using post key, which is "Post-value-")
	storedRecordsMapStore := prefix.NewStore(store, []byte(storePrefix))
	recKey := []byte(address)
	value := storedRecordsMapStore.Get(recKey)
	if value == nil {
		return nil, nil
	}
	var storedRecordsIDs types.RecordsToUser
	if err := k.cdc.Unmarshal(value, &storedRecordsIDs); err != nil {
		return nil, err
	}
	return storedRecordsIDs.Records, nil
}
