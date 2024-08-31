package cmd

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/terramirum/mirumd/app"
)

func initSDKConfig() {
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount(app.Bech32PrefixAccAddr, app.Bech32PrefixAccPub)
	cfg.SetBech32PrefixForValidator(app.Bech32PrefixValAddr, app.Bech32PrefixValPub)
	cfg.SetBech32PrefixForConsensusNode(app.Bech32PrefixConsAddr, app.Bech32PrefixConsPub)
	cfg.SetAddressVerifier(wasmtypes.VerifyAddressLen())
	cfg.SetCoinType(app.CoinType)
	cfg.SetFullFundraiserPath(app.FullFundraiserPath)
	cfg.Seal()
}
