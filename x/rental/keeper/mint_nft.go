package keeper

import (
	context "context"
	"fmt"

	codec "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/terramirum/mirumd/x/rental/types"
)

// MintNft implements types.MsgServer
func (k Keeper) MintNft(context context.Context, mintRequest *types.MsgMintNftRequest) (*types.MsgMintNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)

	store := ctx.KVStore(k.storeKey)
	contractAddress := store.Get(classContractAddressKey(mintRequest.ClassId))
	if string(contractAddress) != mintRequest.ContractOwner {
		return nil, sdkerrors.Wrap(types.ErrNftClassOwnerTheSame, "")
	}

	nfts := k.nftKeeper.GetNFTsOfClass(ctx, mintRequest.ClassId)
	nftId := fmt.Sprintf("%d", len(nfts)+1)
	if len(mintRequest.NftId) > 0 {
		nftId = mintRequest.NftId
	}

	rentDetail, err := codec.NewAnyWithValue(mintRequest.NftRentDetail)
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
