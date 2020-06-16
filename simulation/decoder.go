package simulation

import (
	"bytes"
	"fmt"

	tmkv "github.com/tendermint/tendermint/libs/kv"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/irismod/htlc/types"
)

// NewDecodeStore unmarshals the KVPair's Value to the corresponding HTLC type
func NewDecodeStore(cdc codec.Marshaler) func(kvA, kvB tmkv.Pair) string {
	return func(kvA, kvB tmkv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], types.HTLCKey):
			var htlc1, htlc2 types.HTLC
			cdc.MustUnmarshalBinaryBare(kvA.Value, &htlc1)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &htlc2)
			return fmt.Sprintf("%v\n%v", htlc1, htlc2)

		case bytes.Equal(kvA.Key[:1], types.HTLCExpiredQueueKey):
			return fmt.Sprintf("%v\n%v", kvA.Value, kvB.Value)
		default:
			panic(fmt.Sprintf("invalid HTLC key prefix %X", kvA.Key[:1]))
		}
	}
}
