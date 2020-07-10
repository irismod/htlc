package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateHTLC{}, "irismod/htlc/MsgCreateHTLC", nil)
	cdc.RegisterConcrete(MsgClaimHTLC{}, "irismod/htlc/MsgClaimHTLC", nil)
	cdc.RegisterConcrete(MsgRefundHTLC{}, "irismod/htlc/MsgRefundHTLC", nil)
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

	// ModuleCdc references the global irismod/token module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/staking and
	// defined at the application level.
	ModuleCdc = codec.NewHybridCodec(amino, types.NewInterfaceRegistry())
)

func init() {
	RegisterCodec(amino)
	cryptocodec.RegisterCrypto(amino)
}
