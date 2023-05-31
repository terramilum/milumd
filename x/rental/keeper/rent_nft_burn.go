package keeper

import (
	context "context"

	"github.com/terramirum/mirumd/x/rental/types"
)

// RentNftBurn implements types.MsgServer
func (k Keeper) RentNftBurn(context.Context, *types.MsgRentNftRequest) (*types.MsgRentNftResponse, error) {
	return &types.MsgRentNftResponse{}, nil
}
