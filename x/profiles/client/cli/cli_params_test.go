//go:build norace
// +build norace

package cli_test

import (
	"fmt"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/desmos-labs/desmos/v3/x/profiles/client/cli"
	"github.com/desmos-labs/desmos/v3/x/profiles/types"
)

func (s *IntegrationTestSuite) TestCmdQueryParams() {
	val := s.network.Validators[0]

	testCases := []struct {
		name           string
		args           []string
		shouldErr      bool
		expectedOutput types.QueryParamsResponse
	}{
		{
			name:      "existing params are returned properly",
			args:      []string{fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			shouldErr: false,
			expectedOutput: types.QueryParamsResponse{
				Params: types.DefaultParams(),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryParams()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryParamsResponse
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expectedOutput, response)
			}
		})
	}
}
