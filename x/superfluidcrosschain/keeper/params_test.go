package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/PerseusCore/superfluid-cross-chain/testutil/keeper"
	"github.com/PerseusCore/superfluid-cross-chain/x/superfluidcrosschain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SuperfluidcrosschainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
