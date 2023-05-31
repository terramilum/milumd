package keeper

import (
	context "context"

	"github.com/terramirum/mirumd/x/rental/types"
)

type msgServer struct {
	Keeper
}

// BurnNft implements types.MsgServer
func (*msgServer) BurnNft(context.Context, *types.MsgDeployNftRequest) (*types.MsgDeployNftResponse, error) {
	panic("unimplemented")
}

// DeployNft implements types.MsgServer
func (*msgServer) DeployNft(context.Context, *types.MsgDeployNftRequest) (*types.MsgDeployNftResponse, error) {
	panic("unimplemented")
}

// MintNft implements types.MsgServer
func (*msgServer) MintNft(context.Context, *types.MsgMintNftRequest) (*types.MsgMintNftResponse, error) {
	panic("unimplemented")
}

// NftAccess implements types.MsgServer
func (*msgServer) NftAccess(context.Context, *types.MsgAccessNftRequest) (*types.MsgAccessNftResponse, error) {
	panic("unimplemented")
}

// RentNftBurn implements types.MsgServer
func (*msgServer) RentNftBurn(context.Context, *types.MsgDeployNftRequest) (*types.MsgDeployNftResponse, error) {
	panic("unimplemented")
}

// RentNftMint implements types.MsgServer
func (*msgServer) RentNftMint(context.Context, *types.MsgRentNftRequest) (*types.MsgDeployNftResponse, error) {
	panic("unimplemented")
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = &msgServer{}
