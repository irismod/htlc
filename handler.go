package htlc

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler creates an sdk.Handler for all the HTLC type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *MsgCreateHTLC:
			return handleMsgCreateHTLC(ctx, k, msg)

		case *MsgClaimHTLC:
			return handleMsgClaimHTLC(ctx, k, msg)

		case *MsgRefundHTLC:
			return handleMsgRefundHTLC(ctx, k, msg)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", ModuleName, msg)
		}
	}
}

// handleMsgCreateHTLC handles MsgCreateHTLC
func handleMsgCreateHTLC(ctx sdk.Context, k Keeper, msg *MsgCreateHTLC) (*sdk.Result, error) {
	err := k.CreateHTLC(
		ctx,
		msg.Sender,
		msg.To,
		msg.ReceiverOnOtherChain,
		msg.Amount,
		msg.HashLock,
		msg.Timestamp,
		msg.TimeLock,
	)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			EventTypeCreateHTLC,
			sdk.NewAttribute(AttributeKeySender, msg.Sender.String()),
			sdk.NewAttribute(AttributeKeyReceiver, msg.To.String()),
			sdk.NewAttribute(AttributeKeyReceiverOnOtherChain, msg.ReceiverOnOtherChain),
			sdk.NewAttribute(AttributeKeyAmount, msg.Amount.String()),
			sdk.NewAttribute(AttributeKeyHashLock, msg.HashLock.String()),
			sdk.NewAttribute(AttributeKeyTimeLock, fmt.Sprintf("%d", msg.TimeLock)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

// handleMsgClaimHTLC handles MsgClaimHTLC
func handleMsgClaimHTLC(ctx sdk.Context, k Keeper, msg *MsgClaimHTLC) (*sdk.Result, error) {
	err := k.ClaimHTLC(ctx, msg.HashLock, msg.Secret)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			EventTypeClaimHTLC,
			sdk.NewAttribute(AttributeKeySender, msg.Sender.String()),
			sdk.NewAttribute(AttributeKeyHashLock, msg.HashLock.String()),
			sdk.NewAttribute(AttributeKeySecret, msg.Secret.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

// handleMsgRefundHTLC handles MsgRefundHTLC
func handleMsgRefundHTLC(ctx sdk.Context, k Keeper, msg *MsgRefundHTLC) (*sdk.Result, error) {
	err := k.RefundHTLC(ctx, msg.HashLock)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			EventTypeRefundHTLC,
			sdk.NewAttribute(AttributeKeySender, msg.Sender.String()),
			sdk.NewAttribute(AttributeKeyHashLock, msg.HashLock.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
