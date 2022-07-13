package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// constant used to indicate that some field should not be updated
const (
	TypeMsgTransferDomain = "transfer_domain"
	TypeMsgEditDomain     = "edit_domain"
	TypeMsgRegisterDomain = "register_domain"
	TypeMsgRemoveDomain   = "remove_domain"
)

var (
	_ sdk.Msg = &MsgTransferDomain{}
	_ sdk.Msg = &MsgEditDomain{}
	_ sdk.Msg = &MsgRegisterDomain{}
	_ sdk.Msg = &MsgRemoveDomain{}
)

// NewMsgTransferDomain is a constructor function for MsgSetName
func NewMsgTransferDomain(
	tokenID string, transferPrice uint64, cancelTransfer bool, sender, recipient string,
) *MsgTransferDomain {
	return &MsgTransferDomain{
		Id:        tokenID,
		Price:     transferPrice,
		Cancel:    cancelTransfer,
		Sender:    sender,
		Recipient: recipient,
	}
}

// Route Implements Msg
func (msg MsgTransferDomain) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgTransferDomain) Type() string { return TypeMsgTransferDomain }

// ValidateBasic Implements Msg.
func (msg MsgTransferDomain) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return ValidateTokenID(msg.Id)
}

// GetSignBytes Implements Msg.
func (msg MsgTransferDomain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgTransferDomain) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgEditDomain is a constructor function for MsgSetName
func NewMsgEditDomain(tokenID, tokenData, sender string) *MsgEditDomain {
	return &MsgEditDomain{
		Id:     tokenID,
		Data:   tokenData,
		Sender: sender,
	}
}

// Route Implements Msg
func (msg MsgEditDomain) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgEditDomain) Type() string { return TypeMsgEditDomain }

// ValidateBasic Implements Msg.
func (msg MsgEditDomain) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return ValidateTokenID(msg.Id)
}

// GetSignBytes Implements Msg.
func (msg MsgEditDomain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgEditDomain) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgRegisterDomain is a constructor function for MsgMintNFT
func NewMsgRegisterDomain(tokenID, tokenData, sender, recipient string) *MsgRegisterDomain {
	return &MsgRegisterDomain{
		Id:        tokenID,
		Data:      tokenData,
		Sender:    sender,
		Recipient: recipient,
	}
}

// Route Implements Msg
func (msg MsgRegisterDomain) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgRegisterDomain) Type() string { return TypeMsgRegisterDomain }

// ValidateBasic Implements Msg.
func (msg MsgRegisterDomain) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	if _, err := sdk.AccAddressFromBech32(msg.Recipient); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receipt address (%s)", err)
	}
	return ValidateTokenID(msg.Id)
}

// GetSignBytes Implements Msg.
func (msg MsgRegisterDomain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgRegisterDomain) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgRemoveDomain is a constructor function for MsgBurnNFT
func NewMsgRemoveDomain(sender, tokenID string) *MsgRemoveDomain {
	return &MsgRemoveDomain{
		Sender: sender,
		Id:     tokenID,
	}
}

// Route Implements Msg
func (msg MsgRemoveDomain) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgRemoveDomain) Type() string { return TypeMsgRemoveDomain }

// ValidateBasic Implements Msg.
func (msg MsgRemoveDomain) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return ValidateTokenID(msg.Id)
}

// GetSignBytes Implements Msg.
func (msg MsgRemoveDomain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgRemoveDomain) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}
