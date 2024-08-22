package keeper

import (
	context "context"
	"github.com/cosmos/cosmos-sdk/runtime"

	"cosmossdk.io/store/prefix"
	"github.com/terramirum/mirumd/x/rental/types"
)

func (k Keeper) Renters(c context.Context, req *types.QueryRenterRequest) (*types.QueryRenterResponse, error) {
	store := k.storeService.OpenKVStore(c)

	renterResponse := &types.QueryRenterResponse{
		Renter: []string{},
	}

	sessionIdKey := getStoreWithKey(KeyRentSessionId, req.ClassId, req.NftId, req.SessionId)
	allSessionStore := prefix.NewStore(runtime.KVStoreAdapter(store), sessionIdKey)
	iterator := allSessionStore.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		address := iterator.Value()
		renterResponse.Renter = append(renterResponse.Renter, UnsafeBytesToStr(address))
	}

	return renterResponse, nil
}
