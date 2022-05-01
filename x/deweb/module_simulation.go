package deweb

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/deweb-services/deweb/testutil/sample"
	dewebsimulation "github.com/deweb-services/deweb/x/deweb/simulation"
	"github.com/deweb-services/deweb/x/deweb/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = dewebsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgConnectChain = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgConnectChain int = 100

	opWeightMsgDeleteChainConnect = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteChainConnect int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	dewebGenesis := types.GenesisState{
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&dewebGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgConnectChain int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgConnectChain, &weightMsgConnectChain, nil,
		func(_ *rand.Rand) {
			weightMsgConnectChain = defaultWeightMsgConnectChain
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgConnectChain,
		dewebsimulation.SimulateMsgConnectChain(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteChainConnect int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteChainConnect, &weightMsgDeleteChainConnect, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteChainConnect = defaultWeightMsgDeleteChainConnect
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteChainConnect,
		dewebsimulation.SimulateMsgDeleteChainConnect(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
