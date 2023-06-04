package keeper

import (
	context "context"

	"github.com/terramirum/mirumd/x/rental/types"
)

// RentNftBurn implements types.MsgServer
func (k Keeper) RentNftBurn(context.Context, *types.MsgBurnRentRequest) (*types.MsgBurnRentResponse, error) {
	return &types.MsgBurnRentResponse{}, nil
}
