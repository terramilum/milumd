package rental_test

import (
	"testing"

	keepertest "milumd/testutil/keeper"
	"milumd/testutil/nullify"
	"milumd/x/rental"
	"milumd/x/rental/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RentalKeeper(t)
	rental.InitGenesis(ctx, *k, genesisState)
	got := rental.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}
