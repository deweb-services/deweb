package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSaveWallet{}, "deweb/SaveWallet", nil)
	cdc.RegisterConcrete(&MsgDeleteWallet{}, "deweb/DeleteWallet", nil)
	cdc.RegisterConcrete(&MsgConnectChain{}, "deweb/ConnectChain", nil)
	cdc.RegisterConcrete(&MsgDeleteChainConnect{}, "deweb/DeleteChainConnect", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSaveWallet{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeleteWallet{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgConnectChain{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeleteChainConnect{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
