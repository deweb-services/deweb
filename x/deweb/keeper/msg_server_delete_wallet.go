package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/deweb-services/deweb/x/deweb/types"
)

func (k msgServer) DeleteWallet(goCtx context.Context, msg *types.MsgDeleteWallet) (*types.MsgDeleteWalletResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.Logger(ctx).Error("[DeleteWallet] Starting")

	recordID := msg.Creator + "_" + msg.Address
	userRecord, err := k.getUserKeyRecord(ctx, recordID)
	if err != nil {
		return nil, fmt.Errorf("error getting record by id: %w", err)
	}

	userRecord.Deleted = true
	err = k.writeUserKeyRecord(ctx, userRecord, recordID)
	if err != nil {
		return nil, fmt.Errorf("error updating record by id: %w", err)
	}
	return &types.MsgDeleteWalletResponse{}, nil
}
