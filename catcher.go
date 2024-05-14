// Package catcher provides utilites to gracefully handle crap software.
package catcher

import (
	"fmt"
	"sync"
)

var (
	mu        sync.RWMutex
	notifiers []NotifierFunc
)

// RegisterNotifiers adds all listed notifiers to a global list, each of them
// would be invoked upon an error occurs.
func RegisterNotifiers(fn ...NotifierFunc) {
	mu.Lock()
	notifiers = append(notifiers, fn...)
	mu.Unlock()
}

// Catch must be used together with defer to catch panics from suspicious functions.
// All listed receivers and all the global notifiers will be invoked with the
// error itself, function name and a stack trace.
func Catch(recv ...Receiver) {
	if panicData := recover(); panicData != nil {
		caller := getCaller(5) + " <= " + getCaller(6)
		if err, ok := panicData.(error); ok {
			if len(recv) > 0 {
				stack := getStack()
				for _, r := range recv {
					r.RecvError(err, nil, caller, stack)
				}
			}
			return
		}
		err := fmt.Errorf("%+v", panicData)

		if len(recv) > 0 {
			// invoke receivers
			stack := getStack()
			for _, r := range recv {
				r.RecvPanic(err, nil, caller, stack)
			}
		}
		mu.RLock()
		defer mu.RUnlock()
		if len(notifiers) > 0 {
			wg := new(sync.WaitGroup)
			// notify all registered notifiers
			for _, fn := range notifiers {
				wg.Add(1)
				go func(fn NotifierFunc) {
					fn(err)
					wg.Done()
				}(fn)
			}
			wg.Wait()
		}
	}
}

// CatchWithContext must be used together with defer to catch panics from suspicious functions.
// All listed receivers and all the global notifiers will be invoked with the error itself,
// function name and a stack trace, along with the provided context and meta.
func CatchWithContext(context, meta interface{}, recv ...Receiver) {
	if panicData := recover(); panicData != nil {
		caller := getCaller(5) + " <= " + getCaller(6)
		if err, ok := panicData.(error); ok {
			if len(recv) > 0 {
				stack := getStack()
				for _, r := range recv {
					r.RecvError(err, context, caller, stack)
				}
			}
			return
		}
		err := fmt.Errorf("%+v", panicData)

		if len(recv) > 0 {
			stack := getStack()
			// invoke receivers
			for _, r := range recv {
				r.RecvPanic(err, context, caller, stack)
			}
		}
		mu.RLock()
		defer mu.RUnlock()
		if len(notifiers) > 0 {
			wg := new(sync.WaitGroup)
			// notify all registered notifiers
			for _, fn := range notifiers {
				wg.Add(1)
				go func(fn NotifierFunc) {
					fn(err, context, meta)
					wg.Done()
				}(fn)
			}
			wg.Wait()
		}
	}
}

// NotifierFunc is a func that may be used as a global notifier (via RegisterNotifiers()).
type NotifierFunc func(err error, rawData ...interface{}) error

// Receiver handles panics and errors thrown with a panic.
type Receiver interface {
	RecvError(err error, context interface{}, caller string, stack []byte)
	RecvPanic(err error, context interface{}, caller string, stack []byte)
}

// Error has additional context and stack that can be accessed outside receivers;
// by storing it to a variable using the RecvError receiver.
type Error struct {
	Err     error
	Context interface{}
	Caller  string
	Stack   []byte
}

func (e Error) Error() string {
	return e.Err.Error()
}
