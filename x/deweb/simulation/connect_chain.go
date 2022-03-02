package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/deweb-services/deweb/x/deweb/keeper"
	"github.com/deweb-services/deweb/x/deweb/types"
)

func SimulateMsgConnectChain(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgConnectChain{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the ConnectChain simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "ConnectChain simulation not implemented"), nil, nil
	}
}
