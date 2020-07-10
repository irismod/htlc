package htlc

import (
	"encoding/hex"
	"fmt"

	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/irismod/htlc/keeper"
	"github.com/irismod/htlc/types"
)

// InitGenesis stores the genesis state
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {
	if err := types.ValidateGenesis(data); err != nil {
		panic(err.Error())
	}

	for hashLockStr, htlc := range data.PendingHTLCs {
		hashLock, _ := hex.DecodeString(hashLockStr)

		k.SetHTLC(ctx, htlc, hashLock)
		k.AddHTLCToExpiredQueue(ctx, htlc.ExpirationHeight, hashLock)
	}
}

// ExportGenesis outputs the genesis state
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	pendingHTLCs := make(map[string]types.HTLC)

	k.IterateHTLCs(ctx, func(hlock tmbytes.HexBytes, h types.HTLC) (stop bool) {
		if h.State == types.Open {
			h.ExpirationHeight = h.ExpirationHeight - uint64(ctx.BlockHeight()) + 1
			pendingHTLCs[hlock.String()] = h
		} else if h.State == types.Expired {
			err := k.RefundHTLC(ctx, hlock)
			if err != nil {
				panic(fmt.Errorf("failed to export the HTLC genesis state: %s", hlock.String()))
			}
		}

		return false
	})

	return types.GenesisState{
		PendingHTLCs: pendingHTLCs,
	}
}
