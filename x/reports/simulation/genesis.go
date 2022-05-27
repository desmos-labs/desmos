package simulation

// DONTCOVER

import (
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	postssim "github.com/desmos-labs/desmos/v3/x/posts/simulation"
	poststypes "github.com/desmos-labs/desmos/v3/x/posts/types"
	"github.com/desmos-labs/desmos/v3/x/reports/types"
	subspacessim "github.com/desmos-labs/desmos/v3/x/subspaces/simulation"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

// RandomizeGenState generates a random GenesisState for subspaces
func RandomizeGenState(simState *module.SimulationState) {
	// Read the subspaces data
	subspacesGenesisBz := simState.GenState[subspacestypes.ModuleName]
	var subspacesGenesis subspacestypes.GenesisState
	simState.Cdc.MustUnmarshalJSON(subspacesGenesisBz, &subspacesGenesis)

	// Read the posts data
	postsGenesisBz := simState.GenState[poststypes.ModuleName]
	var postsGenesis poststypes.GenesisState
	simState.Cdc.MustUnmarshalJSON(postsGenesisBz, &postsGenesis)

	reasons := randomReasons(simState.Rand, subspacesGenesis.Subspaces)
	reports := randomReports(simState.Rand, reasons, subspacesGenesis.Subspaces, simState.Accounts, postsGenesis.GenesisPosts)
	subspacesDataEntries := getSubspacesData(subspacesGenesis.Subspaces, reasons, reports)
	params := types.NewParams(GetRandomStandardReasons(simState.Rand, 10))

	// Create the genesis and sanitize it
	reportsGenesis := types.NewGenesisState(subspacesDataEntries, reasons, reports, params)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(reportsGenesis)
}

// randomReasons return a randomly generated slice of reasons
func randomReasons(r *rand.Rand, subspaces []subspacestypes.GenesisSubspace) []types.Reason {
	reasonNumber := r.Intn(20)
	reasons := make([]types.Reason, reasonNumber)
	for i := 0; i < reasonNumber; i++ {
		// Get a random subspace
		subspace := subspacessim.RandomGenesisSubspace(r, subspaces)

		// Generate a random reason
		reasons[i] = types.NewReason(
			subspace.Subspace.ID,
			uint32(i+1),
			GetRandomReasonTitle(r),
			GetRandomReasonDescription(r),
		)
	}
	return reasons
}

// randomReports returns a randomly generated slice of reports
func randomReports(r *rand.Rand, reasons []types.Reason, subspaces []subspacestypes.GenesisSubspace, accs []simtypes.Account, genPosts []poststypes.GenesisPost) []types.Report {
	reportsNumber := r.Intn(100)
	reports := make([]types.Report, reportsNumber)
	for i := 0; i < reportsNumber; i++ {
		// Get a random subspace
		subspace := subspacessim.RandomGenesisSubspace(r, subspaces)

		// Get a random reason
		reason := RandomReason(r, reasons)

		var data types.ReportData
		if r.Intn(101) < 50 {
			// 50% of a post report
			posts := getSubspacePosts(subspace.Subspace.ID, genPosts)
			if len(posts) == 0 {
				continue
			}
			post := postssim.RandomGenesisPost(r, posts)
			data = types.NewPostData(post.ID)
		} else {
			// 50% of a user report
			account, _ := simtypes.RandomAcc(r, accs)
			data = types.NewUserData(account.Address.String())
		}

		// Get a random reporter
		reporter, _ := simtypes.RandomAcc(r, accs)

		// Generate a random report
		reports[i] = types.NewReport(
			subspace.Subspace.ID,
			uint64(i+1),
			reason.ID,
			GetRandomMessage(r),
			data,
			reporter.Address.String(),
			time.Now(),
		)
	}

	return reports
}

// getSubspacePosts gets all the posts for the given subspace from the provided slice
func getSubspacePosts(subspaceID uint64, genPosts []poststypes.GenesisPost) []poststypes.GenesisPost {
	var posts []poststypes.GenesisPost
	for _, post := range genPosts {
		if post.SubspaceID == subspaceID {
			posts = append(posts, post)
		}
	}
	return posts
}

// getSubspacesData gets the subspaces data for the provided subspaces
func getSubspacesData(subspaces []subspacestypes.GenesisSubspace, reasons []types.Reason, reports []types.Report) []types.SubspaceDataEntry {
	entries := make([]types.SubspaceDataEntry, len(subspaces))
	for i, subspace := range subspaces {
		// Get the max reason id
		maxReasonID := uint32(0)
		for _, reason := range reasons {
			if reason.SubspaceID == subspace.Subspace.ID && reason.ID > maxReasonID {
				maxReasonID = reason.ID
			}
		}

		// Get the max report id
		maxReportID := uint64(0)
		for _, report := range reports {
			if report.SubspaceID == subspace.Subspace.ID && report.ID > maxReportID {
				maxReportID = report.ID
			}
		}

		// Generate the entry
		entries[i] = types.NewSubspacesDataEntry(
			subspace.Subspace.ID,
			maxReasonID+1,
			maxReportID+1,
		)
	}
	return entries
}
