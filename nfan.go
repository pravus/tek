package tek

import (
	"sync"
)

func NFan[I, O any](f func(int, I) O, n, capI, capO int) (nfan struct {
	I []chan<- I
	O <-chan O
}) {
	i := make([]<-chan I, n)
	o := make(chan O, capO)
	nfan.I = make([]chan<- I, n)
	nfan.O = o
	w := sync.WaitGroup{}
	w.Add(n)
	for e := 0; e < n; e++ {
		c := make(chan I, capI)
		nfan.I[e], i[e] = c, c
		go func(e int, c <-chan I) {
			for x := range c {
				o <- f(e, x)
			}
			w.Done()
		}(e, i[e])
	}
	go func() {
		w.Wait()
		close(o)
	}()
	return nfan
}
