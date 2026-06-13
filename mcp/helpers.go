package hyperliquidmcp

import (
	"errors"
	"strings"

	"github.com/teslashibe/hyperliquid-go"
	"github.com/teslashibe/mcptool"
)

// resolveWallet picks the explicit per-call wallet when supplied, else falls
// back to the client's configured default. Returns an invalid_input error
// when neither is present so the agent knows to ask for (or have the user
// connect) a wallet.
func resolveWallet(c *hyperliquid.Client, arg string) (string, error) {
	if w := strings.TrimSpace(arg); w != "" {
		return w, nil
	}
	if w := c.DefaultWallet(); w != "" {
		return w, nil
	}
	return "", &mcptool.Error{
		Code:    "invalid_input",
		Message: "no wallet: pass a 0x address or connect one in Settings -> Hyperliquid",
	}
}

// wrapErr converts a hyperliquid-package error into a structured mcptool.Error.
func wrapErr(err error, op string) error {
	if err == nil {
		return nil
	}
	var hlErr *hyperliquid.Error
	if errors.As(err, &hlErr) {
		return &mcptool.Error{
			Code:    hlErr.Code,
			Message: op + ": " + hlErr.Message,
		}
	}
	return err
}
