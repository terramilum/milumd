package keeper

import (
	context "context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/terramirum/mirumd/x/rental/types"
)

func (k Keeper) Classes(c context.Context, req *types.QueryClassRequest) (*types.QueryClassResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)

	var nftClasses []*types.NftClass

	classIdKey := contractAddressClassIdKey(req.ContractOwner)
	allSessionStore := prefix.NewStore(store, classIdKey)
	iterator := allSessionStore.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		classId := string(iterator.Value())
		class, _ := k.nftKeeper.GetClass(ctx, classId)
		nftClass := &types.NftClass{
			Id:          classId,
			Name:        class.Name,
			Symbol:      class.Symbol,
			Description: class.Description,
			Uri:         class.Uri,
		}
		nftClasses = append(nftClasses, nftClass)
	}
	return &types.QueryClassResponse{
		NftClass: nftClasses,
	}, nil
}

func (k Keeper) Nfts(c context.Context, req *types.QueryNftRequest) (*types.QueryNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	nftsRequest := &nft.QueryNFTsRequest{
		ClassId: req.ClassId,
		Pagination: &query.PageRequest{
			Key:        []byte{},
			Offset:     0,
			Limit:      1000,
			CountTotal: false,
			Reverse:    false,
		},
	}

	nfts, err := k.nftKeeper.NFTs(ctx, nftsRequest)
	if err != nil {
		return nil, err
	}

	var nftDefinitions []*types.NftDefinition
	for _, v := range nfts.Nfts {
		var rentDetail types.NftRentDetail
		err = k.cdc.UnpackAny(v.Data, &rentDetail)
		if err != nil {
			return nil, err
		}
		nftDefinition := &types.NftDefinition{
			ClassId:       v.ClassId,
			Id:            v.Id,
			Uri:           v.Uri,
			NftRentDetail: &rentDetail,
		}
		nftDefinitions = append(nftDefinitions, nftDefinition)
	}
	return &types.QueryNftResponse{
		NftDefinition: nftDefinitions,
	}, nil
}
