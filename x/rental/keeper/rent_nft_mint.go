package keeper

import (
	context "context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	nft "github.com/cosmos/cosmos-sdk/x/nft"
	types "github.com/terramirum/mirumd/x/rental/types"
)

// RentNftMint implements types.MsgServer
func (k Keeper) RentNftMint(context context.Context, rentNftRequest *types.MsgMintRentRequest) (*types.MsgMintRentResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)

	if !k.nftKeeper.HasClass(ctx, rentNftRequest.ClassId) {
		return nil, sdkerrors.Wrap(nft.ErrClassNotExists, rentNftRequest.ClassId)
	}

	if k.nftKeeper.HasNFT(ctx, rentNftRequest.ClassId, rentNftRequest.NftId) {
		return nil, sdkerrors.Wrap(nft.ErrNFTExists, rentNftRequest.NftId)
	}

	err := ctx.EventManager().EmitTypedEvent(&types.EventNftRentMint{
		Class:     rentNftRequest.ClassId,
		Id:        rentNftRequest.NftId,
		StartDate: rentNftRequest.StartDateUnix,
		EndDate:   rentNftRequest.EndDateUnix,
	})

	if err != nil {
		return nil, err
	}

	return &types.MsgMintRentResponse{}, nil
}
