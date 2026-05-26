// Package hyperliquid provides a clean, normalized read-only client for
// Hyperliquid account state, positions, and order history.
//
// It wraps github.com/sonirico/go-hyperliquid for the underlying transport
// and exposes a Maxwell-friendly type system with:
//   - typed floats (sonirico uses string fields for precision)
//   - computed fields (PnL %, distance to liquidation, margin utilization)
//   - structured errors with stable codes
//
// V0 is read-only — no signing or order placement. Future versions will
// add execution capabilities through a separate Trader type.
//
// Usage:
//
//	c := hyperliquid.New(hyperliquid.DefaultConfig())
//	defer c.Close()
//
//	state, err := c.FetchAccountState(ctx, "0x...")
//	for _, p := range state.Positions {
//	    fmt.Printf("%s %s @ %.2f, PnL %.2f%%\n", p.Direction, p.Asset, p.EntryPrice, p.UnrealizedPnLPct)
//	}
package hyperliquid
