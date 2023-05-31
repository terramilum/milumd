package keeper

import (
	context "context"

	"github.com/terramirum/mirumd/x/rental/types"
)

// NftAccess implements types.MsgServer
func (k Keeper) NftAccess(context.Context, *types.MsgAccessNftRequest) (*types.MsgAccessNftResponse, error) {
	panic("unimplemented")
}
