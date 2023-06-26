package keeper

import (
	context "context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/terramirum/mirumd/x/rental/types"
)

func (k Keeper) Sessions(context context.Context, req *types.QuerySessionRequest) (*types.QuerySessionResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)
	store := ctx.KVStore(k.storeKey)
	nftRents := []*types.NftRent{}

	var keyRenter []byte
	if len(req.Renter) > 0 {
		if len(req.ClassId) > 0 && len(req.NftId) > 0 && len(req.SessionId) > 0 {
			keyRenter = getStoreWithKey(KeyRentSessionId, req.Renter, req.ClassId, req.NftId, req.SessionId)
		} else if len(req.ClassId) > 0 && len(req.NftId) > 0 {
			keyRenter = getStoreWithKey(KeyRentSessionId, req.Renter, req.ClassId, req.NftId)
		} else if len(req.ClassId) > 0 {
			keyRenter = getStoreWithKey(KeyRentSessionId, req.Renter, req.ClassId)
		} else {
			keyRenter = getStoreWithKey(KeyRentSessionId, req.Renter)
		}
	} else if len(req.ClassId) > 0 && len(req.NftId) > 0 {
		keyRenter = getStoreWithKey(KeyRentDates, req.ClassId, req.NftId)
	} else if len(req.ClassId) > 0 {
		keyRenter = getStoreWithKey(KeyRentDates, req.ClassId)
	} else {
		return nil, sdkerrors.Wrap(types.ErrQuerySessions, "")
	}

	allSessionStore := prefix.NewStore(store, keyRenter)
	iterator := allSessionStore.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var nftRent types.NftRent
		k.cdc.MustUnmarshal(iterator.Value(), &nftRent)
		nftRents = append(nftRents, &nftRent)
	}

	res := &types.QuerySessionResponse{
		NftRent: nftRents,
	}

	return res, nil
}
