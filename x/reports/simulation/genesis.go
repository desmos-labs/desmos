package simulation

// DONTCOVER

import (
	"math/rand"
	"time"

	relationshipstypes "github.com/desmos-labs/desmos/v6/x/relationships/types"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	postssim "github.com/desmos-labs/desmos/v6/x/posts/simulation"
	poststypes "github.com/desmos-labs/desmos/v6/x/posts/types"
	"github.com/desmos-labs/desmos/v6/x/reports/types"
	subspacessim "github.com/desmos-labs/desmos/v6/x/subspaces/simulation"
	subspacestypes "github.com/desmos-labs/desmos/v6/x/subspaces/types"
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

	// Read the relationships data
	relationshipsGenesisBz := simState.GenState[relationshipstypes.ModuleName]
	var relationshipsGenesis relationshipstypes.GenesisState
	simState.Cdc.MustUnmarshalJSON(relationshipsGenesisBz, &relationshipsGenesis)

	reasons := randomReasons(simState.Rand, subspacesGenesis.Subspaces)
	reports := randomReports(simState.Rand, simState.Accounts, subspacesGenesis.Subspaces, relationshipsGenesis.Blocks, postsGenesis.Posts, reasons)
	subspacesDataEntries := getSubspacesData(subspacesGenesis.Subspaces, reasons, reports)
	params := types.NewParams(GetRandomStandardReasons(simState.Rand, 10))

	// Create the genesis and sanitize it
	reportsGenesis := types.NewGenesisState(subspacesDataEntries, reasons, reports, params)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(reportsGenesis)
}

// randomReasons return a randomly generated slice of reasons
func randomReasons(r *rand.Rand, subspaces []subspacestypes.Subspace) []types.Reason {
	if len(subspaces) == 0 {
		return nil
	}

	reasonNumber := r.Intn(20)
	reasons := make([]types.Reason, reasonNumber)
	for i := 0; i < reasonNumber; i++ {
		// Get a random subspace
		subspace := subspacessim.RandomSubspace(r, subspaces)

		// Generate a random reason
		reasons[i] = types.NewReason(
			subspace.ID,
			uint32(i+1),
			GetRandomReasonTitle(r),
			GetRandomReasonDescription(r),
		)
	}
	return reasons
}

// randomReports returns a randomly generated slice of reports
func randomReports(r *rand.Rand, accs []simtypes.Account, subspaces []subspacestypes.Subspace, blocks []relationshipstypes.UserBlock, genPosts []poststypes.Post, reasons []types.Reason) []types.Report {
	if len(subspaces) == 0 || len(reasons) == 0 {
		// No subspaces or valid reasons, so no way we can have a valid post
		return nil
	}

	reportsNumber := r.Intn(100)
	var reports []types.Report
	for i := 0; i < reportsNumber; i++ {
		// Get a random subspace
		subspace := subspacessim.RandomSubspace(r, subspaces)

		// Get a random reporter
		reporter, _ := simtypes.RandomAcc(r, accs)

		// Get a random reason
		subspaceReasons := getSubspaceReasons(subspace.ID, reasons)
		if len(subspaceReasons) == 0 {
			continue
		}

		reason := RandomReason(r, subspaceReasons)

		var data types.ReportTarget
		if r.Intn(101) < 50 {
			// 50% of a post report
			posts := getSubspacePosts(subspace.ID, genPosts)
			if len(posts) == 0 {
				continue
			}
			post := postssim.RandomPost(r, posts)
			if isUserBlocked(reporter.Address.String(), post.Owner, subspace.ID, blocks) {
				continue
			}
			data = types.NewPostTarget(post.ID)
		} else {
			// 50% of a user report
			account, _ := simtypes.RandomAcc(r, accs)
			if isUserBlocked(reporter.Address.String(), account.Address.String(), subspace.ID, blocks) {
				continue
			}
			data = types.NewUserTarget(account.Address.String())
		}

		// Generate a random report
		reports = append(reports, types.NewReport(
			subspace.ID,
			uint64(i+1),
			[]uint32{reason.ID},
			GetRandomMessage(r),
			data,
			reporter.Address.String(),
			time.Now(),
		))
	}

	return reports
}

// getSubspaceReasons returns the reporting reasons for the given subspace filtering the given reasons slice
func getSubspaceReasons(subspaceID uint64, reasons []types.Reason) []types.Reason {
	var subspaceReasons []types.Reason
	for _, reason := range reasons {
		if reason.SubspaceID == subspaceID {
			subspaceReasons = append(subspaceReasons, reason)
		}
	}
	return subspaceReasons
}

// isUserBlocked checks if the given user is blocked by the blocker on the provided subspace checking within the blocks
func isUserBlocked(user string, blocker string, subspaceID uint64, blocks []relationshipstypes.UserBlock) bool {
	for _, block := range blocks {
		if block.Blocked == user && block.Blocker == blocker && block.SubspaceID == subspaceID {
			return true
		}
	}
	return false
}

// getSubspacePosts gets all the posts for the given subspace from the provided slice
func getSubspacePosts(subspaceID uint64, genPosts []poststypes.Post) []poststypes.Post {
	var posts []poststypes.Post
	for _, post := range genPosts {
		if post.SubspaceID == subspaceID {
			posts = append(posts, post)
		}
	}
	return posts
}

// getSubspacesData gets the subspaces data for the provided subspaces
func getSubspacesData(subspaces []subspacestypes.Subspace, reasons []types.Reason, reports []types.Report) []types.SubspaceDataEntry {
	entries := make([]types.SubspaceDataEntry, len(subspaces))
	for i, subspace := range subspaces {
		// Get the max reason id
		maxReasonID := uint32(0)
		for _, reason := range reasons {
			if reason.SubspaceID == subspace.ID && reason.ID > maxReasonID {
				maxReasonID = reason.ID
			}
		}

		// Get the max report id
		maxReportID := uint64(0)
		for _, report := range reports {
			if report.SubspaceID == subspace.ID && report.ID > maxReportID {
				maxReportID = report.ID
			}
		}

		// Generate the entry
		entries[i] = types.NewSubspacesDataEntry(
			subspace.ID,
			maxReasonID+1,
			maxReportID+1,
		)
	}
	return entries
}
