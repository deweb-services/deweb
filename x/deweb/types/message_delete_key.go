package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteWallet{}

func NewMsgDeleteKey(creator string, address string) *MsgDeleteWallet {
	return &MsgDeleteWallet{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgDeleteWallet) Route() string {
	return RouterKey
}

func (msg *MsgDeleteWallet) Type() string {
	return "DeleteWallet"
}

func (msg *MsgDeleteWallet) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteWallet) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteWallet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
