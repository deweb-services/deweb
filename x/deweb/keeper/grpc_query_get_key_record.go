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

func (k Keeper) GetKeyRecord(goCtx context.Context, req *types.QueryGetKeyRecordRequest) (*types.QueryGetKeyRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	k.Logger(ctx).Error(fmt.Sprintf("[GetKeyRecord] Process read"))

	userRecord, err := k.getUserKeyRecord(ctx, req.Uuid)
	if err != nil {
		return nil, fmt.Errorf("error getting record by id: %w", err)
	}
	return &types.QueryGetKeyRecordResponse{
		Uuid:    req.Uuid,
		Message: userRecord.Message,
		Chain:   userRecord.Chain,
		Deleted: userRecord.Deleted,
	}, nil
}

func (k Keeper) getUserKeyRecord(ctx sdk.Context, uuid string) (userRecord types.UserKeyRec, err error) {
	store := ctx.KVStore(k.storeKey)
	postStore := prefix.NewStore(store, []byte(types.RecordsKey))
	postKey := []byte(uuid)
	value := postStore.Get(postKey)
	if value == nil {
		return userRecord, fmt.Errorf("record not found")
	}

	err = k.cdc.Unmarshal(value, &userRecord)
	return
}
