package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
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

	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
