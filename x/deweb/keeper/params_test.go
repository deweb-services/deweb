package keeper_test

import (
	"testing"

	testkeeper "github.com/deweb-services/deweb/testutil/keeper"
	"github.com/deweb-services/deweb/x/deweb/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DewebKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
