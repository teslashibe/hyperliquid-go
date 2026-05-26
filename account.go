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
//
// Hyperliquid Unified Account: collateral can sit in either the perps
// subaccount (clearinghouseState) or as spot USDC. We query both and
// combine for the true Unified Account view that matches the UI.
func (c *Client) FetchAccountState(ctx context.Context, wallet string) (*AccountState, error) {
	if err := validateAddress(wallet); err != nil {
		return nil, err
	}

	perpsState, err := c.info.UserState(ctx, wallet)
	if err != nil {
		return nil, &Error{Code: ErrUpstreamFailed, Message: "user_state: " + err.Error()}
	}

	spotState, err := c.info.SpotUserState(ctx, wallet)
	if err != nil {
		return nil, &Error{Code: ErrUpstreamFailed, Message: "spot_user_state: " + err.Error()}
	}

	perpsAccountValue := parseFloat(perpsState.MarginSummary.AccountValue)
	perpsWithdrawable := parseFloat(perpsState.Withdrawable)
	totalMarginUsed := parseFloat(perpsState.MarginSummary.TotalMarginUsed)
	totalNotional := parseFloat(perpsState.MarginSummary.TotalNtlPos)

	spotUSDC := 0.0
	for _, b := range spotState.Balances {
		if b.Coin == "USDC" {
			spotUSDC += parseFloat(b.Total)
		}
	}

	totalAccountValue := perpsAccountValue + spotUSDC
	totalWithdrawable := perpsWithdrawable + spotUSDC

	out := &AccountState{
		Wallet:           wallet,
		AccountValueUSD:  totalAccountValue,
		WithdrawableUSD:  totalWithdrawable,
		TotalNotionalUSD: totalNotional,
		TotalMarginUsed:  totalMarginUsed,
		FetchedAt:        time.Now(),
	}

	if totalAccountValue > 0 {
		out.MarginUtilization = (totalMarginUsed / totalAccountValue) * 100
	}

	out.Positions = make([]Position, 0, len(perpsState.AssetPositions))
	for _, ap := range perpsState.AssetPositions {
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
