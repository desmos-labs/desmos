package simulation

// DONTCOVER

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/desmos-labs/desmos/x/posts/simulation"

	posts "github.com/desmos-labs/desmos/x/posts/types"
)

var (
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
	PostID  string
	Message string
	Type    string
	Creator simtypes.Account
}

// RandomReportsData returns a randomly generated ReportsData based on the given random posts and accounts list
func RandomReportsData(r *rand.Rand, posts []posts.Post, accs []simtypes.Account) ReportsData {
	post, _ := simulation.RandomPost(r, posts)
	simAccount, _ := simtypes.RandomAcc(r, accs)

	return ReportsData{
		Creator: simAccount,
		PostID:  post.PostID,
		Message: RandomReportMessage(r),
		Type:    RandomReportTypes(r),
	}
}

// RandomPostID returns a randomly generated postID
func RandomPostID(r *rand.Rand) string {
	randBytes := make([]byte, 4)
	_, err := r.Read(randBytes)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(randBytes)
	return hex.EncodeToString(hash[:])
}

func RandomReportMessage(r *rand.Rand) string {
	return messages[r.Intn(len(messages))]
}

func RandomReportTypes(r *rand.Rand) string {
	return repTypes[r.Intn(len(repTypes))]
}
