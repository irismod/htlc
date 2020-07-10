package helpers

import (
	"math/rand"
	"time"

	"github.com/tendermint/tendermint/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// SimAppChainID hardcoded chainID for simulation
const (
	DefaultGenTxGas = 1000000
	SimAppChainID   = "simulation-app"
)

// GenTx generates a signed mock transaction.
func GenTx(msgs []sdk.Msg, feeAmt sdk.Coins, gas uint64, chainID string, accnums []uint64, seq []uint64, priv ...crypto.PrivKey) authtypes.StdTx {
	fee := authtypes.StdFee{
		Amount: feeAmt,
		Gas:    gas,
	}

	sigs := make([]authtypes.StdSignature, len(priv))

	// create a random length memo
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	memo := simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 0, 100))

	for i, p := range priv {
		// use a empty chainID for ease of testing
		sig, err := p.Sign(authtypes.StdSignBytes(chainID, accnums[i], seq[i], fee, msgs, memo))
		if err != nil {
			panic(err)
		}

		sigs[i] = authtypes.StdSignature{
			PubKey:    p.PubKey().Bytes(),
			Signature: sig,
		}
	}

	return authtypes.NewStdTx(msgs, fee, sigs, memo)
}
