package tek

import (
	"sync"
)

func Ecks[I, O any](f func(int, I, []chan<- O), nI, capI, nO, capO int) (ecks struct {
	I []chan<- I
	O []<-chan O
}) {
	i := make([]chan I, nI)
	ecks.I = make([]chan<- I, nI)
	for e := 0; e < nI; e++ {
		ci := make(chan I, capI)
		ecks.I[e], i[e] = ci, ci
	}
	o := make([]chan O, nO)
	fo := make([]chan<- O, nO)
	ecks.O = make([]<-chan O, nO)
	for e := 0; e < nO; e++ {
		co := make(chan O, capO)
		ecks.O[e], fo[e], o[e] = co, co, co
	}
	w := sync.WaitGroup{}
	w.Add(nI)
	for e := 0; e < nI; e++ {
		go func(e int, i chan I) {
			for x := range i {
				f(e, x, fo)
			}
			w.Done()
		}(e, i[e])
	}
	go func() {
		w.Wait()
		for _, o := range o {
			close(o)
		}
	}()
	return ecks
}
