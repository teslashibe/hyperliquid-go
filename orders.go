package hyperliquid

import (
	"context"
	"time"

	hl "github.com/sonirico/go-hyperliquid"
)

// FetchOpenOrders retrieves all pending limit orders for the wallet.
func (c *Client) FetchOpenOrders(ctx context.Context, wallet string) ([]OpenOrder, error) {
	if err := validateAddress(wallet); err != nil {
		return nil, err
	}

	raw, err := c.info.OpenOrders(ctx, wallet)
	if err != nil {
		return nil, &Error{Code: ErrUpstreamFailed, Message: "open_orders: " + err.Error()}
	}

	out := make([]OpenOrder, 0, len(raw))
	for _, o := range raw {
		out = append(out, OpenOrder{
			OrderID:    o.Oid,
			ClientID:   ptrStr(o.Cloid),
			Asset:      o.Coin,
			Side:       sideName(o.Side),
			Size:       o.Size,
			OriginalSz: o.OrigSz,
			LimitPrice: o.LimitPx,
			PlacedAt:   time.UnixMilli(o.Timestamp),
		})
	}

	return out, nil
}

// FetchRecentFills retrieves recent executed trades for the wallet.
// limit caps the number of fills returned (most recent first); 0 means no cap.
func (c *Client) FetchRecentFills(ctx context.Context, wallet string, limit int) ([]Fill, error) {
	if err := validateAddress(wallet); err != nil {
		return nil, err
	}

	raw, err := c.info.UserFills(ctx, hl.UserFillsParams{Address: wallet})
	if err != nil {
		return nil, &Error{Code: ErrUpstreamFailed, Message: "user_fills: " + err.Error()}
	}

	if limit > 0 && len(raw) > limit {
		raw = raw[:limit]
	}

	out := make([]Fill, 0, len(raw))
	for _, f := range raw {
		size := parseFloat(f.Size)
		price := parseFloat(f.Price)
		out = append(out, Fill{
			Asset:       f.Coin,
			Side:        sideName(f.Side),
			Direction:   f.Dir,
			Size:        size,
			Price:       price,
			NotionalUSD: size * price,
			ClosedPnL:   parseFloat(f.ClosedPnl),
			Crossed:     f.Crossed,
			Hash:        f.Hash,
			ExecutedAt:  time.UnixMilli(f.Time),
		})
	}

	return out, nil
}

func sideName(side string) string {
	switch side {
	case "B":
		return "BUY"
	case "A":
		return "SELL"
	}
	return side
}

func ptrStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
