package hyperliquidmcp

import (
	"errors"

	"github.com/teslashibe/hyperliquid-go"
	"github.com/teslashibe/mcptool"
)

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
