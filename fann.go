package tek

func FanN[I, O any](f func(int, I) O, n, capI, capO int) (fann struct {
	I chan<- I
	O []<-chan O
}) {
	i := make(chan I, capI)
	o := make([]chan O, n)
	fann.I = i
	fann.O = make([]<-chan O, n)
	for i := 0; i < n; i++ {
		c := make(chan O, capO)
		fann.O[i], o[i] = c, c
	}
	go func() {
		for x := range i {
			for i, o := range o {
				o <- f(i, x)
			}
		}
		for _, o := range o {
			close(o)
		}
	}()
	return fann
}
