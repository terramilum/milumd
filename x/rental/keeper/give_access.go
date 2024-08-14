package keeper

import (
	context "context"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/terramirum/mirumd/x/rental/types"
)

func (k Keeper) RentNftGiveAccess(context context.Context, rentGiveAccessRequest *types.MsgRentGiveAccessRequest) (*types.MsgRentGiveAccessResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)
	store := k.storeService.OpenKVStore(ctx)

	req := &types.QuerySessionRequest{
		ClassId:   rentGiveAccessRequest.ClassId,
		NftId:     rentGiveAccessRequest.NftId,
		Renter:    rentGiveAccessRequest.Renter,
		SessionId: rentGiveAccessRequest.SessionId,
	}

	res, err := k.Sessions(context, req)
	if err != nil {
		return nil, err
	}

	if len(res.SessionDetail) != 1 {
		return nil, sdkerrors.Wrap(types.ErrQuerySessionsNotFound, "")
	}

	rentOwnerKey := getStoreWithKey(KeyRentDatesOwner, rentGiveAccessRequest.ClassId, rentGiveAccessRequest.NftId, rentGiveAccessRequest.SessionId, rentGiveAccessRequest.Renter)
	rentOwner, err := store.Get(rentOwnerKey)
	if err != nil {
		return nil, err
	}
	if rentGiveAccessRequest.Renter != string(rentOwner) {
		return nil, sdkerrors.Wrap(types.ErrNftRentAccessGive, "")
	}

	newRenter := getStoreWithKey(KeyRentDatesOwner, rentGiveAccessRequest.ClassId, rentGiveAccessRequest.NftId, rentGiveAccessRequest.SessionId, rentGiveAccessRequest.NewRenter)
	store.Set(newRenter, rentOwner)

	rentersKey := getStoreWithKey(KeyRentSessionId, rentGiveAccessRequest.NewRenter, rentGiveAccessRequest.ClassId, rentGiveAccessRequest.NftId, rentGiveAccessRequest.SessionId)
	bz := k.cdc.MustMarshal(res.SessionDetail[0].NftRent)
	store.Set(rentersKey, bz)

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
