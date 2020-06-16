package models_test

import (
	"encoding/json"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profile/internal/types/models"
	"github.com/stretchr/testify/require"
)

func TestPostQueryResponse_MarshalJSON(t *testing.T) {
	min := sdk.NewInt(5)
	max := sdk.NewInt(10)
	nameSurnameParams := models.NewNameSurnameLenParams(&min, &max)
	monikerParams := models.NewMonikerLenParams(&min, &max)
	bioParams := models.NewBioLenParams(sdk.NewInt(100))

	paramQueryResponse := models.NewParamsQueryResponse(
		nameSurnameParams,
		monikerParams,
		bioParams,
	)

	expResponse := `{"name_surname_len_params":{"min_name_surname_len":"5","max_name_surname_len":"10"},"moniker_len_params":{"min_moniker_len":"5","max_moniker_len":"10"},"bio_len_params":{"max_bio_len":"100"}}`

	jsonData, err := json.Marshal(&paramQueryResponse)
	require.NoError(t, err)
	require.Equal(t, expResponse, string(jsonData))
}
