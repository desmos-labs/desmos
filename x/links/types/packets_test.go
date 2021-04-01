package types_test

import (
	fmt "fmt"
	"testing"

	"github.com/desmos-labs/desmos/x/links/types"
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
				"desmos",
				"desmos1tw3jl54lmwn3mq6hjfvl5nsk4q70v34wc9nsyk",
				"desmospub1addwnpepqfrxkfzkyduxzvfz2em0hn6wkk3jeq663tx8xw5cnt69kr9mes8gg49l8u8",
				"desmos1488h84vd9rc0dmwxx9gzskmymwr7afcemegt9q",
				"",
				"",
			),
			expErr: nil,
		},
	}
	link := types.NewLink("desmos1tw3jl54lmwn3mq6hjfvl5nsk4q70v34wc9nsyk", "desmos1488h84vd9rc0dmwxx9gzskmymwr7afcemegt9q")
	linkBytes, _ := link.Marshal()
	fmt.Printf("%b", linkBytes)
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, test.packet.Validate())
		})
	}
}
