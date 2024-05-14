package suplog

import (
	"encoding/json"

	"github.com/xlab/catcher"
	"github.com/xlab/suplog"
)

// RecvSuplog logs the error caught using the suplog logger.
// See https://github.com/xlab/suplog
func RecvSuplog(logger suplog.Logger) catcher.Receiver {
	return suplogReceiver{
		logger: logger,
	}
}

type suplogReceiver struct {
	logger suplog.Logger
}

func (s suplogReceiver) RecvError(err error, context interface{}, caller string, stack []byte) {
	logger := s.logger

	if len(caller) > 0 {
		logger = logger.WithField("caller", caller)
	}

	if context != nil {
		vv, _ := json.Marshal(context)
		logger = logger.WithField("context", string(vv))
	}

	if len(stack) > 0 {
		logger = logger.WithField("stacktrace", string(stack))
	}

	logger.WithError(err).Error("catcher received an error")
}

func (s suplogReceiver) RecvPanic(err error, context interface{}, caller string, stack []byte) {
	logger := s.logger

	if len(caller) > 0 {
		logger = logger.WithField("caller", caller)
	}

	if context != nil {
		vv, _ := json.Marshal(context)
		logger = logger.WithField("context", string(vv))
	}

	if len(stack) > 0 {
		logger = logger.WithField("stacktrace", string(stack))
	}

	logger.WithError(err).Error("catcher received a panic")
}
