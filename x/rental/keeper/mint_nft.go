package keeper

import (
	context "context"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	"cosmossdk.io/x/nft"
	codec "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/terramirum/mirumd/x/rental/types"
)

// MintNft implements types.MsgServer
func (k Keeper) MintNft(context context.Context, mintRequest *types.MsgMintNftRequest) (*types.MsgMintNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)

	store := k.storeService.OpenKVStore(ctx)

	isOwner, err := store.Get(getStoreWithKey(KeyContractClassId, mintRequest.ContractOwner, mintRequest.ClassId))
	if err != nil {
		return nil, err
	}
	if string(isOwner) != "1" {
		return nil, sdkerrors.Wrap(types.ErrNftClassOwnerTheSame, "")
	}

	nfts := k.nftKeeper.GetNFTsOfClass(ctx, mintRequest.ClassId)
	nftId := fmt.Sprintf("%d", len(nfts)+1)
	if len(mintRequest.NftId) > 0 {
		nftId = mintRequest.NftId
	}

	rentDetail, err := codec.NewAnyWithValue(mintRequest.Detail)
	if err != nil {
		return nil, err
	}

	uri := "/" + nftId
	if len(mintRequest.Uri) > 3 {
		uri = mintRequest.Uri
	}

	nft := nft.NFT{
		ClassId: mintRequest.ClassId,
		Id:      nftId,
		Uri:     uri,
		Data:    rentDetail,
	}

	reciever, err := sdk.AccAddressFromBech32(mintRequest.Reciever)
	if err != nil {
		return nil, err
	}

	err = k.nftKeeper.Mint(ctx, nft, reciever)
	if err != nil {
		return nil, err
	}

	return &types.MsgMintNftResponse{
		NftId: mintRequest.NftId,
	}, nil
}
