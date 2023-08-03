package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/anshulWeb3/belkyc/testutil/keeper"

	"github.com/anshulWeb3/belkyc/x/belkyc/keeper"
	"github.com/anshulWeb3/belkyc/x/belkyc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BelkycKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}
