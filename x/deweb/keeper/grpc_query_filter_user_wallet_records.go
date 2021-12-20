package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/deweb-services/deweb/x/deweb/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FilterUserWalletRecords(goCtx context.Context, req *types.QueryFilterUserWalletRecordsRequest) (*types.QueryFilterUserWalletRecordsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	userRecordsIDs, err := k.getUserRecordsIDs(ctx, req.GetOwner())
	if err != nil {
		return nil, fmt.Errorf("error getting user records list: %w", err)
	}

	resultedRecords := make([]*types.WalletRecordResponse, 0)
	processed := 0
	requestedOffset := int(req.Offset)
	requestedLimit := int(req.Limit)
	for _, recID := range userRecordsIDs {
		k.Logger(ctx).With("method", "FilterUserWalletRecords").Error(
			fmt.Sprintf("%s", recID))
		record, err := k.getUserKeyRecord(ctx, recID)
		if err != nil {
			k.Logger(ctx).With("method", "FilterUserWalletRecords").Error(
				fmt.Sprintf("error getting record %s for address %s", record, req.GetAddress()))
		}
		if req.Address != "" && req.Address != record.Address {
			continue
		}
		if !req.Deleted && record.Deleted {
			continue
		}
		if req.Chain != "" && req.Chain != record.Chain {
			continue
		}
		processed += 1
		if processed <= requestedOffset {
			continue
		}
		resRecord := &types.WalletRecordResponse{
			Owner:        req.Owner,
			Address:      record.Address,
			EncryptedKey: record.EncryptedKey,
			Chain:        record.Chain,
			Deleted:      record.Deleted,
		}
		resultedRecords = append(resultedRecords, resRecord)
		if len(resultedRecords) == requestedLimit {
			break
		}
	}

	return &types.QueryFilterUserWalletRecordsResponse{
		Records: resultedRecords,
	}, nil
}
