package terminator

/*
Encapuslate a mechanism to terminate the execution of a goroutine.  In general,
this interface wraps a signal, external to the given goroutine's context, to
forward the termination request to the goroutine.  However, this interface can
be decorated with a signal local to the goroutine's context.

IsNot() - Implements a non-blocking function to test for a termination signal.

Chan() - Exposes a blocking channel mechanism that broadcasts a termination
signal to the goroutine. Use on golang "select" statement to cause blocking instead of
spinning the CPU.
*/
type I interface {
	IsNot() bool
	Chan() chan bool
}

/*
Extends the I interface by requiring the signaled goroutine to acknowledge it's
terimatation.  The goroute signals its termination by encoding the Done() method
in one of its defer statements.   This permits goroutines, that depend on the
termination of one or more other goroutines, to detect the termination of the
goroutines that they are observing.
*/
type Isync interface {
	I
	Add(int)
	Done()
	Wait()
}
