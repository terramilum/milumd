package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "milumd/testutil/keeper"
	"milumd/x/rental/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.RentalKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
