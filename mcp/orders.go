package hyperliquidmcp

import (
	"context"

	"github.com/teslashibe/hyperliquid-go"
	"github.com/teslashibe/mcptool"
)

// GetOpenOrdersInput is the typed input for hyperliquid_get_open_orders.
type GetOpenOrdersInput struct {
	Wallet string `json:"wallet" jsonschema:"description=42-character 0x-prefixed Ethereum address.,required"`
}

// GetRecentFillsInput is the typed input for hyperliquid_get_recent_fills.
type GetRecentFillsInput struct {
	Wallet string `json:"wallet" jsonschema:"description=42-character 0x-prefixed Ethereum address.,required"`
	Limit  int    `json:"limit,omitempty" jsonschema:"description=Maximum number of recent fills to return (most recent first). Default 50.,minimum=1,maximum=500,default=50"`
}

func getOpenOrders(ctx context.Context, c *hyperliquid.Client, in GetOpenOrdersInput) (any, error) {
	orders, err := c.FetchOpenOrders(ctx, in.Wallet)
	if err != nil {
		return nil, wrapErr(err, "get_open_orders")
	}
	return orders, nil
}

func getRecentFills(ctx context.Context, c *hyperliquid.Client, in GetRecentFillsInput) (any, error) {
	limit := in.Limit
	if limit <= 0 {
		limit = 50
	}
	fills, err := c.FetchRecentFills(ctx, in.Wallet, limit)
	if err != nil {
		return nil, wrapErr(err, "get_recent_fills")
	}
	return fills, nil
}

var orderTools = []mcptool.Tool{
	mcptool.Define[*hyperliquid.Client, GetOpenOrdersInput](
		"hyperliquid_get_open_orders",
		"List all pending limit orders for a wallet on Hyperliquid (asset, side, size, limit price, placed timestamp).",
		"FetchOpenOrders",
		getOpenOrders,
	),
	mcptool.Define[*hyperliquid.Client, GetRecentFillsInput](
		"hyperliquid_get_recent_fills",
		"Retrieve recent executed trades for a wallet on Hyperliquid (most recent first), including realized PnL per fill.",
		"FetchRecentFills",
		getRecentFills,
	),
}
