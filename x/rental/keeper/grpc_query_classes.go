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

	classIdKey := getStoreWithKey(KeyContractClassId, req.ContractOwner)
	iterator := sdk.KVStorePrefixIterator(store, classIdKey)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		classId := getParsedStoreKey(iterator.Key())[2]
		val := string(iterator.Value())
		if val != "1" {
			continue
		}
		class, _ := k.nftKeeper.GetClass(ctx, classId)

		classDetail := &types.Detail{}
		err := classDetail.Unmarshal(class.Data.Value)
		if err != nil {
			return nil, err
		}

		nftClass := &types.NftClass{
			Id:          classId,
			Name:        class.Name,
			Symbol:      class.Symbol,
			Description: class.Description,
			Uri:         class.Uri,
			Detail:      classDetail,
		}
		nftClasses = append(nftClasses, nftClass)
	}
	return &types.QueryClassResponse{
		NftClass: nftClasses,
	}, nil
}
