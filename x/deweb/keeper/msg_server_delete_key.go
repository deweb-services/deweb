package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/deweb-services/deweb/x/deweb/types"
)

func (k msgServer) DeleteKey(goCtx context.Context, msg *types.MsgDeleteKey) (*types.MsgDeleteKeyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.Logger(ctx).Error("[DeleteKey] Starting")

	userRecord, err := k.getUserKeyRecord(ctx, msg.Uuid)
	if err != nil {
		return nil, fmt.Errorf("error getting record by id: %w", err)
	}

	if userRecord.Creator != msg.Creator {
		return nil, fmt.Errorf("permission to perform this operation was denied")
	}

	userRecord.Deleted = true
	err = k.writeUserKeyRecord(ctx, userRecord, msg.Uuid)
	if err != nil {
		return nil, fmt.Errorf("error updating record by id: %w", err)
	}
	return &types.MsgDeleteKeyResponse{}, nil
}
