package cli

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/terramirum/mirumd/x/rental/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		NewRentNft(),
	)

	return cmd
}

func NewRentNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "class [class_id] [nft_id] [Renter] [FromDate] [ToDate]",
		Short: "rent funds from one account to another.",
		Long: `rent funds from one account to another.
Note, the '--from' flag is ignored as it is implied from [from_key_or_address].
When using '--dry-run' a key name cannot be used, only a bech32 address.
`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			renter, err := sdk.AccAddressFromBech32(args[2])
			if err != nil {
				return err
			}

			startDate, err := strconv.ParseInt(args[3], 10, 64)
			if err != nil {
				return err
			}

			endDate, err := strconv.ParseInt(args[4], 10, 64)
			if err != nil {
				return err
			}

			msg := &types.MsgMintRentRequest{
				ContractOwner: clientCtx.GetFromAddress().String(),
				ClassId:       args[0],
				NftId:         args[1],
				Renter:        renter.String(),
				StartDate:     startDate,
				EndDate:       endDate,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
