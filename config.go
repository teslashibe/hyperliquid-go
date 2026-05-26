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
}

// DefaultConfig returns a mainnet config with sensible defaults.
func DefaultConfig() Config {
	return Config{
		Network: NetworkMainnet,
		Timeout: 15 * time.Second,
		Debug:   false,
	}
}
