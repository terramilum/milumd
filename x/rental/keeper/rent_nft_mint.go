package keeper

import (
	context "context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	nft "github.com/cosmos/cosmos-sdk/x/nft"
	types "github.com/terramirum/mirumd/x/rental/types"
)

// RentNftMint implements types.MsgServer
func (k Keeper) RentNftMint(context context.Context, rentNftRequest *types.MsgMintRentRequest) (*types.MsgMintRentResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)
	store := ctx.KVStore(k.storeKey)

	if !k.nftKeeper.HasClass(ctx, rentNftRequest.ClassId) {
		return nil, sdkerrors.Wrap(nft.ErrClassNotExists, rentNftRequest.ClassId)
	}

	if !k.nftKeeper.HasNFT(ctx, rentNftRequest.ClassId, rentNftRequest.NftId) {
		return nil, sdkerrors.Wrap(nft.ErrNFTExists, rentNftRequest.NftId)
	}

	nftOwner := k.nftKeeper.GetOwner(ctx, rentNftRequest.ClassId, rentNftRequest.NftId)
	if nftOwner.String() != rentNftRequest.ContractOwner {
		sdkerrors.Wrap(types.ErrNftOwnerCanRent, "")
	}

	params := k.GetParams(ctx)

	currentDate := getNowUtcAddMin(params.RentMinStartUnit)
	if rentNftRequest.StartDate < currentDate {
		sdkerrors.Wrap(types.ErrNftOwnerCanRent, fmt.Sprintf("Start Date: %d", currentDate))
	}

	nftRents := k.GetSessionIdsOfNft(ctx, rentNftRequest.ClassId, rentNftRequest.NftId)
	k.clearOldSession(ctx, rentNftRequest.ClassId, rentNftRequest.NftId, nftRents)

	if k.hasAvaliableSession(nftRents, rentNftRequest) {
		sessionId := fmt.Sprintf("%d", rentNftRequest.StartDate)
		keySession := getStoreWithKey(KeyRentSessionId, rentNftRequest.ClassId, rentNftRequest.NftId, sessionId)
		nftRent := &types.NftRent{
			SessionId: sessionId,
			StartDate: rentNftRequest.StartDate,
			EndDate:   rentNftRequest.EndDate,
		}
		bz := k.cdc.MustMarshal(nftRent)
		store.Set(keySession, bz)

		keyRenter := getStoreWithKey(KeyRentSessionId, rentNftRequest.ClassId, rentNftRequest.NftId, rentNftRequest.Renter)
		store.Set(keyRenter, keySession)

		rentersKey := getStoreWithKey(KeyRentSessionId, rentNftRequest.ClassId, rentNftRequest.NftId, sessionId, rentNftRequest.Renter)
		store.Set(rentersKey, UnsafeStrToBytes("1"))
	} else {
		sdkerrors.Wrap(types.ErrNftOwnerCanRent, fmt.Sprintf("Start Date: %d", currentDate))
	}

	// ilgili tarihler icin mint eden adress başka kişiye yetki verebilir.
	// Contract owner aynı tarihte birden fazla kişiye yetki verebilir.
	// Bu kişiler birbirile yeni bir session id ile baglanacak
	// herbir tarih yeni bir session id alacak.
	// session id başlangıc ve bitiş toplamı olabilir
	// session id aynı zamanda yetkili kişileri icermekte olacak

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeRentNft,
		sdk.NewAttribute(types.AttributeKeyNftRentReciever, rentNftRequest.Renter),
		sdk.NewAttribute(types.AttributeKeyNftRentStart, fmt.Sprintf("%d", rentNftRequest.StartDate)),
		sdk.NewAttribute(types.AttributeKeyNftRentEnd, fmt.Sprintf("%d", rentNftRequest.EndDate)),
		sdk.NewAttribute(types.AttributeKeyClassId, rentNftRequest.ClassId),
		sdk.NewAttribute(types.AttributeKeyNftId, rentNftRequest.NftId),
	))

	return &types.MsgMintRentResponse{}, nil
}

func (k Keeper) hasAvaliableSession(nftRents []types.NftRent, rentNftRequest *types.MsgMintRentRequest) bool {
	for _, v := range nftRents {
		if v.StartDate <= rentNftRequest.StartDate && v.EndDate >= rentNftRequest.StartDate {
			return false
		}
		if v.StartDate <= rentNftRequest.EndDate && v.EndDate >= rentNftRequest.EndDate {
			return false
		}
	}
	return true
}

func (k Keeper) GetSessionIdsOfNft(ctx sdk.Context, classId, nftId string) (nftRents []types.NftRent) {
	store := ctx.KVStore(k.storeKey)
	key := getStoreWithKey(KeyRentDates, classId, nftId)
	allSessionStore := prefix.NewStore(store, key)
	iterator := allSessionStore.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var nftRent types.NftRent
		k.cdc.MustUnmarshal(iterator.Value(), &nftRent)
		nftRents = append(nftRents, nftRent)
	}
	return nftRents
}

// clear old sessions.
func (k Keeper) clearOldSession(ctx sdk.Context, classId, nftId string, nftRents []types.NftRent) {
	currentDate := getNowUtc()
	store := ctx.KVStore(k.storeKey)
	key := getStoreWithKey(KeyRentDates, classId, nftId)
	allSessionStore := prefix.NewStore(store, key)
	for _, v := range nftRents {
		if v.EndDate < currentDate {
			sessionIdKey := getStoreWithKey(KeyRentSessionId, classId, nftId, v.SessionId)
			allSessionStore.Delete(sessionIdKey)

			k.clearSessionIdRenters(ctx, classId, nftId, v.SessionId)
		}
	}
}

// clear renters given accessed by renter to session id.
func (k Keeper) clearSessionIdRenters(ctx sdk.Context, classId, nftId, sessionId string) {
	store := ctx.KVStore(k.storeKey)
	key := getStoreWithKey(KeyRentSessionId, classId, nftId, sessionId)
	sessionIdRenters := prefix.NewStore(store, key)
	iterator := sessionIdRenters.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		renter := string(iterator.Value())
		keyRenter := getStoreWithKey(KeyRentDates, classId, nftId, renter)
		store.Delete(keyRenter)
	}
	store.Delete(key)
}
