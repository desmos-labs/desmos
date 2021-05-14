package types_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/desmos-labs/desmos/x/ibc/profiles/types"
	"github.com/stretchr/testify/require"
)

func TestIBCAccountConnectionPacketData_Validate(t *testing.T) {
	tests := []struct {
		name   string
		packet types.IBCAccountConnectionPacketData
		expErr error
	}{
		{
			name: "Valid IBCAccountConnectionPacketData",
			packet: types.NewIBCAccountConnectionPacketData(
				"cosmos",
				"cosmos1u33w3u4ler4654phrpt6xqvh92ch0v6mcjrj97",
				"02bcd0738e3b7e0f6650c8e6eb10bd4266fd6818c92a0283b4cb0884f046051c3e",
				"cosmos1qdasq0mzpajknaaj32kf9lk5nmcy8g65mddd4p",
				"143bbd3131d76232f973f84a9ea7be751243044315056dbecb968942da97e474401caa1c8f2c4ce5e48052cf44066717f2166c21a7277de9911d75c57eca598d",
				"1a18d5f012ce0e8258fd3455c01b48249bb019231e416c4323ab2bb170b4ad0951b370138d2ea69a376feb942d3c619c9152d63a6d2e0232aaff77162df66636",
			),
			expErr: nil,
		},
		{
			name: "Source Pubkey and address are mismatched",
			packet: types.NewIBCAccountConnectionPacketData(
				"cosmos",
				"cosmos1u33w3u4ler4654phrpt6xqvh92ch0v6mcjrj97",
				"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
				"cosmos1qdasq0mzpajknaaj32kf9lk5nmcy8g65mddd4p",
				"143bbd3131d76232f973f84a9ea7be751243044315056dbecb968942da97e474401caa1c8f2c4ce5e48052cf44066717f2166c21a7277de9911d75c57eca598d",
				"1a18d5f012ce0e8258fd3455c01b48249bb019231e416c4323ab2bb170b4ad0951b370138d2ea69a376feb942d3c619c9152d63a6d2e0232aaff77162df66636",
			),
			expErr: fmt.Errorf("source pubkey and source address are mismatched"),
		},
		{
			name: "Empty source prefix",
			packet: types.NewIBCAccountConnectionPacketData(
				"",
				"cosmos1u33w3u4ler4654phrpt6xqvh92ch0v6mcjrj97",
				"02bcd0738e3b7e0f6650c8e6eb10bd4266fd6818c92a0283b4cb0884f046051c3e",
				"cosmos1qdasq0mzpajknaaj32kf9lk5nmcy8g65mddd4p",
				"143bbd3131d76232f973f84a9ea7be751243044315056dbecb968942da97e474401caa1c8f2c4ce5e48052cf44066717f2166c21a7277de9911d75c57eca598d",
				"1a18d5f012ce0e8258fd3455c01b48249bb019231e416c4323ab2bb170b4ad0951b370138d2ea69a376feb942d3c619c9152d63a6d2e0232aaff77162df66636",
			),
			expErr: fmt.Errorf("chain prefix cannot be empty"),
		},
		{
			name: "Invalid format source address",
			packet: types.NewIBCAccountConnectionPacketData(
				"cosmos",
				"cosmos1u33w3u4ler4654phrpt6xqvh92ch0v6mcjrj",
				"02bcd0738e3b7e0f6650c8e6eb10bd4266fd6818c92a0283b4cb0884f046051c3e",
				"cosmos1qdasq0mzpajknaaj32kf9lk5nmcy8g65mddd4p",
				"143bbd3131d76232f973f84a9ea7be751243044315056dbecb968942da97e474401caa1c8f2c4ce5e48052cf44066717f2166c21a7277de9911d75c57eca598d",
				"1a18d5f012ce0e8258fd3455c01b48249bb019231e416c4323ab2bb170b4ad0951b370138d2ea69a376feb942d3c619c9152d63a6d2e0232aaff77162df66636",
			),
			expErr: fmt.Errorf("failed to parse source address"),
		},
		{
			name: "Invalid format of source pubkey",
			packet: types.NewIBCAccountConnectionPacketData(
				"cosmos",
				"cosmos1u33w3u4ler4654phrpt6xqvh92ch0v6mcjrj97",
				"02bcd0738e3b7e0f6650c8e6eb10bd4266fd6818c92a0283b4cb0884f046051",
				"cosmos1qdasq0mzpajknaaj32kf9lk5nmcy8g65mddd4p",
				"143bbd3131d76232f973f84a9ea7be751243044315056dbecb968942da97e474401caa1c8f2c4ce5e48052cf44066717f2166c21a7277de9911d75c57eca598d",
				"1a18d5f012ce0e8258fd3455c01b48249bb019231e416c4323ab2bb170b4ad0951b370138d2ea69a376feb942d3c619c9152d63a6d2e0232aaff77162df66636",
			),
			expErr: fmt.Errorf("failed to decode source pubkey"),
		},
		{
			name: "Invalid format of destination address",
			packet: types.NewIBCAccountConnectionPacketData(
				"cosmos",
				"cosmos1u33w3u4ler4654phrpt6xqvh92ch0v6mcjrj97",
				"02bcd0738e3b7e0f6650c8e6eb10bd4266fd6818c92a0283b4cb0884f046051c3e",
				"cosmos1qdasq0mzpajknaaj32kf9lk5nmcy8g65mddd",
				"143bbd3131d76232f973f84a9ea7be751243044315056dbecb968942da97e474401caa1c8f2c4ce5e48052cf44066717f2166c21a7277de9911d75c57eca598d",
				"1a18d5f012ce0e8258fd3455c01b48249bb019231e416c4323ab2bb170b4ad0951b370138d2ea69a376feb942d3c619c9152d63a6d2e0232aaff77162df66636",
			),
			expErr: fmt.Errorf("failed to parse destination address"),
		},
		{
			name: "Invalid pubkey for signature",
			packet: types.NewIBCAccountConnectionPacketData(
				"cosmos",
				"cosmos1u33w3u4ler4654phrpt6xqvh92ch0v6mcjrj97",
				"02bcd0738e3b7e0f6650c8e6eb10bd4266fd6818c92a0283b4cb0884f046051c3e",
				"cosmos1qdasq0mzpajknaaj32kf9lk5nmcy8g65mddd4p",
				"1a18d5f012ce0e8258fd3455c01b48249bb019231e416c4323ab2bb170b4ad0951b370138d2ea69a376feb942d3c619c9152d63a6d2e0232aaff77162df66636",
				"1a18d5f012ce0e8258fd3455c01b48249bb019231e416c4323ab2bb170b4ad0951b370138d2ea69a376feb942d3c619c9152d63a6d2e0232aaff77162df66636",
			),
			expErr: fmt.Errorf("failed to verify source signature"),
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, test.packet.Validate())
		})
	}
}

