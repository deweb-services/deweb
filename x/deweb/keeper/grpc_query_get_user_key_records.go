package keeper

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/deweb-services/deweb/x/deweb/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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

func (k Keeper) GetUserKeyRecords(goCtx context.Context, req *types.QueryGetUserKeyRecordsRequest) (*types.QueryGetUserKeyRecordsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	k.Logger(ctx).Error(fmt.Sprintf("[GetUserKeyRecords] Process read"))

	userRecordsIDs, err := k.getUserRecordsIDs(ctx, req.Address)
	if err != nil {
		return nil, err
	}

	return &types.QueryGetUserKeyRecordsResponse{
		Uuids: userRecordsIDs,
	}, nil
}
