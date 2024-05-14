package catcher

import (
	"fmt"
	"io"
	"log"
	"os"
)

// RecvError puts the error catched in the variable by pointer. If useSimple is set to true,
// the error would not contain anything except the message.
func RecvError(err *error, useSimple ...bool) Receiver {
	s := len(useSimple) > 0 && useSimple[0]
	return errReceiver{err: err, useSimple: s}
}

type errReceiver struct {
	err       *error
	useSimple bool
}

func (e errReceiver) RecvError(err error, context interface{}, caller string, stack []byte) {
	if e.useSimple {
		*e.err = err
		return
	}
	*e.err = &Error{
		Err:     err,
		Caller:  caller,
		Context: context,
		Stack:   stack,
	}
}

func (e errReceiver) RecvPanic(err error, context interface{}, caller string, stack []byte) {
	if e.useSimple {
		*e.err = err
		return
	}
	*e.err = &Error{
		Err:     err,
		Context: context,
		Caller:  caller,
		Stack:   stack,
	}
}

// RecvDie forces app to exit with the provided exit code. If onPanic is set to true,
// the app will die only on wild panics. N.B. panic(errors.New(...)) is not a wild panic.
func RecvDie(ret int, onPanic ...bool) Receiver {
	p := len(onPanic) > 0 && onPanic[0]
	return dieReceiver{ret: ret, onPanic: p}
}

type dieReceiver struct {
	onPanic bool
	ret     int
}

func (d dieReceiver) RecvError(err error, context interface{}, caller string, stack []byte) {
	if !d.onPanic { // die on an error too
		os.Exit(d.ret)
	}
}

func (d dieReceiver) RecvPanic(err error, context interface{}, caller string, stack []byte) {
	os.Exit(d.ret)
}

// RecvWrite writes the error caught to the writer, presumably the os.Stderr.
// If writeStack is set to true, a stacktrace will be written too.
func RecvWrite(wr io.Writer, writeStack ...bool) Receiver {
	s := len(writeStack) > 0 && writeStack[0]
	return writeReceiver{wr: wr, writeStack: s}
}

type writeReceiver struct {
	wr         io.Writer
	writeStack bool
}

func (w writeReceiver) RecvError(err error, context interface{}, caller string, stack []byte) {
	var callPrefix string
	if w.writeStack {
		callPrefix = fmt.Sprintf("%s: ", caller)
	}
	if context != nil {
		fmt.Fprintf(w.wr, "caught error: %s%s (has context)\n", callPrefix, err.Error())
	}
	fmt.Fprintf(w.wr, "caught error: %s%s\n", callPrefix, err.Error())
	if w.writeStack {
		fmt.Fprintf(w.wr, "stacktrace: %s\n", stack)
	}
}

func (w writeReceiver) RecvPanic(err error, context interface{}, caller string, stack []byte) {
	fmt.Fprintf(w.wr, "caught panic: %s", err.Error())
	if len(caller) > 0 {
		fmt.Fprintf(w.wr, " from %s", caller)
	}
	if context != nil {
		fmt.Fprint(w.wr, " (has context)")
	}
	fmt.Fprint(w.wr, "\n")
	if w.writeStack {
		fmt.Fprintf(w.wr, "stacktrace: %s\n", stack)
	}
}

// RecvLog logs the error caught using the standard logger.
// If writeStack is set to true, a stacktrace will be written too.
func RecvLog(writeStack ...bool) Receiver {
	s := len(writeStack) > 0 && writeStack[0]
	return logReceiver{writeStack: s}
}

type logReceiver struct {
	writeStack bool
}

func (w logReceiver) RecvError(err error, context interface{}, caller string, stack []byte) {
	var callPrefix string
	if w.writeStack {
		callPrefix = fmt.Sprintf("%s: ", caller)
	}
	if context != nil {
		fmt.Printf("caught error: %s%s (has context)\n", callPrefix, err.Error())
	}
	log.Printf("caught error: %s%s\n", callPrefix, err.Error())
	if w.writeStack {
		log.Printf("stacktrace: %s\n", stack)
	}
}

func (w logReceiver) RecvPanic(err error, context interface{}, caller string, stack []byte) {
	fmt.Printf("caught panic: %s", err.Error())
	if len(caller) > 0 {
		fmt.Printf(" from %s", caller)
	}
	if context != nil {
		fmt.Print(" (has context)")
	}
	fmt.Println()
	if w.writeStack {
		fmt.Printf("stacktrace: %s\n", stack)
	}
}
