package catcher

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCatch(t *testing.T) {
	assert := assert.New(t)
	assert.Panics(func() {
		suspiciousFunc()
	})
	assert.NotPanics(func() {
		defer Catch()
		suspiciousFunc()
	})
	err := func() (err error) {
		defer Catch(RecvError(&err))
		suspiciousFunc()
		return
	}()
	assert.Error(err)
	e, ok := err.(*Error)
	if assert.True(ok) {
		assert.Equal("sorry pls", e.Error())
		assert.Equal("TestCatch at catcher_test.go:24 <= tRunner at testing.go:1689", string(e.Caller))
	}
}

// ExampleCatch shows how to treat suspicious funcs properly.
func ExampleCatch() {
	defer Catch( // if this example panics (oh sure it would), just output the
		// panic message and its stacktrace to the stderr.
		RecvWrite(os.Stderr, true),
	)

	// SafeCall is an example of a function that uses two receivers.
	// First one will put the panic message into the error value;
	// second one will yield the message to the stderr without the stracktrace.
	SafeCall := func() (err error) {
		defer Catch(
			RecvError(&err),
			RecvWrite(os.Stderr),
		)

		suspiciousFunc()
		return
	}
	// treat suspiciousFunc like a normal func that may return an error
	if err := SafeCall(); err != nil {
		log.Println("[ERR] SafeCall failed with:", err)
	}

	// don't be afraid to call this func anymore
	suspiciousFunc()
}

// suspiciousFunc will definitely panic. Usually this kind of functions
// panic only on Saturdays or holidays, but for test simplicity this one
// will panic 100% of the time.
func suspiciousFunc() {
	panic("sorry pls")
}
