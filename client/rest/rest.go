package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
)

// Rest variable names
// nolint
const (
	RestHashLock = "hash-lock"
)

// RegisterRoutes defines routes that get registered by the main application
func RegisterRoutes(cliCtx client.Context, r *mux.Router) {
	registerQueryRoutes(cliCtx, r)
	registerTxRoutes(cliCtx, r)
}
