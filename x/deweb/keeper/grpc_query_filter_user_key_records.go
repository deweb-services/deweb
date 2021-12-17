package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/deweb-services/deweb/x/deweb/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FilterUserKeyRecords(goCtx context.Context, req *types.QueryFilterUserKeyRecordsRequest) (*types.QueryFilterUserKeyRecordsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	userRecordsIDs, err := k.getUserRecordsIDs(ctx, req.GetAddress())
	if err != nil {
		return nil, fmt.Errorf("error getting user records list: %w", err)
	}

	resultedRecords := make([]*types.QueryGetKeyRecordResponse, 0)
	for _, recID := range userRecordsIDs {
		record, err := k.getUserKeyRecord(ctx, recID)
		if err != nil {
			k.Logger(ctx).With("method", "FilterUserKeyRecords").Error(
				fmt.Sprintf("error getting record %s for address %s", record, req.GetAddress()))
		}
		if !req.Deleted && record.Deleted {
			continue
		}
		if req.Chain != "" && req.Chain != record.Chain {
			continue
		}
		resRecord := &types.QueryGetKeyRecordResponse{
			Uuid:    recID,
			Message: record.Message,
			Chain:   record.Chain,
			Deleted: record.Deleted,
		}
		resultedRecords = append(resultedRecords, resRecord)
	}

	return &types.QueryFilterUserKeyRecordsResponse{
		Records: resultedRecords,
	}, nil
}
