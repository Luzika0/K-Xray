// Package engine defines the shared mobile engine lifecycle contract.
package engine

import "errors"

type Mode string

const (
	ModeXray       Mode = "xray"
	ModePingTunnel Mode = "pingtunnel"
)

var ErrNotImplemented = errors.New("engine adapter is not integrated")

// Start is the single entry point intended for gomobile/native bindings.
// tunFD is owned by the platform VPN service; implementations must duplicate
// it before handing it to a native engine and must not close the caller's FD.
// This stub intentionally fails closed until real adapters are integrated.
func Start(mode Mode, tunFD int, configJSON string) error {
	return ErrNotImplemented
}

// Stop is idempotent by contract. The current scaffold has no running engine.
func Stop() error {
	return ErrNotImplemented
}
