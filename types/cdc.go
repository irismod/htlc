package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(&MsgCreateHTLC{}, "irismod/htlc/MsgCreateHTLC", nil)
	cdc.RegisterConcrete(&MsgClaimHTLC{}, "irismod/htlc/MsgClaimHTLC", nil)
	cdc.RegisterConcrete(&MsgRefundHTLC{}, "irismod/htlc/MsgRefundHTLC", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateHTLC{},
		&MsgClaimHTLC{},
		&MsgRefundHTLC{},
	)
}

// ModuleCdc defines the module codec
var (
	amino = codec.New()

	ModuleCdc = codec.NewHybridCodec(amino, types.NewInterfaceRegistry())
)

func init() {
	RegisterCodec(amino)
	cryptocodec.RegisterCrypto(amino)
}
