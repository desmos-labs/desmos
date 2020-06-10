package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/desmos-labs/desmos/x/profile/internal/types/models"
	"github.com/gorilla/mux"
	"net/http"

	govrest "github.com/cosmos/cosmos-sdk/x/gov/client/rest"
)

func registerProposalRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/profiles/proposals/name-surname-params-edit", nameSurnameParamsEditHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/profiles/proposals/moniker-params-edit", monikerParamsEditHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/profiles/proposals/biography-params-edit", bioParamsEditHandler(cliCtx)).Methods("POST")
}

type NameSurnameParamsEditRequest struct {
	BaseReq           rest.BaseReq                `json:"base_req" yaml:"base_req"`
	Title             string                      `json:"title" yaml:"title"`
	Description       string                      `json:"description" yaml:"description"`
	Deposit           sdk.Coins                   `json:"deposit" yaml:"deposit"`
	NameSurnameParams models.NameSurnameLenParams `json:"name_surname_params" yaml:"name_surname_params"`
}

type MonikerParamsEditRequest struct {
	BaseReq       rest.BaseReq            `json:"base_req" yaml:"base_req"`
	Title         string                  `json:"title" yaml:"title"`
	Description   string                  `json:"description" yaml:"description"`
	Deposit       sdk.Coins               `json:"deposit" yaml:"deposit"`
	MonikerParams models.MonikerLenParams `json:"moniker_params" yaml:"moniker_params"`
}

type BioParamsEditRequest struct {
	BaseReq      rest.BaseReq        `json:"base_req" yaml:"base_req"`
	Title        string              `json:"title" yaml:"title"`
	Description  string              `json:"description" yaml:"description"`
	Deposit      sdk.Coins           `json:"deposit" yaml:"deposit"`
	BioLenParams models.BioLenParams `json:"bio_len_params" yaml:"bio_len_params"`
}

func NameSurnameParamsEditProposalRESTHandler(cliCtx context.CLIContext) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "profile/name-surname-params-edit",
		Handler:  nameSurnameParamsEditHandler(cliCtx),
	}
}

func nameSurnameParamsEditHandler(cliCtx context.CLIContext) http.HandlerFunc {
	var req NameSurnameParamsEditRequest
	return func(w http.ResponseWriter, r *http.Request) {
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		editProp := models.NewNameSurnameParamsEditProposal(req.Title, req.Description, req.NameSurnameParams)

		msg := gov.NewMsgSubmitProposal(editProp, req.Deposit, fromAddr)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func MonikerParamsEditProposalRESTHandler(cliCtx context.CLIContext) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "profile/moniker-params-edit",
		Handler:  monikerParamsEditHandler(cliCtx),
	}
}

func monikerParamsEditHandler(cliCtx context.CLIContext) http.HandlerFunc {
	var req MonikerParamsEditRequest
	return func(w http.ResponseWriter, r *http.Request) {
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		editProp := models.NewMonikerParamsEditProposal(req.Title, req.Description, req.MonikerParams)

		msg := gov.NewMsgSubmitProposal(editProp, req.Deposit, fromAddr)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func BioParamsEditProposalRESTHandler(cliCtx context.CLIContext) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "profile/biography-params-edit",
		Handler:  monikerParamsEditHandler(cliCtx),
	}
}

func bioParamsEditHandler(cliCtx context.CLIContext) http.HandlerFunc {
	var req BioParamsEditRequest
	return func(w http.ResponseWriter, r *http.Request) {
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		editProp := models.NewBioParamsEditProposal(req.Title, req.Description, req.BioLenParams)

		msg := gov.NewMsgSubmitProposal(editProp, req.Deposit, fromAddr)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}
