package hyperliquid

import "time"

// Network identifiers.
const (
	NetworkMainnet = "mainnet"
	NetworkTestnet = "testnet"
)

// Config controls client behavior.
type Config struct {
	Network string        // "mainnet" or "testnet"
	Timeout time.Duration // HTTP timeout
	Debug   bool          // verbose logging from underlying SDK

	// Wallet is an optional default 0x address. When set, callers (e.g. the
	// MCP tools) may resolve an empty wallet argument to this address, so a
	// host that stores one wallet per user need not pass it on every call.
	// It does not restrict which wallet a call may target; an explicit
	// argument always wins.
	Wallet string
}

// DefaultConfig returns a mainnet config with sensible defaults.
func DefaultConfig() Config {
	return Config{
		Network: NetworkMainnet,
		Timeout: 15 * time.Second,
		Debug:   false,
	}
}
