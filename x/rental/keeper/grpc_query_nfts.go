package keeper

import (
	context "context"

	"cosmossdk.io/x/nft"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/terramirum/mirumd/x/rental/types"
)

func (k Keeper) Nfts(c context.Context, req *types.QueryNftRequest) (*types.QueryNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	pageRequest := &query.PageRequest{
		Offset:     0,
		Limit:      1000,
		CountTotal: false,
		Reverse:    false,
	}

	nftsRequest := &nft.QueryNFTsRequest{
		ClassId:    req.ClassId,
		Pagination: pageRequest,
	}

	classesRequest := &nft.QueryClassesRequest{
		Pagination: pageRequest,
	}
	clss, err := k.nftKeeper.Classes(ctx, classesRequest)
	if err != nil {
		return nil, err
	}
	_ = clss

	nfts, err := k.nftKeeper.NFTs(ctx, nftsRequest)
	if err != nil {
		return nil, err
	}

	var nftDefinitions []*types.NftDefinition
	for _, v := range nfts.Nfts {
		rentDetail := &types.Detail{}
		err = rentDetail.Unmarshal(v.Data.Value)
		if err != nil {
			return nil, err
		}

		nftDefinition := &types.NftDefinition{
			ClassId: v.ClassId,
			Id:      v.Id,
			Uri:     v.Uri,
			Detail:  rentDetail,
		}
		nftDefinitions = append(nftDefinitions, nftDefinition)
	}
	return &types.QueryNftResponse{
		NftDefinition: nftDefinitions,
	}, nil
}
