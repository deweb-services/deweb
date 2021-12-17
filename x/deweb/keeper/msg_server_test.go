package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/deweb-services/deweb/testutil/keeper"
	"github.com/deweb-services/deweb/x/deweb/keeper"
	"github.com/deweb-services/deweb/x/deweb/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DewebKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
