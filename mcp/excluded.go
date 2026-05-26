package hyperliquidmcp

// Excluded lists *hyperliquid.Client methods intentionally not exposed.
var Excluded = map[string]string{
	"Close": "lifecycle owned by the host process; not an agent action",
}
