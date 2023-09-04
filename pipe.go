package tek

func Pipe[I, O any](f func(I) O, capI, capO int) (pipe struct {
	I chan<- I
	O <-chan O
}) {
	i := make(chan I, capI)
	o := make(chan O, capO)
	go func() {
		for x := range i {
			o <- f(x)
		}
		close(o)
	}()
	pipe.I, pipe.O = i, o
	return pipe
}
