package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/deweb-services/deweb/x/dns_module/types"
)

func registerQueryRoutes(cliCtx client.Context, r *mux.Router, queryRoute string) {
	// Get the params of module
	r.HandleFunc(fmt.Sprintf("/%s/params", types.ModuleName), queryParams(cliCtx, queryRoute)).Methods("GET")
	// Query a single domain NFT
	r.HandleFunc(fmt.Sprintf("/deweb/domains/v1beta1/domain/{%s}", RestParamDomainName), queryDomain(cliCtx, queryRoute)).Methods("GET")
}

func queryParams(cliCtx client.Context, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		res, height, err := cliCtx.Query(
			fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryParams),
		)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryDomain(cliCtx client.Context, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		domainName := vars[RestParamDomainName]
		if err := types.ValidateTokenID(domainName); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		params := types.NewQueryDomainsParams(domainName)
		bz, err := cliCtx.LegacyAmino.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		res, height, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryDomain), bz,
		)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}
