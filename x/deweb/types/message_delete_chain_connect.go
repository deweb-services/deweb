package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeleteChainConnect = "delete_chain_connect"

var _ sdk.Msg = &MsgDeleteChainConnect{}

func NewMsgDeleteChainConnect(creator string, chain string, address string) *MsgDeleteChainConnect {
	return &MsgDeleteChainConnect{
		Creator: creator,
		Chain:   chain,
		Address: address,
	}
}

func (msg *MsgDeleteChainConnect) Route() string {
	return RouterKey
}

func (msg *MsgDeleteChainConnect) Type() string {
	return TypeMsgDeleteChainConnect
}

func (msg *MsgDeleteChainConnect) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteChainConnect) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteChainConnect) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
