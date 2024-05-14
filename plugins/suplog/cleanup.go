package suplog

import (
	"strings"

	"github.com/bugsnag/bugsnag-go"
)

// BugsnagCleanup should be used with bugsnag.OnBeforeNotify.
func BugsnagCleanup(e *bugsnag.Event, c *bugsnag.Configuration) error {
	var pos = -1
	for idx, st_item := range e.Stacktrace {
		switch {
		case strings.Contains(st_item.File, "xlab/catcher"):
			pos = idx
			continue
		case strings.HasPrefix(st_item.File, "runtime"):
			pos = idx
			continue
		}
		if pos >= 0 {
			e.Stacktrace = e.Stacktrace[pos+1:]
			break
		}
	}
	e.ErrorClass = e.Message
	return nil
}
