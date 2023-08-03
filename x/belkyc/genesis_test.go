package belkyc_test

import (
	"testing"

	keepertest "github.com/anshulWeb3/belkyc/testutil/keeper"
	"github.com/anshulWeb3/belkyc/testutil/nullify"

	"github.com/anshulWeb3/belkyc/x/belkyc"
	"github.com/anshulWeb3/belkyc/x/belkyc/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		KycList: []types.Kyc{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BelkycKeeper(t)
	belkyc.InitGenesis(ctx, *k, genesisState)
	got := belkyc.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.KycList, got.KycList)
	// this line is used by starport scaffolding # genesis/test/assert
}
