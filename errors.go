package hyperliquid

import "fmt"

// Stable error codes for programmatic handling.
const (
	ErrInvalidWallet   = "invalid_wallet"
	ErrUpstreamFailed  = "upstream_failed"
	ErrUpstreamTimeout = "upstream_timeout"
	ErrParseFailure    = "parse_failure"
)

// Error is a structured error with a stable code.
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}
