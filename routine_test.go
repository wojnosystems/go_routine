// go_routine Copyright © 2019 Chris Wojno. All rights reserved.

package routine

import "testing"

func TestControl_StopAndJoinError(t *testing.T) {
	wasStarted := false
	wasEnded := false
	r := Go(func(stop <-chan bool) error {
		wasStarted = true
		select {
		case <-stop:
			wasEnded = true
			return nil
		}
	})

	err := r.StopAndJoinError()
	if err != nil {
		t.Error("expected nil error")
	}

	if !wasStarted {
		t.Error("expected routine to start")
	}
	if !wasEnded {
		t.Error("expected routine to end")
	}

	err = r.StopAndJoinError()
	if err != ErrAlreadyStopped {
		t.Error("expected to get ErrAlreadyStopped")
	}
}
