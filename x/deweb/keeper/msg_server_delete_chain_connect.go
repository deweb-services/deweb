package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/deweb-services/deweb/x/deweb/types"
)

func (k msgServer) DeleteChainConnect(goCtx context.Context, msg *types.MsgDeleteChainConnect) (*types.MsgDeleteChainConnectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.Logger(ctx).Error("[DeleteChainConnect] Starting")

	recordID := msg.Creator + "_" + msg.Chain + "_" + msg.Address
	chainRecord, err := k.getChainMappingRecord(ctx, recordID)
	if err != nil {
		return nil, fmt.Errorf("error getting record by id: %w", err)
	}

	chainRecord.Deleted = true
	err = k.writeChainMappingRecord(ctx, chainRecord, recordID)
	if err != nil {
		return nil, fmt.Errorf("error updating record by id: %w", err)
	}

	return &types.MsgDeleteChainConnectResponse{}, nil
}
