package hyperliquid

import "time"

// AccountState is the normalized perpetuals account snapshot.
type AccountState struct {
	Wallet            string     `json:"wallet"`
	AccountValueUSD   float64    `json:"account_value_usd"`
	WithdrawableUSD   float64    `json:"withdrawable_usd"`
	TotalNotionalUSD  float64    `json:"total_notional_usd"`
	TotalMarginUsed   float64    `json:"total_margin_used_usd"`
	MarginUtilization float64    `json:"margin_utilization_pct"`
	Positions         []Position `json:"positions"`
	FetchedAt         time.Time  `json:"fetched_at"`
}

// Position is a single open perpetual position with derived fields computed.
type Position struct {
	Asset             string  `json:"asset"`
	Direction         string  `json:"direction"`
	Size              float64 `json:"size"`
	EntryPrice        float64 `json:"entry_price"`
	PositionValueUSD  float64 `json:"position_value_usd"`
	UnrealizedPnLUSD  float64 `json:"unrealized_pnl_usd"`
	UnrealizedPnLPct  float64 `json:"unrealized_pnl_pct"`
	LeverageType      string  `json:"leverage_type"`
	LeverageValue     int     `json:"leverage_value"`
	LiquidationPrice  float64 `json:"liquidation_price"`
	LiquidationDistPct float64 `json:"liquidation_distance_pct"`
	MarginUsedUSD     float64 `json:"margin_used_usd"`
	MaxLeverage       int     `json:"max_leverage,omitempty"`
	FundingPaidTotal  float64 `json:"funding_paid_total_usd,omitempty"`
}

// OpenOrder is a normalized pending limit order.
type OpenOrder struct {
	OrderID    int64     `json:"order_id"`
	ClientID   string    `json:"client_id,omitempty"`
	Asset      string    `json:"asset"`
	Side       string    `json:"side"`
	Size       float64   `json:"size"`
	OriginalSz float64   `json:"original_size"`
	LimitPrice float64   `json:"limit_price"`
	PlacedAt   time.Time `json:"placed_at"`
}

// Fill is a normalized executed trade.
type Fill struct {
	Asset      string    `json:"asset"`
	Side       string    `json:"side"`
	Direction  string    `json:"direction"`
	Size       float64   `json:"size"`
	Price      float64   `json:"price"`
	NotionalUSD float64  `json:"notional_usd"`
	ClosedPnL  float64   `json:"closed_pnl_usd"`
	Crossed    bool      `json:"crossed"`
	Hash       string    `json:"tx_hash"`
	ExecutedAt time.Time `json:"executed_at"`
}
