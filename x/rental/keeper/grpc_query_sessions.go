package keeper

import (
	context "context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/terramirum/mirumd/x/rental/types"
)

func (k Keeper) Sessions(context context.Context, req *types.QuerySessionRequest) (*types.QuerySessionResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)
	store := ctx.KVStore(k.storeKey)
	nftRents := []*types.NftRent{}

	keyRenter := getStoreWithKey(KeyRentDates, req.ClassId, req.NftId)
	if len(req.Renter) > 0 {
		keyRenter = getStoreWithKey(KeyRentSessionId, req.ClassId, req.NftId, req.Renter)
	}
	allSessionStore := prefix.NewStore(store, keyRenter)
	iterator := allSessionStore.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		sessionId := string(getParsedStoreKey(iterator.Key())[1])
		var nftRent types.NftRent
		k.cdc.MustUnmarshal(iterator.Value(), &nftRent)
		nftRent.SessionId = sessionId
		nftRents = append(nftRents, &nftRent)
	}

	res := &types.QuerySessionResponse{
		NftRent: nftRents,
	}

	return res, nil
}
