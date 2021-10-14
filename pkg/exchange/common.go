package exchange

import (
	"github.com/kataras/golog"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/tools/state"
)

var (
	logger = golog.Child("protocols/exchange")
)

// checkParams Checks if Non-nil Parameters were passed
func checkParams(host *host.SNRHost, em *state.Emitter) error {
	if host == nil {
		logger.Error("Host provided is nil", ErrParameters)
		return ErrParameters
	}
	if em == nil {
		logger.Error("Emitter provided is nil", ErrParameters)
		return ErrParameters
	}
	return host.HasRouting()
}