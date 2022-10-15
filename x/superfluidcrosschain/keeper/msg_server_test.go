package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/PerseusCore/superfluid-cross-chain/x/superfluidcrosschain/types"
    "github.com/PerseusCore/superfluid-cross-chain/x/superfluidcrosschain/keeper"
    keepertest "github.com/PerseusCore/superfluid-cross-chain/testutil/keeper"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.SuperfluidcrosschainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
