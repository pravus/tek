package tek

import (
	"sync"
)

func NFan[I, O any](f func(int, I) O, n, capI, capO int) (nfan struct {
	I []chan<- I
	O <-chan O
}) {
	i := make([]chan I, n)
	o := make(chan O, capO)
	w := sync.WaitGroup{}
	w.Add(n)
	for e := 0; e < n; e++ {
		e := e
		c := make(chan I, capI)
		go func() {
			for x := range c {
				o <- f(e, x)
			}
			w.Done()
		}()
		i[e] = c
	}
	go func() {
		w.Wait()
		close(o)
	}()
	nfan.I = make([]chan<- I, n)
	for e := 0; e < n; e++ {
		nfan.I[e] = i[e]
	}
	nfan.O = o
	return nfan
}
