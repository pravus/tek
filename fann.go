package tek

func FanN[I, O any](f func(int, I) O, n, capI, capO int) (fann struct {
	I chan<- (I)
	O []<-chan (O)
}) {
	i := make(chan (I), capI)
	o := make([]chan (O), n)
	for e := 0; e < n; e++ {
		o[e] = make(chan (O), capO)
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
	fann.I = i
	fann.O = make([]<-chan (O), n)
	for e := 0; e < n; e++ {
		fann.O[e] = o[e]
	}
	return fann
}
