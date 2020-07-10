package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/irismod/htlc/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) HTLC(c context.Context, request *types.QueryHTLCRequest) (*types.QueryHTLCResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	htlc, found := k.GetHTLC(ctx, request.HashLock)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrUnknownHTLC, request.HashLock.String())
	}

	return &types.QueryHTLCResponse{Htlc: &htlc}, nil
}
