package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/deweb-services/deweb/x/dns_module/types"
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

func (m msgServer) RegisterDomain(goCtx context.Context, msg *types.MsgRegisterDomain) (*types.MsgRegisterDomainResponse, error) {
	recipient, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, err
	}

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.RegisterDomain(ctx, msg.Id, msg.Data, sender, recipient); err != nil {
		return nil, err
	}

	return &types.MsgRegisterDomainResponse{}, nil
}

func (m msgServer) EditDomain(goCtx context.Context, msg *types.MsgEditDomain) (*types.MsgEditdomainResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	_, err = ParseDomainData([]byte(msg.Data))
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal, "invalid data")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.Keeper.EditDomain(ctx, msg.Id, msg.Data, sender); err != nil {
		return nil, err
	}

	return &types.MsgEditdomainResponse{}, nil
}

func (m msgServer) TransferDomain(goCtx context.Context, msg *types.MsgTransferDomain) (*types.MsgTransferDomainResponse, error) {
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

	return &types.MsgTransferDomainResponse{}, nil
}

func (m msgServer) RemoveDomain(goCtx context.Context, msg *types.MsgRemoveDomain) (*types.MsgRemoveDomainResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.Keeper.RemoveDomain(ctx, msg.Id, sender); err != nil {
		return nil, err
	}

	return &types.MsgRemoveDomainResponse{}, nil
}
