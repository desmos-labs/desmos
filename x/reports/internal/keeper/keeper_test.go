package keeper_test

import (
	"github.com/desmos-labs/desmos/x/posts"
	"testing"

	"github.com/desmos-labs/desmos/x/reports/internal/types"
	"github.com/desmos-labs/desmos/x/reports/internal/types/models"
	"github.com/stretchr/testify/require"
)

func TestKeeper_CheckExistence(t *testing.T) {
	existentPost := posts.NewPost(postID,
		"",
		"Post",
		false,
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		map[string]string{},
		testPostCreationDate,
		creator,
	)

	tests := []struct {
		name         string
		existentPost *posts.Post
		postID       posts.PostID
		expBool      bool
	}{
		{
			name:         "Post not exist",
			existentPost: nil,
			postID:       postID,
			expBool:      false,
		},
		{
			name:         "Post exist",
			existentPost: &existentPost,
			postID:       postID,
			expBool:      true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k, pk := SetupTestInput()
			if test.existentPost != nil {
				pk.SavePost(ctx, *test.existentPost)
			}

			actualBool := k.CheckPostExistence(ctx, postID)
			require.Equal(t, test.expBool, actualBool)
		})
	}
}

func TestKeeper_SaveReport(t *testing.T) {
	expReports := models.Reports{models.NewReport("type", "message", creator)}
	report := models.NewReport("type", "message", creator)

	ctx, k, _ := SetupTestInput()
	store := ctx.KVStore(k.StoreKey)

	k.SaveReport(ctx, postID, report)

	var reports models.Reports
	k.Cdc.MustUnmarshalBinaryBare(store.Get(types.ReportStoreKey(postID)), &reports)
	require.Equal(t, expReports, reports)

}

func TestKeeper_GetPostReports(t *testing.T) {
	tests := []struct {
		name       string
		expReports models.Reports
	}{
		{
			name: "Returns a non-empty reports array",
			expReports: models.Reports{
				{Type: "type", Message: "message", User: creator},
			},
		},
		{
			name:       "Returns an empty reports array",
			expReports: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k, _ := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)
			if test.expReports != nil {
				store.Set(types.ReportStoreKey(postID), k.Cdc.MustMarshalBinaryBare(&test.expReports))
			}

			actualRep := k.GetPostReports(ctx, postID)
			require.Equal(t, test.expReports, actualRep)
		})
	}
}

func TestKeeper_GetReportsMap(t *testing.T) {
	reports := models.Reports{
		{Type: "type", Message: "message", User: creator},
	}
	tests := []struct {
		name            string
		existingReports models.Reports
		expReportsMap   map[string]models.Reports
	}{
		{
			name:            "Returns a non-empty reports map",
			existingReports: reports,
			expReportsMap: map[string]models.Reports{
				postID.String(): reports,
			},
		},
		{
			name:            "Returns an empty reports map",
			existingReports: nil,
			expReportsMap:   map[string]models.Reports{},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k, _ := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)
			if test.existingReports != nil {
				store.Set(types.ReportStoreKey(postID), k.Cdc.MustMarshalBinaryBare(&test.existingReports))
			}

			actualRep := k.GetReportsMap(ctx)
			require.Equal(t, test.expReportsMap, actualRep)
		})
	}
}
