package keeper

import (
	context "context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/terramirum/mirumd/x/rental/types"
)

func (k Keeper) Classes(c context.Context, req *types.QueryClassRequest) (*types.QueryClassResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)

	var nftClasses []*types.NftClass

	classIdKey := contractOwnerClasseseKey(req.ContractOwner)
	iterator := sdk.KVStorePrefixIterator(store, classIdKey)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		_, classId := parseContractAddressClassIdKey(iterator.Key())
		val := string(iterator.Value())
		if val != "1" {
			continue
		}
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
