package main

import (
	"errors"
	"log"
	"os"

	"github.com/xlab/catcher"
)

func init() {
	log.SetFlags(log.Lshortfile)
	// bugsnag.OnBeforeNotify(suplog.BugsnagCleanup)
	// catcher.RegisterNotifiers(bugsnag.Notify)
}

func main() {
	defer catcher.Catch(
		catcher.RecvWrite(os.Stderr, true),
	)

	if err := safeCall(); err != nil {
		log.Println("[ERR] safeCall failed with:", err)
	}

	suspiciousFunc()
}

// suspiciousFunc will definitely panic. Usually this kind of functions
// panic only on Saturdays or holidays, but for test simplicity this one
// will panic 100% of the time.
func suspiciousFunc() {
	panic("sorry pls")
}

// safeCall is an example of a function that uses two receivers.
// First one will put the panic message into the error value;
// second one will yield the message to the stderr without the stracktrace.
func safeCall() (err error) {
	defer catcher.Catch(
		catcher.RecvError(&err),
		catcher.RecvWrite(os.Stderr),
	)

	overthinkFunc()
	return
}

// overthinkFunc will throw an error like it was Java.
func overthinkFunc() {
	panic(errors.New("must catch me"))
}
