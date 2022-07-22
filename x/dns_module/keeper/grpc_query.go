package keeper

import (
	"context"
	"fmt"
	"sort"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/deweb-services/deweb/x/dns_module/types"
)

type queryServer struct {
	Keeper
	dnsDenomName string
}

var _ types.QueryServer = queryServer{}

const (
	DefaultDomainsCount = 10
)

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
	domainRecordData, err := ParseDomainData([]byte(baseDomainNFT.GetData()))
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal,
			fmt.Sprintf("cannot process record for domain %s", baseDomainNFT.GetID()), err)
	}
	var domainTransferOffer *types.TransferOffer
	if domainRecordData.TransferOffer != nil {
		domainTransferOffer = &types.TransferOffer{
			Price:                domainRecordData.TransferOffer.Price,
			ExpectedOwnerAddress: domainRecordData.TransferOffer.ExpectedOwnerAddress,
		}
	}
	resRecords := make([]*types.DNSRecords, 0, len(domainRecordData.Records))
	for _, rec := range domainRecordData.Records {
		resRecord := &types.DNSRecords{
			Type:   rec.RecordType,
			Values: rec.RecordValues,
		}
		resRecords = append(resRecords, resRecord)
	}
	respDomain := &types.ResponseDomain{
		Id:                  baseDomainNFT.Id,
		Issued:              domainRecordData.Issued.Format(time.RFC3339),
		ValidTil:            domainRecordData.ValidTill.Format(time.RFC3339),
		TransferOffer:       domainTransferOffer,
		Records:             resRecords,
		SubDomainsOnSale:    domainRecordData.SubDomainsOnSale,
		SubDomainsSalePrice: domainRecordData.SubDomainsSalePrice,
		Owner:               baseDomainNFT.Owner,
	}
	return &types.QueryDomainResponse{Domain: respDomain}, nil
}

func (q queryServer) OwnedDomains(c context.Context, request *types.QueryOwnedDomainsRequest) (*types.QueryOwnedDomainsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	address, err := sdk.AccAddressFromBech32(request.Address)
	if err != nil {
		return nil, err
	}
	ownedCollections := q.GetByOwner(ctx, address)

	var domains []string
	for _, collection := range ownedCollections.IDCollections {
		if collection.DenomId != q.dnsDenomName {
			continue
		}
		domains = collection.TokenIds
	}
	if domains == nil || len(domains) == 0 {
		return &types.QueryOwnedDomainsResponse{}, nil
	}
	res := &types.QueryOwnedDomainsResponse{
		Total:   int64(len(domains)),
		Domains: make([]*types.ResponseDomain, 0, request.Count),
	}
	sort.Strings(domains)
	count := int(request.Count)
	if count == 0 {
		count = DefaultDomainsCount
	}
	offset := int(request.Offset)
	for i, domainName := range domains {
		if i < offset {
			continue
		}
		req := &types.QueryDomainRequest{DomainName: domainName}
		domainResp, err := q.Domain(c, req)
		if err != nil {
			return nil, sdkerrors.Wrapf(err, "cannot get info for domain %s", domainName)
		}
		res.Domains = append(res.Domains, domainResp.Domain)
		if len(res.Domains) == count {
			break
		}
	}
	return res, nil
}
