package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/desmos-labs/desmos/v4/app"
	profilescliutils "github.com/desmos-labs/desmos/v4/x/profiles/client/utils"
	profilestypes "github.com/desmos-labs/desmos/v4/x/profiles/types"
)

func main() {
	dataFilePath := os.Args[1]
	chainLinkJSON, err := os.ReadFile(dataFilePath)
	if err != nil {
		panic(err)
	}

	cdc, legacyCdc := app.MakeCodecs()

	var link profilescliutils.ChainLinkJSON
	if err := cdc.UnmarshalJSON(chainLinkJSON, &link); err != nil {
		panic(err)
	}

	var addrData profilestypes.AddressData
	err = cdc.UnpackAny(link.Address, &addrData)
	if err != nil {
		panic(err)
	}

	err = link.Proof.Verify(cdc, legacyCdc, "desmos1n8345tvzkg3jumkm859r2qz0v6xsc3henzddcj", addrData)
	if err != nil {
		panic(err)
	}

	bz, err := link.Proof.Signature.Marshal()
	if err != nil {
		panic(err)
	}

	fmt.Println(hex.EncodeToString(bz))
}
