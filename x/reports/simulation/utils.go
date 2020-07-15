package simulation

// DONTCOVER

import (
	"math/rand"
	"time"

	sim "github.com/cosmos/cosmos-sdk/x/simulation"

	posts "github.com/desmos-labs/desmos/x/posts/types"
)

var (
	subspaces = []string{
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		"2bdf5932925584b9a86470bea60adce69041608a447f84a3317723aa5678ec88",
		"3d59f7548e1af2151b64135003ce63c0a484c26b9b8b166a7b1c1805ec34b00a",
		"ec8202b6f9fb16f9e26b66367afa4e037752f3c09a18cefab426165e06a424b1",
		"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
		"3f40462915a3e6026a4d790127b95ded4d870f6ab18d9af2fcbc454168255237",
	}

	messages = []string{
		"it's a trap",
		"it's an offense",
		"it's scam",
		"it'' racism",
	}

	repTypes = []string{
		"nudity",
		"violence",
		"intimidation",
		"suicide or self-harm",
		"fake news",
		"spam",
		"unauthorized sale",
		"hatred incitement",
		"promotion of drug use",
		"non-consensual intimate images",
		"pornography",
		"children abuse",
		"animals abuse",
		"bullying",
		"scam",
	}
)

type ReportsData struct {
	Creator sim.Account
	PostID  posts.PostID
	Message string
	Type    string
}

// RandomReportsData returns a randomly generated ReportsData based on the given random and accounts list
func RandomReportsData(r *rand.Rand, accs []sim.Account) ReportsData {
	simAccount, _ := sim.RandomAcc(r, accs)
	return ReportsData{
		Creator: simAccount,
		PostID:  RandomPostID(r, accs),
		Message: RandomReportMessage(r),
		Type:    RandomReportTypes(r),
	}
}

func RandomPostID(r *rand.Rand, accs []sim.Account) posts.PostID {
	return posts.ComputeID(time.Now(), accs[r.Intn(len(accs))].Address, subspaces[r.Intn(len(subspaces))])
}

func RandomReportMessage(r *rand.Rand) string {
	return messages[r.Intn(len(messages))]
}

func RandomReportTypes(r *rand.Rand) string {
	return repTypes[r.Intn(len(repTypes))]
}
