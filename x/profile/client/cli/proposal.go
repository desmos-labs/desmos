package cli

import (
	"bufio"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	"github.com/desmos-labs/desmos/x/profile/internal/types/models"
	"github.com/spf13/cobra"
)

// parseProposalCommonFields parse the common fields of proposals
func parseProposalCommonFields(cmd *cobra.Command) (string, string, sdk.Coins, error) {
	title, err := cmd.Flags().GetString(cli.FlagTitle)
	if err != nil {
		return "", "", nil, err
	}

	description, err := cmd.Flags().GetString(cli.FlagDescription)
	if err != nil {
		return "", "", nil, err
	}

	depositStr, err := cmd.Flags().GetString(cli.FlagDeposit)
	if err != nil {
		return "", "", nil, err
	}

	deposit, err := sdk.ParseCoins(depositStr)
	if err != nil {
		return "", "", nil, err
	}

	return title, description, deposit, nil
}

// GetCmdSubmitNameSurnameParamsEditProposal is the CLI command for submitting an edit proposal to name/surname params transaction.
func GetCmdSubmitNameSurnameParamsEditProposal(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ns-params-edit",
		Short: "Submit an edit proposal for name and surname params lengths",
		Long: fmt.Sprintf(`
		Submit an edit proposal for name and surname params lengths.
You should specify at least one of the two parameters otherwise the proposal will not be considered valid.

%s tx gov submit-proposal ns-params-edit \
--title editProp  \
--description="My awesome proposal" \
--deposit 10desmos \ 
--min-len 3 \
--max-len 200 \
--from leo
`, version.ClientName),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)
			from := cliCtx.GetFromAddress()

			title, description, deposit, err := parseProposalCommonFields(cmd)
			if err != nil {
				return err
			}

			minNSLen, err := cmd.Flags().GetInt64(flagMinParamsLen)
			if err != nil {
				return err
			}
			maxNSLen, err := cmd.Flags().GetInt64(flagMaxParamsLen)
			if err != nil {
				return err
			}

			var nsLenParams models.NameSurnameLengths

			if minNSLen == -1 && maxNSLen == -1 {
				return fmt.Errorf("invalid proposal. At least one parameter should be specified")
			} else {
				maxMonikerLenParam := sdk.NewInt(maxNSLen)
				minMonikerLenParam := sdk.NewInt(minNSLen)
				if minNSLen == -1 && maxNSLen != -1 {
					nsLenParams = models.NewNameSurnameLenParams(nil, &maxMonikerLenParam)
				} else if minNSLen != -1 && maxNSLen == -1 {
					nsLenParams = models.NewNameSurnameLenParams(&minMonikerLenParam, nil)
				} else {
					nsLenParams = models.NewNameSurnameLenParams(&minMonikerLenParam, &maxMonikerLenParam)
				}
			}

			content := models.NewNameSurnameParamsEditProposal(title, description, nsLenParams)

			msg := gov.NewMsgSubmitProposal(content, deposit, from)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(cli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(cli.FlagDeposit, "", "deposit of proposal")
	cmd.Flags().Int64(flagMinParamsLen, -1, "The min value for a param length")
	cmd.Flags().Int64(flagMaxParamsLen, -1, "The max value for a param length")

	return cmd
}

// GetCmdSubmitMonikerParamsEditProposal is the CLI command for submitting an edit proposal to moniker params transaction.
func GetCmdSubmitMonikerParamsEditProposal(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "moniker-params-edit",
		Short: "Submit an edit proposal for moniker params lengths",
		Long: fmt.Sprintf(`
		Submit an edit proposal for moniker params lengths.
You should specify at least one of the two parameters otherwise the proposal will not be considered valid.

%s tx gov submit-proposal moniker-params-edit \
--title editProp  \
--description="My awesome proposal" \
--deposit 10desmos \ 
--min-len 5 \
--max-len 40 \
--from leo
`, version.ClientName),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)
			from := cliCtx.GetFromAddress()

			title, description, deposit, err := parseProposalCommonFields(cmd)
			if err != nil {
				return err
			}

			minMonikerLen, err := cmd.Flags().GetInt64(flagMinParamsLen)
			if err != nil {
				return err
			}
			maxMonikerLen, err := cmd.Flags().GetInt64(flagMaxParamsLen)
			if err != nil {
				return err
			}

			var monikerParams models.MonikerLengths

			if minMonikerLen == -1 && maxMonikerLen == -1 {
				return fmt.Errorf("invalid proposal. At least one parameter should be specified")
			} else {
				minMonikerLenParam := sdk.NewInt(minMonikerLen)
				maxMonikerLenParam := sdk.NewInt(maxMonikerLen)
				if minMonikerLen == -1 && maxMonikerLen != -1 {
					monikerParams = models.NewMonikerLenParams(nil, &maxMonikerLenParam)
				} else if minMonikerLen != -1 && maxMonikerLen == -1 {
					monikerParams = models.NewMonikerLenParams(&minMonikerLenParam, nil)
				} else {
					monikerParams = models.NewMonikerLenParams(&minMonikerLenParam, &maxMonikerLenParam)
				}
			}

			content := models.NewMonikerParamsEditProposal(title, description, monikerParams)

			msg := gov.NewMsgSubmitProposal(content, deposit, from)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})

		},
	}
	cmd.Flags().String(cli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(cli.FlagDeposit, "", "deposit of proposal")
	cmd.Flags().Int64(flagMinParamsLen, -1, "The min value for a param length")
	cmd.Flags().Int64(flagMaxParamsLen, -1, "The max value for a param length")

	return cmd
}

// GetCmdSubmitBioParamsEditProposal is the CLI command for submitting an edit proposal to biography params transaction.
func GetCmdSubmitBioParamsEditProposal(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bio-params-edit",
		Short: "Submit an edit proposal for biography param lengths",
		Long: fmt.Sprintf(`
		Submit an edit proposal for biography param lengths.
%s tx gov submit-proposal bio-params-edit \
--title editProp  \
--description="My awesome proposal" \
--deposit 10desmos \
--max-len 200 \
--from leo
`, version.ClientName),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)
			from := cliCtx.GetFromAddress()

			title, description, deposit, err := parseProposalCommonFields(cmd)
			if err != nil {
				return err
			}

			maxBioLen, err := cmd.Flags().GetInt64(flagMaxParamsLen)
			if err != nil {
				return err
			}

			var bioParams models.BiographyLengths

			if maxBioLen == -1 {
				return fmt.Errorf("invalid proposal. No parameters specified")
			} else {
				bioParams = models.NewBioLenParams(sdk.NewInt(maxBioLen))
			}

			content := models.NewBioParamsEditProposal(title, description, bioParams)

			msg := gov.NewMsgSubmitProposal(content, deposit, from)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})

		},
	}
	cmd.Flags().String(cli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(cli.FlagDeposit, "", "deposit of proposal")
	cmd.Flags().Int64(flagMaxParamsLen, -1, "The max value for a param length")

	return cmd
}
