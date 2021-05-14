package cli_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/cosmos/go-bip39"
	"github.com/stretchr/testify/suite"

	"github.com/desmos-labs/desmos/testutil"
	"github.com/desmos-labs/desmos/x/ibc/profiles/client/cli"
	"github.com/desmos-labs/desmos/x/ibc/profiles/types"
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

	var gs types.GenesisState
	s.Require().NoError(cfg.Codec.UnmarshalJSON(genesisState[types.ModuleName], &gs))

	gs.PortId = "ibc-profiles"

	gsBz, err := cfg.Codec.MarshalJSON(&gs)
	s.Require().NoError(err)
	genesisState[types.ModuleName] = gsBz
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

func TestGetQueryCmd(t *testing.T) {
	cmd := cli.GetQueryCmd("link")
	if cmd == nil {
		t.Errorf("get no query command")
	}
}

func TestNewTxCmd(t *testing.T) {
	cmd := cli.NewTxCmd()
	if cmd == nil {
		t.Errorf("Failed to get tx command")
	}
}

func (s *IntegrationTestSuite) TestGetCmdCreateIBCAccountConnection() {
	ctx := s.network.Validators[0].ClientCtx

	tests := []struct {
		name     string
		args     []string
		malleate func()
		expPass  bool
		expErr   error
	}{
		{
			name: "Empty keybase",
			args: []string{
				"ibcprofiles",
				"channel-0",
				"desmos",
				".",
				"test",
				fmt.Sprintf("--%s=%s", flags.FlagKeyringBackend, keyring.BackendMemory),
			},
			malleate: func() {
				ctx.Keyring = keyring.NewInMemory()
			},
			expPass: false,
			expErr:  fmt.Errorf("The specified item could not be found in the keyring"),
		},
		{
			name: "Invalid destination keybase",
			args: []string{
				"ibcprofiles",
				"channel-0",
				"desmos",
				".",
				"test",
			},
			malleate: func() {
				keybase, _ := generateMemoryKeybase("could not get destination key")
				ctx.Keyring = keybase
			},
			expPass: false,
			expErr:  fmt.Errorf("The specified item could not be found in the keyring"),
		},
		{
			name: "Wrong destination key name for destination keybase",
			args: []string{
				"ibcprofiles",
				"channel-0",
				"desmos",
				"",
				"wrongname",
				fmt.Sprintf("--%s=%s", cli.FlagTesting, "true"),
			},
			malleate: func() {
				keybase, _ := generateMemoryKeybase("")
				ctx.Keyring = keybase
			},
			expPass: false,
			expErr:  fmt.Errorf("could not get destination key"),
		},
		{
			name: "Channel is not available",
			args: []string{
				"ibcprofiles",
				"channel-0",
				"desmos",
				".",
				"test",
				fmt.Sprintf("--%s=%s", cli.FlagTesting, "true"),
			},
			malleate: func() {
				keybase, _ := generateMemoryKeybase("")
				ctx.Keyring = keybase
			},
			expPass: false,
		},
		{
			name: "Invalid args number",
			args: []string{
				"ibcprofiles",
				"channel-0",
				"desmos",
				".",
				"test",
				"hi123",
			},
			malleate: func() {},
			expPass:  false,
			expErr:   fmt.Errorf("accepts 5 arg(s), received 6"),
		},
	}

	for _, test := range tests {
		test := test

		s.Run(test.name, func() {
			test.malleate()
			cmd := cli.GetCmdCreateIBCAccountConnection()
			_, err := clitestutil.ExecTestCLICmd(ctx, cmd, test.args)

			if !test.expPass {
				s.Require().Error(err)
				if test.expErr != nil {
					s.Require().Equal(test.expErr, err)
				}
			} else {
				s.Require().NoError(err)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestGetCmdCreateIBCAccountLink() {
	ctx := s.network.Validators[0].ClientCtx

	tests := []struct {
		name     string
		args     []string
		malleate func()
		expPass  bool
		expErr   error
	}{
		{
			name: "Empty keybase",
			args: []string{
				"ibcprofiles",
				"channel-0",
				"desmos",
				fmt.Sprintf("--%s=%s", flags.FlagKeyringBackend, keyring.BackendMemory),
			},
			malleate: func() {
				ctx.Keyring = keyring.NewInMemory()
			},
			expPass: false,
			expErr:  fmt.Errorf("The specified item could not be found in the keyring"),
		},
		{
			name: "Channel is not available",
			args: []string{
				"ibcprofiles",
				"channel-0",
				"desmos",
			},
			malleate: func() {
				keybase, _ := generateMemoryKeybase("")
				ctx.Keyring = keybase
			},
			expPass: false,
		},
		{
			name: "Invalid args number",
			args: []string{
				"ibcprofiles",
				"channel-0",
				"desmos",
				"456",
			},
			malleate: func() {
			},
			expPass: false,
			expErr:  fmt.Errorf("accepts 3 arg(s), received 4"),
		},
	}

	for _, test := range tests {
		test := test

		s.Run(test.name, func() {
			test.malleate()
			cmd := cli.GetCmdCreateIBCAccountLink()
			_, err := clitestutil.ExecTestCLICmd(ctx, cmd, test.args)

			if !test.expPass {
				s.Require().Error(err)
				if test.expErr != nil {
					s.Require().Equal(test.expErr, err)
				}
			} else {
				s.Require().NoError(err)
			}
		})
	}
}

// ___________________________________________________________________________________________________________________

func generateMemoryKeybase(keyname string) (keyring.Keyring, keyring.Info) {
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
		expErr   error
	}{
		{
			name: "Get packet successfully",
			malleate: func() {
				destChainPrefix = "cosmos"
				srcKeybase, srcKey = generateMemoryKeybase("test")
				destKeybase, destKey = generateMemoryKeybase("test")
			},
			expPass: true,
		},
		{
			name: "Wrong source key name",
			malleate: func() {
				destChainPrefix = "cosmos"
				srcKeybase, _ = generateMemoryKeybase("test")
				_, srcKey = generateMemoryKeybase("wrong")
				destKeybase, destKey = generateMemoryKeybase("test")
			},
			expPass: false,
			expErr:  fmt.Errorf("The specified item could not be found in the keyring"),
		},
		{
			name: "Wrong dest key name",
			malleate: func() {
				destChainPrefix = "cosmos"
				srcKeybase, srcKey = generateMemoryKeybase("test")
				destKeybase, _ = generateMemoryKeybase("test")
				_, destKey = generateMemoryKeybase("wrong")
			},
			expPass: false,
			expErr:  fmt.Errorf("The specified item could not be found in the keyring"),
		},
	}

	for _, test := range tests {
		test := test

		s.Run(test.name, func() {
			test.malleate()
			_, err := cli.GetIBCAccountConnectionPacket(srcKeybase, srcKey, destKeybase, destKey, destChainPrefix)

			if !test.expPass {
				s.Require().Error(err)
				if test.expErr != nil {
					s.Require().Equal(test.expErr, err)
				}
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
		expErr   error
	}{
		{
			name: "Get packet successfully",
			malleate: func() {
				destChainPrefix = "cosmos"
				srcKeybase, srcKey = generateMemoryKeybase("test")
			},
			expPass: true,
		},
		{
			name: "Wrong source key name",
			malleate: func() {
				destChainPrefix = "cosmos"
				srcKeybase, _ = generateMemoryKeybase("test")
				_, srcKey = generateMemoryKeybase("wrong")
			},
			expPass: false,
			expErr:  fmt.Errorf("The specified item could not be found in the keyring"),
		},
	}

	for _, test := range tests {
		test := test

		s.Run(test.name, func() {
			test.malleate()
			_, err := cli.GetIBCAccountLinkPacket(srcKeybase, srcKey, destChainPrefix)

			if !test.expPass {
				s.Require().Error(err)
				if test.expErr != nil {
					s.Require().Equal(test.expErr, err)
				}
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
		expErr   error
	}{
		{
			name: "Get key info successfully",
			malleate: func() {
				srcKeybase, srcKey := generateMemoryKeybase("test")
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
			expErr:  fmt.Errorf("The specified item could not be found in the keyring"),
		},
	}

	for _, test := range tests {
		test := test

		s.Run(test.name, func() {
			test.malleate()
			_, _, err := cli.GetSourceKeyInfo(val.ClientCtx)

			if !test.expPass {
				s.Require().Error(err)
				if test.expErr != nil {
					s.Require().Equal(test.expErr, err)
				}
			} else {
				s.Require().NoError(err)
			}
		})
	}
}
