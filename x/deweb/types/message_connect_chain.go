package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgConnectChain = "connect_chain"

var _ sdk.Msg = &MsgConnectChain{}

func NewMsgConnectChain(creator string, chain string, address string) *MsgConnectChain {
	return &MsgConnectChain{
		Creator: creator,
		Chain:   chain,
		Address: address,
	}
}

func (msg *MsgConnectChain) Route() string {
	return RouterKey
}

func (msg *MsgConnectChain) Type() string {
	return TypeMsgConnectChain
}

func (msg *MsgConnectChain) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgConnectChain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgConnectChain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
