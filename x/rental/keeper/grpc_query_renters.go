package keeper

import (
	context "context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/terramirum/mirumd/x/rental/types"
)

func (k Keeper) Renters(c context.Context, req *types.QueryRenterRequest) (*types.QueryRenterResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)

	renterResponse := &types.QueryRenterResponse{
		Renter: []string{},
	}

	sessionIdKey := nftSessionIdRentersStoreKey(req.ClassId, req.NftId, req.SessionId)
	allSessionStore := prefix.NewStore(store, sessionIdKey)
	iterator := allSessionStore.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		address := iterator.Value()
		renterResponse.Renter = append(renterResponse.Renter, UnsafeBytesToStr(address))
	}

	return renterResponse, nil
}
