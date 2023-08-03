package keeper_test

import (
	"testing"

	testkeeper "github.com/anshulWeb3/belkyc/testutil/keeper"

	"github.com/anshulWeb3/belkyc/x/belkyc/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BelkycKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
