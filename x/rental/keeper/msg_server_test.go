package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/terramirum/mirumd/testutil/keeper"
	"github.com/terramirum/mirumd/x/rental/keeper"
	"github.com/terramirum/mirumd/x/rental/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.RentalKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
