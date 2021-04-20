package cli_test

import (
	"fmt"
	"testing"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/stretchr/testify/suite"
	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/desmos-labs/desmos/testutil"
	"github.com/desmos-labs/desmos/x/links/client/cli"
	"github.com/desmos-labs/desmos/x/links/types"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	cfg := testutil.DefaultConfig()
	genesisState := cfg.GenesisState
	cfg.NumValidators = 3

	var linksData types.GenesisState
	s.Require().NoError(cfg.Codec.UnmarshalJSON(genesisState[types.ModuleName], &linksData))

	linksData.Links = []types.Link{
		types.NewLink(
			"desmos1tw3jl54lmwn3mq6hjfvl5nsk4q70v34wc9nsyk",
			"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
		),
	}

	linksData.PortId = "links"

	linsDataBz, err := cfg.Codec.MarshalJSON(&linksData)
	s.Require().NoError(err)
	genesisState[types.ModuleName] = linsDataBz
	cfg.GenesisState = genesisState

	s.cfg = cfg
	s.network = network.New(s.T(), cfg)

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

// ___________________________________________________________________________________________________________________

func (s *IntegrationTestSuite) TestCmdQueryLink() {
	val := s.network.Validators[0]

	tests := []struct {
		name      string
		args      []string
		expErr    bool
		expOutput types.QueryLinkResponse
	}{
		{
			name: "empty slice is returned properly",
			args: []string{
				s.network.Validators[2].Address.String(),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			expErr: false,
			expOutput: types.QueryLinkResponse{
				Link: types.Link{},
			},
		},
		{
			name: "existing user blocks are returned properly",
			args: []string{
				"desmos1tw3jl54lmwn3mq6hjfvl5nsk4q70v34wc9nsyk",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			expErr: false,
			expOutput: types.QueryLinkResponse{
				Link: types.Link{},
			},
		},
	}

	for _, test := range tests {
		test := test

		s.Run(test.name, func() {

			cmd := cli.GetCmdQueryLink()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, test.args)

			if test.expErr {
				s.Require().Error(err)

				var response types.QueryLinkResponse
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(test.expOutput, response)
			}
		})
	}
}
