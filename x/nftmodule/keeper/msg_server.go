package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/deweb-services/deweb/x/nftmodule/types"
)

type msgServer struct {
	Keeper
	dnsDenomName string
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the NFT MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{
		Keeper:       keeper,
		dnsDenomName: DNSDenomName,
	}
}

// IssueDenom issue a new denom.
func (m msgServer) IssueDenom(goCtx context.Context, msg *types.MsgIssueDenom) (*types.MsgIssueDenomResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	if msg.Id != m.dnsDenomName {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDenom, "only denomID %s allowed", m.dnsDenomName)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.Keeper.IssueDenom(ctx, msg.Id, msg.Name, msg.Schema, msg.Symbol, sender,
		msg.MintRestricted, msg.UpdateRestricted,
		msg.Description, msg.Uri, msg.UriHash, msg.Data,
	); err != nil {
		return nil, err
	}

	return &types.MsgIssueDenomResponse{}, nil
}

func (m msgServer) MintNFT(goCtx context.Context, msg *types.MsgMintNFT) (*types.MsgMintNFTResponse, error) {
	recipient, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, err
	}

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.MintNFT(ctx, msg.Id, msg.Data, sender, recipient); err != nil {
		return nil, err
	}

	return &types.MsgMintNFTResponse{}, nil
}

func (m msgServer) EditNFT(goCtx context.Context, msg *types.MsgEditNFT) (*types.MsgEditNFTResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	_, err = ParseNFTData([]byte(msg.Data))
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal, "invalid data")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.Keeper.EditNFT(ctx, msg.Id, msg.Data, sender); err != nil {
		return nil, err
	}

	return &types.MsgEditNFTResponse{}, nil
}

func (m msgServer) TransferNFT(goCtx context.Context, msg *types.MsgTransferNFT) (*types.MsgTransferNFTResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	if len(msg.Recipient) > 0 {
		_, err = sdk.AccAddressFromBech32(msg.Recipient)
		if err != nil {
			return nil, err
		}
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.Keeper.TransferDomainOwner(ctx, msg.Id, msg.Cancel, msg.Price, sender, msg.Recipient); err != nil {
		return nil, err
	}

	return &types.MsgTransferNFTResponse{}, nil
}

func (m msgServer) BurnNFT(goCtx context.Context, msg *types.MsgBurnNFT) (*types.MsgBurnNFTResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.Keeper.BurnNFT(ctx, msg.Id, sender); err != nil {
		return nil, err
	}

	return &types.MsgBurnNFTResponse{}, nil
}
