package superfluidcrosschain_test

import (
	"testing"

	keepertest "github.com/PerseusCore/superfluid-cross-chain/testutil/keeper"
	"github.com/PerseusCore/superfluid-cross-chain/testutil/nullify"
	"github.com/PerseusCore/superfluid-cross-chain/x/superfluidcrosschain"
	"github.com/PerseusCore/superfluid-cross-chain/x/superfluidcrosschain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SuperfluidcrosschainKeeper(t)
	superfluidcrosschain.InitGenesis(ctx, *k, genesisState)
	got := superfluidcrosschain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}
