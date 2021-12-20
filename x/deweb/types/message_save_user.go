package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSaveWallet{}

func NewMsgSaveUser(creator string, address string, encKey string, chain string) *MsgSaveWallet {
	return &MsgSaveWallet{
		Creator:      creator,
		Address:      address,
		EncryptedKey: encKey,
		Chain:        chain,
	}
}

func (msg *MsgSaveWallet) Route() string {
	return RouterKey
}

func (msg *MsgSaveWallet) Type() string {
	return "SaveWallet"
}

func (msg *MsgSaveWallet) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSaveWallet) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSaveWallet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
