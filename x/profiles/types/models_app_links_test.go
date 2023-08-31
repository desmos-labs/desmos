package types_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v6/app"
	"github.com/desmos-labs/desmos/v6/x/profiles/types"
)

func TestApplicationLink_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		link      types.ApplicationLink
		shouldErr bool
	}{
		{
			name: "invalid user returns error",
			link: types.NewApplicationLink(
				"",
				types.NewData("twitter", "twitteruser"),
				types.ApplicationLinkStateInitialized,
				types.NewOracleRequest(
					0,
					1,
					types.NewOracleRequestCallData(
						"twitter",
						"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
					),
					"client_id",
				),
				nil,
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "invalid data returns error",
			link: types.NewApplicationLink(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				types.NewData("", "twitteruser"),
				types.ApplicationLinkStateInitialized,
				types.NewOracleRequest(
					0,
					1,
					types.NewOracleRequestCallData(
						"twitter",
						"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
					),
					"client_id",
				),
				nil,
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "invalid oracle request returns error",
			link: types.NewApplicationLink(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				types.NewData("twitter", "twitteruser"),
				types.ApplicationLinkStateInitialized,
				types.NewOracleRequest(
					0,
					0,
					types.NewOracleRequestCallData(
						"twitter",
						"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
					),
					"client_id",
				),
				nil,
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "invalid date returns error",
			link: types.NewApplicationLink(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				types.NewData("twitter", "twitteruser"),
				types.ApplicationLinkStateInitialized,
				types.NewOracleRequest(
					0,
					1,
					types.NewOracleRequestCallData(
						"twitter",
						"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
					),
					"client_id",
				),
				nil,
				time.Time{},
				time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "invalid error result returns error",
			link: types.NewApplicationLink(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				types.NewData("twitter", "twitteruser"),
				types.ApplicationLinkStateInitialized,
				types.NewOracleRequest(
					0,
					1,
					types.NewOracleRequestCallData(
						"twitter",
						"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
					),
					"client_id",
				),
				types.NewErrorResult(""),
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "invalid success result returns error",
			link: types.NewApplicationLink(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				types.NewData("twitter", "twitteruser"),
				types.ApplicationLinkStateInitialized,
				types.NewOracleRequest(
					0,
					1,
					types.NewOracleRequestCallData(
						"twitter",
						"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
					),
					"client_id",
				),
				types.NewSuccessResult("value", "signature"),
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "invalid expiration time returns error",
			link: types.NewApplicationLink(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				types.NewData("twitter", "twitteruser"),
				types.ApplicationLinkStateInitialized,
				types.NewOracleRequest(
					0,
					1,
					types.NewOracleRequestCallData(
						"twitter",
						"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
					),
					"client_id",
				),
				nil,
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				time.Time{},
			),
			shouldErr: true,
		},
		{
			name: "expiration time before creation time returns error",
			link: types.NewApplicationLink(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				types.NewData("twitter", "twitteruser"),
				types.ApplicationLinkStateInitialized,
				types.NewOracleRequest(
					0,
					1,
					types.NewOracleRequestCallData(
						"twitter",
						"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
					),
					"client_id",
				),
				nil,
				time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
				time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "valid link returns no error",
			link: types.NewApplicationLink(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				types.NewData("twitter", "twitteruser"),
				types.ApplicationLinkStateInitialized,
				types.NewOracleRequest(
					0,
					1,
					types.NewOracleRequestCallData(
						"twitter",
						"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
					),
					"client_id",
				),
				nil,
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.link.Validate()

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMustUnmarshalApplicationLink(t *testing.T) {
	link := types.NewApplicationLink(
		"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
		types.NewData("twitter", "twitteruser"),
		types.ApplicationLinkStateInitialized,
		types.NewOracleRequest(
			0,
			1,
			types.NewOracleRequestCallData(
				"twitter",
				"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
			),
			"client_id",
		),
		nil,
		time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
		time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
	)

	cdc, _ := app.MakeCodecs()
	marshalled := types.MustMarshalApplicationLink(cdc, link)
	unmarshalled := types.MustUnmarshalApplicationLink(cdc, marshalled)
	require.Equal(t, link, unmarshalled)
}

// --------------------------------------------------------------------------------------------------------------------

func TestData_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		data      types.Data
		shouldErr bool
	}{
		{
			name:      "invalid application returns error",
			data:      types.NewData("   ", "twitteruser"),
			shouldErr: true,
		},
		{
			name:      "invalid username returns error",
			data:      types.NewData("twitter", "  "),
			shouldErr: true,
		},
		{
			name:      "valid data returns no error",
			data:      types.NewData("twitter", "twitteruser"),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.data.Validate()

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func TestOracleRequest_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		request   types.OracleRequest
		shouldErr bool
	}{
		{
			name: "invalid script id returns error",
			request: types.NewOracleRequest(
				0,
				0,
				types.NewOracleRequestCallData(
					"twitter",
					"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
				),
				"client_id",
			),
			shouldErr: true,
		},
		{
			name: "invalid call data returns error",
			request: types.NewOracleRequest(
				0,
				1,
				types.NewOracleRequestCallData(
					"",
					"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
				),
				"client_id",
			),
			shouldErr: true,
		},
		{
			name: "invalid client id returns error",
			request: types.NewOracleRequest(
				0,
				1,
				types.NewOracleRequestCallData(
					"twitter",
					"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
				),
				"",
			),
			shouldErr: true,
		},
		{
			name: "valid request returns no error",
			request: types.NewOracleRequest(
				0,
				1,
				types.NewOracleRequestCallData(
					"twitter",
					"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
				),
				"client_id",
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.request.Validate()

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestOracleRequest_CallData_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		data      types.OracleRequest_CallData
		shouldErr bool
	}{
		{
			name: "invalid application returns error",
			data: types.NewOracleRequestCallData(
				"  ",
				"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
			),
			shouldErr: true,
		},
		{
			name:      "empty call data returns error",
			data:      types.NewOracleRequestCallData("twitter", "   "),
			shouldErr: true,
		},
		{
			name:      "non hex call data returns error",
			data:      types.NewOracleRequestCallData("twitter", "call_data"),
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			data: types.NewOracleRequestCallData(
				"twitter",
				"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.data.Validate()

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestApplicationLink_IsVerificationOngoing(t *testing.T) {
	testCases := []struct {
		name     string
		link     types.ApplicationLink
		expected bool
	}{
		{
			name:     "non on going state returns false",
			link:     types.ApplicationLink{State: types.AppLinkStateVerificationSuccess},
			expected: false,
		},
		{
			name:     "verification initialized returns true",
			link:     types.ApplicationLink{State: types.ApplicationLinkStateInitialized},
			expected: true,
		},
		{
			name:     "verification error returns true",
			link:     types.ApplicationLink{State: types.AppLinkStateVerificationStarted},
			expected: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, tc.link.IsVerificationOngoing())
		})
	}
}

func TestApplicationLink_IsVerificationCompleted(t *testing.T) {
	testCases := []struct {
		name     string
		link     types.ApplicationLink
		expected bool
	}{
		{
			name:     "non completed state returns false",
			link:     types.ApplicationLink{State: types.ApplicationLinkStateInitialized},
			expected: false,
		},
		{
			name:     "verification success returns true",
			link:     types.ApplicationLink{State: types.AppLinkStateVerificationSuccess},
			expected: true,
		},
		{
			name:     "verification error returns true",
			link:     types.ApplicationLink{State: types.AppLinkStateVerificationError},
			expected: true,
		},
		{
			name:     "verification timeout returns true",
			link:     types.ApplicationLink{State: types.AppLinkStateVerificationTimedOut},
			expected: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, tc.link.IsVerificationCompleted())
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func TestResult_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		result    *types.Result
		shouldErr bool
	}{
		{
			name:      "invalid error result returns error",
			result:    types.NewErrorResult(" "),
			shouldErr: true,
		},
		{
			name:      "invalid success result returns error",
			result:    types.NewSuccessResult(" ", " "),
			shouldErr: true,
		},
		{
			name:      "valid error result returns no error",
			result:    types.NewErrorResult("error"),
			shouldErr: false,
		},
		{
			name: "valid success result returns no error",
			result: types.NewSuccessResult(
				"value",
				"a00a7d5bd45e42615645fcaeb4d800af22704e54937ab235e5e50bebd38e88b765fdb696c22712c0cab1176756b6346cbc11481c544d1f7828cb233620c06173",
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.result.Validate()

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestResult_Failed__Validate(t *testing.T) {
	testCases := []struct {
		name      string
		result    *types.Result
		shouldErr bool
	}{
		{
			name:      "invalid error returns error",
			result:    types.NewErrorResult(" "),
			shouldErr: true,
		},
		{
			name:      "valid error returns no error",
			result:    types.NewErrorResult("error"),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.result.Validate()

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestResult_Success__Validate(t *testing.T) {
	testCases := []struct {
		name      string
		result    *types.Result
		shouldErr bool
	}{
		{
			name: "invalid value returns error",
			result: types.NewSuccessResult(
				" ",
				"a00a7d5bd45e42615645fcaeb4d800af22704e54937ab235e5e50bebd38e88b765fdb696c22712c0cab1176756b6346cbc11481c544d1f7828cb233620c06173",
			),
			shouldErr: true,
		},
		{
			name:      "empty signature returns error",
			result:    types.NewSuccessResult("value", "  "),
			shouldErr: true,
		},
		{
			name:      "invalid signature returns error",
			result:    types.NewSuccessResult("value", "signature"),
			shouldErr: true,
		},
		{
			name: "valid result returns no error",
			result: types.NewSuccessResult(
				"value",
				"a00a7d5bd45e42615645fcaeb4d800af22704e54937ab235e5e50bebd38e88b765fdb696c22712c0cab1176756b6346cbc11481c544d1f7828cb233620c06173",
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.result.Validate()

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
