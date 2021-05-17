package types_test

import (
	"fmt"
	"testing"

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
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
				"cosmos13rzf5gph4drs3qnf63jmuyf4g9q7a4cv9n0uqq",
				"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
				"42dd1f8d98c5de91a12259cf46098104132f69b61eaa24e112bf504d17e1a0b71274dad981bbb4a13dc440905a19be92eaf4497940751f431c530cc4d68e78b0",
			),
			expErr: nil,
		},
		{
			name: "Source Pubkey and address are mismatched",
			packet: types.NewIBCAccountConnectionPacketData(
				"cosmos",
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
				"cosmos13rzf5gph4drs3qnf63jmuyf4g9q7a4cv9n0uqq",
				"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
				"42dd1f8d98c5de91a12259cf46098104132f69b61eaa24e112bf504d17e1a0b71274dad981bbb4a13dc440905a19be92eaf4497940751f431c530cc4d68e78b0",
			),
			expErr: fmt.Errorf("source pubkey and source address are mismatched"),
		},
		{
			name: "Empty source prefix",
			packet: types.NewIBCAccountConnectionPacketData(
				"",
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
				"cosmos13rzf5gph4drs3qnf63jmuyf4g9q7a4cv9n0uqq",
				"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
				"42dd1f8d98c5de91a12259cf46098104132f69b61eaa24e112bf504d17e1a0b71274dad981bbb4a13dc440905a19be92eaf4497940751f431c530cc4d68e78b0",
			),
			expErr: fmt.Errorf("chain prefix cannot be empty"),
		},
		{
			name: "Invalid format source address",
			packet: types.NewIBCAccountConnectionPacketData(
				"cosmos",
				"=",
				"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
				"cosmos13rzf5gph4drs3qnf63jmuyf4g9q7a4cv9n0uqq",
				"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
				"42dd1f8d98c5de91a12259cf46098104132f69b61eaa24e112bf504d17e1a0b71274dad981bbb4a13dc440905a19be92eaf4497940751f431c530cc4d68e78b0",
			),
			expErr: fmt.Errorf("failed to parse source address"),
		},
		{
			name: "Invalid format of source pubkey",
			packet: types.NewIBCAccountConnectionPacketData(
				"cosmos",
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				"=",
				"cosmos13rzf5gph4drs3qnf63jmuyf4g9q7a4cv9n0uqq",
				"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
				"42dd1f8d98c5de91a12259cf46098104132f69b61eaa24e112bf504d17e1a0b71274dad981bbb4a13dc440905a19be92eaf4497940751f431c530cc4d68e78b0",
			),
			expErr: fmt.Errorf("failed to decode source pubkey"),
		},
		{
			name: "Invalid format of destination address",
			packet: types.NewIBCAccountConnectionPacketData(
				"cosmos",
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
				"=",
				"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
				"42dd1f8d98c5de91a12259cf46098104132f69b61eaa24e112bf504d17e1a0b71274dad981bbb4a13dc440905a19be92eaf4497940751f431c530cc4d68e78b0",
			),
			expErr: fmt.Errorf("failed to parse destination address"),
		},
		{
			name: "Invalid pubkey for signature",
			packet: types.NewIBCAccountConnectionPacketData(
				"cosmos",
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
				"cosmos13rzf5gph4drs3qnf63jmuyf4g9q7a4cv9n0uqq",
				"42dd1f8d98c5de91a12259cf46098104132f69b61eaa24e112bf504d17e1a0b71274dad981bbb4a13dc440905a19be92eaf4497940751f431c530cc4d68e78b0",
				"42dd1f8d98c5de91a12259cf46098104132f69b61eaa24e112bf504d17e1a0b71274dad981bbb4a13dc440905a19be92eaf4497940751f431c530cc4d68e78b0",
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
		"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
		"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
		"cosmos13rzf5gph4drs3qnf63jmuyf4g9q7a4cv9n0uqq",
		"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
		"42dd1f8d98c5de91a12259cf46098104132f69b61eaa24e112bf504d17e1a0b71274dad981bbb4a13dc440905a19be92eaf4497940751f431c530cc4d68e78b0",
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
				"cosmos",
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
				"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
			),
			expErr: nil,
		},
		{
			name: "Source Pubkey and address are mismatched",
			packet: types.NewIBCAccountLinkPacketData(
				"cosmos",
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				"02466b245623786131225676fbcf4eb5a32c835a8acc733a989af45b0cbbcc0e84",
				"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
			),
			expErr: fmt.Errorf("source pubkey and source address are mismatched"),
		},
		{
			name: "Empty source prefix",
			packet: types.NewIBCAccountLinkPacketData(
				"",
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
				"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
			),
			expErr: fmt.Errorf("chain prefix cannot be empty"),
		},
		{
			name: "Invalid source address",
			packet: types.NewIBCAccountLinkPacketData(
				"cosmos",
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs",
				"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
				"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
			),
			expErr: fmt.Errorf("failed to parse source address"),
		},
		{
			name: "Invalid hex string of source pubkey",
			packet: types.NewIBCAccountLinkPacketData(
				"cosmos",
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				"=",
				"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
			),
			expErr: fmt.Errorf("failed to decode source pubkey"),
		},
		{
			name: "Invalid hex string of source signature",
			packet: types.NewIBCAccountLinkPacketData(
				"cosmos",
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
				"=",
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
		"cosmos",
		"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
		"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
		"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
	)
	_, err := p.GetBytes()
	require.NoError(t, err)
}
