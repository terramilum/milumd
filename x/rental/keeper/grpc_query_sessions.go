package keeper

import (
	context "context"

	sdkstore "cosmossdk.io/core/store"
	sdkerrors "cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/terramirum/mirumd/x/rental/types"
)

func (k Keeper) Sessions(context context.Context, req *types.QuerySessionRequest) (*types.QuerySessionResponse, error) {
	store := k.storeService.OpenKVStore(context)
	sessionDetails := []*types.SessionDetail{}

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

	allSessionStore := prefix.NewStore(runtime.KVStoreAdapter(store), keyRenter)
	iterator := allSessionStore.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		sessionDetail, err := k.getSessionDetail(keyRenter, iterator.Key(), req)
		if err != nil {
			return nil, err
		}

		var nftRent types.NftRent
		k.cdc.MustUnmarshal(iterator.Value(), &nftRent)
		if len(req.SessionId) > 0 {
			if req.SessionId == nftRent.SessionId {
				sessionDetail.NftRent = &nftRent
				sessionDetails = append(sessionDetails, sessionDetail)
			}
		} else {
			sessionDetail.NftRent = &nftRent
			sessionDetails = append(sessionDetails, sessionDetail)
		}
	}

	for _, v := range sessionDetails {
		if len(v.Renter) == 0 {
			v.Renter = k.getOwnerOfSession(v, store)
		}
	}

	return &types.QuerySessionResponse{
		SessionDetail: sessionDetails,
	}, nil
}

func (k Keeper) getSessionDetail(queryKeyFirst, queryKeySecond []byte, req *types.QuerySessionRequest) (*types.SessionDetail, error) {
	keys := getParsedStoreKey(queryKeyFirst)
	keys = append(keys, getParsedStoreKey(queryKeySecond)...)
	if keys[0] == string(KeyRentDates) {
		return &types.SessionDetail{
			Renter:  "",
			ClassId: keys[1],
			NftId:   keys[2],
		}, nil
	} else if keys[0] == string(KeyRentSessionId) {
		return &types.SessionDetail{
			Renter:  keys[1],
			ClassId: keys[3],
			NftId:   keys[4],
		}, nil
	} else {
		return &types.SessionDetail{
			Renter:  "",
			ClassId: "",
			NftId:   "",
		}, nil
	}
}

func (k Keeper) toNftRent(sessionDetails []*types.SessionDetail) []*types.NftRent {
	var nftRents []*types.NftRent
	for _, v := range sessionDetails {
		nftRents = append(nftRents, v.NftRent)
	}
	return nftRents
}

func (k Keeper) getOwnerOfSession(sessionDetail *types.SessionDetail, store sdkstore.KVStore) string {
	keySessionOwner := getStoreWithKey(KeyRentDatesOwner, sessionDetail.ClassId,
		sessionDetail.NftId, sessionDetail.NftRent.SessionId)
	allSessionStore := prefix.NewStore(runtime.KVStoreAdapter(store), keySessionOwner)
	iterator := allSessionStore.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		keys := getParsedStoreKey(iterator.Key())
		if keys[1] == string(iterator.Value()) {
			return keys[1]
		}
	}
	return ""
}
