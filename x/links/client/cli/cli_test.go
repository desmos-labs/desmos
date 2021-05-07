package cli_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/cosmos/go-bip39"
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

// ___________________________________________________________________________________________________________________

func generateKey(keyname string) (keyring.Keyring, keyring.Info) {
	keyBase := keyring.NewInMemory()
	keyringAlgos, _ := keyBase.SupportedAlgorithms()
	algo, _ := keyring.NewSigningAlgoFromString("secp256k1", keyringAlgos)
	hdPath := hd.CreateHDPath(0, 0, 0).String()
	entropySeed, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropySeed)
	key, _ := keyBase.NewAccount(keyname, mnemonic, "", hdPath, algo)
	return keyBase, key
}

func (s *IntegrationTestSuite) TestGetIBCAccountConnectionPacket() {
	var (
		destChainPrefix string
		destKey         keyring.Info
		srcKey          keyring.Info
		destKeybase     keyring.Keyring
		srcKeybase      keyring.Keyring
	)

	tests := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			name: "Get packet successfully",
			malleate: func() {
				destChainPrefix = "cosmos"
				srcKeybase, srcKey = generateKey("test")
				destKeybase, destKey = generateKey("test")
			},
			expPass: true,
		},
		{
			name: "Wrong source key name",
			malleate: func() {
				destChainPrefix = "cosmos"
				srcKeybase, _ = generateKey("test")
				_, srcKey = generateKey("wrong")
				destKeybase, destKey = generateKey("test")
			},
			expPass: false,
		},
		{
			name: "Wrong dest key name",
			malleate: func() {
				destChainPrefix = "cosmos"
				srcKeybase, srcKey = generateKey("test")
				destKeybase, _ = generateKey("test")
				_, destKey = generateKey("wrong")
			},
			expPass: false,
		},
	}

	for _, test := range tests {
		test := test

		s.Run(test.name, func() {
			test.malleate()
			_, err := cli.GetIBCAccountConnectionPacket(srcKeybase, srcKey, destKeybase, destKey, destChainPrefix)

			if !test.expPass {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestGetIBCAccountLinkPacket() {
	var (
		destChainPrefix string
		srcKey          keyring.Info
		srcKeybase      keyring.Keyring
	)

	tests := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			name: "Get packet successfully",
			malleate: func() {
				destChainPrefix = "cosmos"
				srcKeybase, srcKey = generateKey("test")
			},
			expPass: true,
		},
		{
			name: "Wrong source key name",
			malleate: func() {
				destChainPrefix = "cosmos"
				srcKeybase, _ = generateKey("test")
				_, srcKey = generateKey("wrong")
			},
			expPass: false,
		},
	}

	for _, test := range tests {
		test := test

		s.Run(test.name, func() {
			test.malleate()
			_, err := cli.GetIBCAccountLinkPacket(srcKeybase, srcKey, destChainPrefix)

			if !test.expPass {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestGetSourceKeyInfo() {
	val := s.network.Validators[0]
	tests := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			name: "Get key info successfully",
			malleate: func() {
				srcKeybase, srcKey := generateKey("test")
				val.ClientCtx.Keyring = srcKeybase
				val.ClientCtx.FromAddress = srcKey.GetAddress()
				val.ClientCtx.FromName = srcKey.GetName()
			},
			expPass: true,
		},
		{
			name: "Empty keybase",
			malleate: func() {
				val.ClientCtx.Keyring = keyring.NewInMemory()
			},
			expPass: false,
		},
	}

	for _, test := range tests {
		test := test

		s.Run(test.name, func() {
			test.malleate()
			_, _, err := cli.GetSourceKeyInfo(val.ClientCtx)

			if !test.expPass {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
			}
		})
	}
}
