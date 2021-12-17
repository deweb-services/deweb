package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSaveUser{}

func NewMsgSaveUser(creator string, message string, chain string) *MsgSaveUser {
	return &MsgSaveUser{
		Creator: creator,
		Message: message,
		Chain:   chain,
	}
}

func (msg *MsgSaveUser) Route() string {
	return RouterKey
}

func (msg *MsgSaveUser) Type() string {
	return "SaveUser"
}

func (msg *MsgSaveUser) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSaveUser) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSaveUser) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
