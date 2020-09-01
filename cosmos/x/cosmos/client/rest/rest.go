package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers cosmos-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding
	r.HandleFunc("/cosmos/asset", listAssetHandler(cliCtx, "cosmos")).Methods("GET")
	r.HandleFunc("/cosmos/asset", createAssetHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/cosmos/comment", listCommentHandler(cliCtx, "cosmos")).Methods("GET")
	r.HandleFunc("/cosmos/comment", createCommentHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/cosmos/post", listPostHandler(cliCtx, "cosmos")).Methods("GET")
	r.HandleFunc("/cosmos/post", createPostHandler(cliCtx)).Methods("POST")
}
