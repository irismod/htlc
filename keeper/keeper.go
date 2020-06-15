package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"

	"github.com/irismod/htlc/types"
)

// Keeper defines the HTLC keeper
type Keeper struct {
	storeKey      sdk.StoreKey
	cdc           codec.Marshaler
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
}

// NewKeeper creates a new HTLC Keeper instance
func NewKeeper(
	cdc codec.Marshaler,
	key sdk.StoreKey,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) Keeper {
	// ensure the HTLC module account is set
	if addr := accountKeeper.GetModuleAddress(types.HTLCAccName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.HTLCAccName))
	}

	return Keeper{
		storeKey:      key,
		cdc:           cdc,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
	}
}

// Logger returns a module-specific logger
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("irismod/%s", types.ModuleName))
}

// GetHTLCAccount returns the HTLC module account
func (k Keeper) GetHTLCAccount(ctx sdk.Context) exported.ModuleAccountI {
	return k.accountKeeper.GetModuleAccount(ctx, types.HTLCAccName)
}