func TestIBCAccountConnectionPacketData_GetBytes(t *testing.T) {
	p := types.NewIBCAccountConnectionPacketData(
		"cosmos",
		"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
		"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
		"cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn",
		"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
		"fc0bc7dd041c736b8fa3bb6638fc003944b430aaa656d08b823836894338d30d5bb8c96e43d4c40d820acf2f6d03c8123df525c59eed114564b877ed1f7dd561",
	)
	_, err := p.GetBytes()
	require.NoError(t, err)
}

func TestIBCAccountLinkPacketData_Validate(t *testing.T) {
	tests := []struct {
		name   string
		packet types.IBCAccountLinkPacketData
		expErr error
	}{
		{
			name: "Valid IBCAccountConnectionPacketData",
			packet: types.NewIBCAccountLinkPacketData(
				"desmos",
				"desmos1tw3jl54lmwn3mq6hjfvl5nsk4q70v34wc9nsyk",
				"02466b245623786131225676fbcf4eb5a32c835a8acc733a989af45b0cbbcc0e84",
				"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
			),
			expErr: nil,
		},
		{
			name: "Source Pubkey and address are mismatched",
			packet: types.NewIBCAccountLinkPacketData(
				"desmos",
				"desmos1tw3jl54lmwn3mq6hjfvl5nsk4q70v34wc9nsyk",
				"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
				"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
			),
			expErr: fmt.Errorf("source pubkey and source address are mismatched"),
		},
		{
			name: "Empty source prefix",
			packet: types.NewIBCAccountLinkPacketData(
				"",
				"desmos1tw3jl54lmwn3mq6hjfvl5nsk4q70v34wc9nsyk",
				"02466b245623786131225676fbcf4eb5a32c835a8acc733a989af45b0cbbcc0e84",
				"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
			),
			expErr: fmt.Errorf("chain prefix cannot be empty"),
		},
		{
			name: "Invalid source address",
			packet: types.NewIBCAccountLinkPacketData(
				"desmos",
				"desmos1tw3jl54lmwn3mq6hjfvl5nsk4q70v34wc9nsy",
				"02466b245623786131225676fbcf4eb5a32c835a8acc733a989af45b0cbbcc0e84",
				"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
			),
			expErr: fmt.Errorf("failed to parse source address"),
		},
		{
			name: "Invalid source pubkey",
			packet: types.NewIBCAccountLinkPacketData(
				"desmos",
				"desmos1tw3jl54lmwn3mq6hjfvl5nsk4q70v34wc9nsyk",
				"02466b245623786131225676fbcf4eb5a32c835a8acc733a989af45b0cbbcc0e8",
				"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
			),
			expErr: fmt.Errorf("failed to decode source pubkey"),
		},
		{
			name: "Invalid source signature",
			packet: types.NewIBCAccountLinkPacketData(
				"desmos",
				"desmos1tw3jl54lmwn3mq6hjfvl5nsk4q70v34wc9nsyk",
				"02466b245623786131225676fbcf4eb5a32c835a8acc733a989af45b0cbbcc0e84",
				"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b",
			),
			expErr: fmt.Errorf("failed to decode source signature"),
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, test.packet.Validate())
		})
	}
}

