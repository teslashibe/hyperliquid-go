package hyperliquid

import (
	"context"
	"math"
	"strconv"
	"strings"
	"time"

	hl "github.com/sonirico/go-hyperliquid"
)

// FetchAccountState retrieves the user's perpetuals account state and
// converts the underlying string-typed fields into typed floats with
// derived metrics computed.
func (c *Client) FetchAccountState(ctx context.Context, wallet string) (*AccountState, error) {
	if err := validateAddress(wallet); err != nil {
		return nil, err
	}

	state, err := c.info.UserState(ctx, wallet)
	if err != nil {
		return nil, &Error{Code: ErrUpstreamFailed, Message: "user_state: " + err.Error()}
	}

	out := &AccountState{
		Wallet:           wallet,
		AccountValueUSD:  parseFloat(state.MarginSummary.AccountValue),
		WithdrawableUSD:  parseFloat(state.Withdrawable),
		TotalNotionalUSD: parseFloat(state.MarginSummary.TotalNtlPos),
		TotalMarginUsed:  parseFloat(state.MarginSummary.TotalMarginUsed),
		FetchedAt:        time.Now(),
	}

	if out.AccountValueUSD > 0 {
		out.MarginUtilization = (out.TotalMarginUsed / out.AccountValueUSD) * 100
	}

	out.Positions = make([]Position, 0, len(state.AssetPositions))
	for _, ap := range state.AssetPositions {
		out.Positions = append(out.Positions, normalizePosition(ap))
	}

	return out, nil
}

func normalizePosition(ap hl.AssetPosition) Position {
	p := ap.Position
	size := parseFloat(p.Szi)
	direction := "LONG"
	if size < 0 {
		direction = "SHORT"
	}

	pos := Position{
		Asset:            p.Coin,
		Direction:        direction,
		Size:             math.Abs(size),
		EntryPrice:       parsePtrFloat(p.EntryPx),
		PositionValueUSD: parseFloat(p.PositionValue),
		UnrealizedPnLUSD: parseFloat(p.UnrealizedPnl),
		LeverageType:     p.Leverage.Type,
		LeverageValue:    p.Leverage.Value,
		LiquidationPrice: parsePtrFloat(p.LiquidationPx),
		MarginUsedUSD:    parseFloat(p.MarginUsed),
	}

	if pos.PositionValueUSD > 0 {
		pos.UnrealizedPnLPct = (pos.UnrealizedPnLUSD / pos.PositionValueUSD) * 100
	}
	if pos.LiquidationPrice > 0 && pos.EntryPrice > 0 {
		pos.LiquidationDistPct = math.Abs(pos.EntryPrice-pos.LiquidationPrice) / pos.EntryPrice * 100
	}
	if p.CumFunding != nil {
		pos.FundingPaidTotal = parseFloat(p.CumFunding.AllTime)
	}

	return pos
}

func parseFloat(s string) float64 {
	v, _ := strconv.ParseFloat(strings.TrimSpace(s), 64)
	return v
}

func parsePtrFloat(s *string) float64 {
	if s == nil {
		return 0
	}
	return parseFloat(*s)
}
