package tek

func FanN[I, O any](f func(int, I) O, n, capI, capO int) (fann struct {
	I chan<- I
	O []<-chan O
}) {
	i := make(chan I, capI)
	o := make([]chan<- O, n)
	fann.I = i
	fann.O = make([]<-chan O, n)
	for e := 0; e < n; e++ {
		c := make(chan O, capO)
		fann.O[e], o[e] = c, c
	}
	go func() {
		for x := range i {
			for e, o := range o {
				o <- f(e, x)
			}
		}
		for _, o := range o {
			close(o)
		}
	}()
	return fann
}
