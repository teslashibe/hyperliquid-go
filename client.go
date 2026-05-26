package hyperliquid

import (
	"context"
	"strings"

	hl "github.com/sonirico/go-hyperliquid"
)

// Client is a thin wrapper over sonirico/go-hyperliquid focused on
// read-only account state and order queries.
type Client struct {
	cfg  Config
	info *hl.Info
}

// New constructs a Client. The underlying SDK fetches metadata at
// construction time, so this performs network IO.
func New(cfg Config) (*Client, error) {
	baseURL := hl.MainnetAPIURL
	if cfg.Network == NetworkTestnet {
		baseURL = hl.TestnetAPIURL
	}

	info := hl.NewInfo(context.Background(), baseURL, true, nil, nil, nil)

	return &Client{
		cfg:  cfg,
		info: info,
	}, nil
}

// Close releases resources. Currently a no-op (no persistent connection).
func (c *Client) Close() error { return nil }

// validateAddress performs a minimal sanity check on a wallet address.
func validateAddress(addr string) error {
	a := strings.TrimSpace(addr)
	if !strings.HasPrefix(a, "0x") || len(a) != 42 {
		return &Error{Code: ErrInvalidWallet, Message: "wallet must be 42-char 0x-prefixed hex"}
	}
	return nil
}
