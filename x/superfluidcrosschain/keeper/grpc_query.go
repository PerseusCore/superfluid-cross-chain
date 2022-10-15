package keeper

import (
	"github.com/PerseusCore/superfluid-cross-chain/x/superfluidcrosschain/types"
)

var _ types.QueryServer = Keeper{}
