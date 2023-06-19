package keeper

import (
	context "context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/terramirum/mirumd/x/rental/types"
)

func (k Keeper) RentNftGiveAccess(context context.Context, rentGiveAccessRequest *types.MsgRentGiveAccessRequest) (*types.MsgRentGiveAccessResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)
	store := ctx.KVStore(k.storeKey)

	var renters []string
	key := getStoreWithKey(KeyRentSessionId, rentGiveAccessRequest.ClassId, rentGiveAccessRequest.NftId, rentGiveAccessRequest.SessionId)
	rentersStore := prefix.NewStore(store, key)
	iterator := rentersStore.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		renters = append(renters, string(iterator.Value()))
	}

	renterExists := false
	for _, v := range renters {
		if v == rentGiveAccessRequest.Renter {
			renterExists = true
			store.Set(key, []byte(rentGiveAccessRequest.NewRenter))
		}
	}

	if !renterExists {
		return nil, sdkerrors.Wrap(types.ErrNftRentAccessGive, "")
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeAccessGiveNft,
		sdk.NewAttribute(types.AttributeKeyClassId, rentGiveAccessRequest.ClassId),
		sdk.NewAttribute(types.AttributeKeyNftId, rentGiveAccessRequest.NftId),
		sdk.NewAttribute(types.AttributeKeySessionId, rentGiveAccessRequest.SessionId),
		sdk.NewAttribute(types.AttributeKeyNftRenter, rentGiveAccessRequest.Renter),
		sdk.NewAttribute(types.AttributeKeyNftNewRenter, rentGiveAccessRequest.NewRenter),
	))

	return &types.MsgRentGiveAccessResponse{}, nil
}
