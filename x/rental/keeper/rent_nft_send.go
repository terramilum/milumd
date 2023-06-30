package keeper

import (
	context "context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	types "github.com/terramirum/mirumd/x/rental/types"
)

// SendSession implements types.MsgServer
func (k Keeper) SendSession(context context.Context, sendSessionRequest *types.MsgSendSessionRequest) (*types.MsgSendSessionResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)
	store := ctx.KVStore(k.storeKey)

	if len(sendSessionRequest.ClassId) == 0 {
		return nil, sdkerrors.Wrap(types.ErrFieldIsRequired, "Class Id")
	}

	if len(sendSessionRequest.NftId) == 0 {
		return nil, sdkerrors.Wrap(types.ErrFieldIsRequired, "Class Id")
	}

	querySessions := &types.QuerySessionRequest{
		ClassId:   sendSessionRequest.ClassId,
		NftId:     sendSessionRequest.NftId,
		Renter:    sendSessionRequest.FromRenter,
		SessionId: sendSessionRequest.SessionId,
	}
	res, err := k.Sessions(context, querySessions)
	if err != nil {
		return nil, err
	}

	if len(res.SessionDetail) != 1 {
		return nil, sdkerrors.Wrap(types.ErrQuerySessionsNotFound, "")
	}

	currentDate := getNowUtc()

	if res.SessionDetail[0].NftRent.EndDate < currentDate {
		nftRents := k.toNftRent(res.SessionDetail)
		k.clearOldSession(ctx, sendSessionRequest.ClassId, sendSessionRequest.NftId, nftRents)
		return nil, sdkerrors.Wrap(types.ErrQueryOldSessionsNotTransfer, "")
	}

	// storing owner to KeyRentDates.
	keySessionOwner := getStoreWithKey(KeyRentDatesOwner, sendSessionRequest.ClassId, sendSessionRequest.NftId, sendSessionRequest.SessionId, sendSessionRequest.FromRenter)
	sessionOwner := store.Get(keySessionOwner)
	if sendSessionRequest.FromRenter != string(sessionOwner) {
		return nil, sdkerrors.Wrap(types.ErrSessionOwnerCanTransfer, "")
	}

	// get session info
	rentersKey := getStoreWithKey(KeyRentSessionId, sendSessionRequest.FromRenter, sendSessionRequest.ClassId, sendSessionRequest.NftId, sendSessionRequest.SessionId)
	rentSession := store.Get(rentersKey)

	// clear old owners datas
	k.clearKeyRentDatesOwner(ctx, sendSessionRequest.ClassId, sendSessionRequest.NftId, sendSessionRequest.SessionId)

	// setting new owner.
	keySessionOwner = getStoreWithKey(KeyRentDatesOwner, sendSessionRequest.ClassId, sendSessionRequest.NftId, sendSessionRequest.SessionId, sendSessionRequest.ToRenter)
	store.Set(keySessionOwner, []byte(sendSessionRequest.ToRenter))

	// setting new owner.
	rentersKey = getStoreWithKey(KeyRentSessionId, sendSessionRequest.ToRenter, sendSessionRequest.ClassId, sendSessionRequest.NftId, sendSessionRequest.SessionId)
	store.Set(rentersKey, rentSession)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeRentSend,
		sdk.NewAttribute(types.AttributeKeyNftFromRenter, sendSessionRequest.FromRenter),
		sdk.NewAttribute(types.AttributeKeyNftToRenter, sendSessionRequest.ToRenter),
		sdk.NewAttribute(types.AttributeKeyClassId, sendSessionRequest.ClassId),
		sdk.NewAttribute(types.AttributeKeyNftId, sendSessionRequest.NftId),
		sdk.NewAttribute(types.AttributeKeySessionId, sendSessionRequest.SessionId),
	))

	return &types.MsgSendSessionResponse{}, nil
}
