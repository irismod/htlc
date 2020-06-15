package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// HTLC module sentinel errors
var (
	ErrInvalidHashLock = sdkerrors.Register(ModuleName, 3, "invalid hash lock")
	ErrInvalidTimeLock = sdkerrors.Register(ModuleName, 5, "invalid time lock")
	ErrInvalidSecret   = sdkerrors.Register(ModuleName, 6, "invalid secret")
	ErrHTLCExists      = sdkerrors.Register(ModuleName, 4, "htlc already exists")
	ErrUnknownHTLC     = sdkerrors.Register(ModuleName, 4, "unknown htlc")
	ErrHTLCNotOpen     = sdkerrors.Register(ModuleName, 7, "htlc not open")
	ErrHTLCNotExpired  = sdkerrors.Register(ModuleName, 8, "htlc not expired")
)
