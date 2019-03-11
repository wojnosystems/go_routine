// go_routine Copyright © 2019 Chris Wojno. All rights reserved.

package go_routine

import (
	"errors"
)

// routine.control is a way to control a go-routine to be started and stopped
// This is not exported because it should be created with a call to "Go" and be terminated with a call to Stop or StopAndJoinError
type control struct {
	// stopChan is how we signal to the function passed to "Go" that it needs to exit as soon as possible
	stopChan chan bool

	// waitForError is optionally set when StopAndJoinError is called
	waitForError *chan error
}

var ErrAlreadyStopped = errors.New("control has already been stopped")

// Go starts the routine(s) on their own go-Routine
// @param routine is a function passed by the user of the library
func Go(routine func(stop <-chan bool) error) StopWithJoiner {
	c := &control{
		stopChan: make(chan bool, 1),
	}
	// This forks the go-routine so that start can return immediately
	go func(ctrl *control) {
		err := routine(ctrl.stopChan)
		if ctrl.waitForError != nil {
			*ctrl.waitForError <- err
		}
	}(c)
	return c
}

// Tells the go-routines to stop, but does not wait for
// them to do so
func (c *control) Stop() {
	c.stopChan <- true
}

// StopAndJoinError will tell the go-routines to stop
// but will wait for them, returning the first error returned
func (c *control) StopAndJoinError() (err error) {
	if c.waitForError != nil {
		return ErrAlreadyStopped
	}
	ch := make(chan error, 1)
	c.waitForError = &ch
	c.Stop()
	err = <-*c.waitForError
	// We're done with the channel
	close(*c.waitForError)
	return err
}
