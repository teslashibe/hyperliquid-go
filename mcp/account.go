package hyperliquidmcp

import (
	"context"

	"github.com/teslashibe/hyperliquid-go"
	"github.com/teslashibe/mcptool"
)

// GetAccountStateInput is the typed input for hyperliquid_get_account_state.
type GetAccountStateInput struct {
	Wallet string `json:"wallet" jsonschema:"description=42-character 0x-prefixed Ethereum address (e.g. 0x5FB80F80C397F01d0287b009B74530eBeAA170B5).,required"`
}

func getAccountState(ctx context.Context, c *hyperliquid.Client, in GetAccountStateInput) (any, error) {
	state, err := c.FetchAccountState(ctx, in.Wallet)
	if err != nil {
		return nil, wrapErr(err, "get_account_state")
	}
	return state, nil
}

var accountTools = []mcptool.Tool{
	mcptool.Define[*hyperliquid.Client, GetAccountStateInput](
		"hyperliquid_get_account_state",
		"Fetch a wallet's full Hyperliquid perpetuals account state: account value, withdrawable balance, margin utilization, and all open positions with entry price, size, leverage, unrealized PnL, and liquidation distance.",
		"FetchAccountState",
		getAccountState,
	),
}
