// Package hyperliquidmcp exposes github.com/teslashibe/hyperliquid-go as a
// set of mcptool.Tool values backing a single mcptool.Provider.
//
// Tools:
//   - hyperliquid_get_account_state   full account snapshot with positions
//   - hyperliquid_get_open_orders     pending limit orders
//   - hyperliquid_get_recent_fills    recent executed trades
//
// All tools take a *hyperliquid.Client (read-only) and are safe to share
// across concurrent MCP requests.
package hyperliquidmcp

import "github.com/teslashibe/mcptool"

// Provider implements mcptool.Provider for hyperliquid-go.
// Zero value is ready to use.
type Provider struct{}

// Platform returns "hyperliquid".
func (Provider) Platform() string { return "hyperliquid" }

// Tools returns every hyperliquid_* tool exposed by this provider.
func (Provider) Tools() []mcptool.Tool {
	out := make([]mcptool.Tool, 0, 3)
	out = append(out, accountTools...)
	out = append(out, orderTools...)
	return out
}
