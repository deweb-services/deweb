package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/deweb-services/deweb/x/dns_module/types"
)

type queryServer struct {
	Keeper
	dnsDenomName string
}

var _ types.QueryServer = queryServer{}

// NewQueryServerImpl returns an implementation of the NFT QueryServer interface
// for the provided Keeper.
func NewQueryServerImpl(keeper Keeper) types.QueryServer {
	return &queryServer{
		Keeper:       keeper,
		dnsDenomName: DNSDenomName,
	}
}

func (q queryServer) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	params := q.GetParams(ctx)

	return &types.QueryParamsResponse{
		Params: params,
	}, nil
}

func (q queryServer) Domain(c context.Context, request *types.QueryDomainRequest) (*types.QueryDomainResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	nft, err := q.GetDomain(ctx, request.DomainName)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrUnknownDomain, "invalid domains name %s ", request.DomainName)
	}

	baseDomainNFT, ok := nft.(types.BaseDomain)
	if !ok {
		return nil, sdkerrors.Wrapf(types.ErrUnknownDomain, "invalid type of domain NFT %s", request.DomainName)
	}

	return &types.QueryDomainResponse{Domain: &baseDomainNFT}, nil
}
