package routine

// Stopper is a routine that can be stopped. Stopping is
// not synchronous and Stop should not wait for the
// routine to complete
type Stopper interface {
	// Tells the go-routines to stop, but does not wait for
	// them to do so
	Stop()
}

// StopJoiner is a routine that can be stopped and then
// synchronously joined to wait for the first error to be encountered
// to return from the routine started with routine.Start
type StopJoiner interface {
	// StopAndJoinError will tell the go-routines to stop
	// but will wait for them, returning the first error returned
	StopAndJoinError() error
}

// StopWithJoiner is the same as Stopper + StopJoiner
// This is a convenience method
type StopWithJoiner interface {
	Stopper
	StopJoiner
}
