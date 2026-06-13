package hyperliquidmcp

import (
	"context"

	"github.com/teslashibe/hyperliquid-go"
	"github.com/teslashibe/mcptool"
)

// GetAccountStateInput is the typed input for hyperliquid_get_account_state.
type GetAccountStateInput struct {
	Wallet string `json:"wallet,omitempty" jsonschema:"description=42-character 0x-prefixed address (e.g. 0x5FB80F80C397F01d0287b009B74530eBeAA170B5). Optional when the client has a default wallet configured."`
}

func getAccountState(ctx context.Context, c *hyperliquid.Client, in GetAccountStateInput) (any, error) {
	wallet, err := resolveWallet(c, in.Wallet)
	if err != nil {
		return nil, err
	}
	state, err := c.FetchAccountState(ctx, wallet)
	if err != nil {
		return nil, wrapErr(err, "get_account_state")
	}
	return state, nil
}

var accountTools = []mcptool.Tool{
	mcptool.Define[*hyperliquid.Client, GetAccountStateInput](
		"hyperliquid_get_account_state",
		"Fetch a wallet's Hyperliquid account: value, margin use, and open positions (entry, size, leverage, PnL, liquidation).",
		"FetchAccountState",
		getAccountState,
	),
}
