package cli

import (
	"context"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/irismod/htlc/types"
)

// GetQueryCmd returns the cli query commands for the module.
func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	htlcQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the HTLC module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	htlcQueryCmd.AddCommand(flags.GetCommands(
		GetCmdQueryHTLC(cdc),
	)...)

	return htlcQueryCmd
}

// GetCmdQueryHTLC implements the query HTLC command.
func GetCmdQueryHTLC(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "htlc [hash-lock]",
		Short: "Query an HTLC",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query details of an HTLC with the specified hash lock.

Example:
$ %s query htlc htlc <hash-lock>
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := client.NewContext().WithCodec(cdc).WithJSONMarshaler(cdc)

			hashLock, err := hex.DecodeString(args[0])
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(cliCtx)
			param := types.QueryHTLCRequest{HashLock: hashLock}
			response, err := queryClient.HTLC(context.Background(), &param)
			if err != nil {
				return err
			}
			return cliCtx.PrintOutput(response.Htlc)
		},
	}
}