func TestIBCAccountLinkPacketData_GetBytes(t *testing.T) {
	p := types.NewIBCAccountLinkPacketData(
		"desmos",
		"desmos1tw3jl54lmwn3mq6hjfvl5nsk4q70v34wc9nsyk",
		"02466b245623786131225676fbcf4eb5a32c835a8acc733a989af45b0cbbcc0e84",
		"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
	)
	_, err := p.GetBytes()
	require.NoError(t, err)
}

func Test_VerifySignature(t *testing.T) {
	tests := []struct {
		name   string
		msg    string // Hex-encoded msg
		sig    string // Hex-encoded sig
		pubKey string // Hex-encoded pubKey
		expRes bool
	}{
		{
			name:   "Valid signature",
			msg:    "28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
			sig:    "fc0bc7dd041c736b8fa3bb6638fc003944b430aaa656d08b823836894338d30d5bb8c96e43d4c40d820acf2f6d03c8123df525c59eed114564b877ed1f7dd561",
			pubKey: "02b493a33f104de068e93d51ffe9929409a20635a68d0c2bc2b51d95e186e58f07",
			expRes: true,
		},
		{
			name:   "Invalid signature",
			msg:    "28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
			sig:    "fc0bc7dd041c736b8fa3bb6638fc003944b430aaa656d08b823836894338d30d5bb8c96e43d4c40d820acf2f6d03c8123df525c59eed114564b877ed1f7dd561",
			pubKey: "02466b245623786131225676fbcf4eb5a32c835a8acc733a989af45b0cbbcc0e84",
			expRes: false,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			msg, _ := hex.DecodeString(test.msg)
			sig, _ := hex.DecodeString(test.sig)
			pubKeyBs, _ := hex.DecodeString(test.pubKey)
			pubKey := &secp256k1.PubKey{Key: pubKeyBs}
			require.Equal(t, test.expRes, types.VerifySignature(msg, sig, pubKey))
		})
	}
}
