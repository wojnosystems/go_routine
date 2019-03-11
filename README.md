# Overview

go_routine is a unified method for managing and synchronizing with Go's routines. This library doesn't do much on it's own, but it makes it easier to provide a uniform way of stopping and signalling errors between routines.

When you call go_routine.Go, it forks a go routine for you and runs the method passed into Go. When the invoking routine (or any routine, just make sure it happens only once) calls Stop or StopAndJoinError, the go-routine will be signaled via that channel.

If you call Stop, the signal will be sent, but it won't block to wait for the routine to stop. However, if you call StopAndJoinError, Stop will be called, but the error value returned by the routine passed to Go will be returned as well. This method will wait until the routine completes.

# Copyright

Copyright © 2019 Chris Wojno. All rights reserved.

No Warranties. Use this software at your own risk.

# License

Attribution 4.0 International https://creativecommons.org/licenses/by/4.0/