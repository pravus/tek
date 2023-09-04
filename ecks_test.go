package tek_test

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"testing"

	"github.com/pravus/tek"
	"github.com/stretchr/testify/require"
)

func testEcks[I, O any](nI, nO int, i []I, f func(int, I, []chan<- O), t func([]O)) {
	ecks := tek.Ecks(f, nI, 1, nO, 1)
	for _, c := range ecks.I {
		go func(c chan<- I) {
			for _, v := range i {
				c <- v
			}
			close(c)
		}(c)
	}
	all := make([][]O, nO)
	wg := sync.WaitGroup{}
	wg.Add(nO)
	for e, o := range ecks.O {
		all[e] = []O{}
		go func(e int, o <-chan O) {
			for v := range o {
				all[e] = append(all[e], v)
			}
			wg.Done()
		}(e, o)
	}
	wg.Wait()
	got := []O{}
	for _, o := range all {
		got = append(got, o...)
	}
	t(got)
}

func TestEcks(t *testing.T) {
	// fann
	testEcks(1, 3,
		[]int{1, 20, 300},
		func(n, v int, o []chan<- string) {
			for _, o := range o {
				o <- fmt.Sprintf(`%03d%03d`, n, v)
			}
		},
		func(got []string) {
			want := []string{
				`000001`, `000020`, `000300`,
				`000001`, `000020`, `000300`,
				`000001`, `000020`, `000300`,
			}
			require.Equal(t, want, got)
		},
	)

	// nfan
	testEcks(3, 1,
		[]int{1, 20, 300},
		func(n, v int, o []chan<- string) {
			for _, o := range o {
				o <- fmt.Sprintf(`%03d%03d`, n, v)
			}
		},
		func(got []string) {
			sort.Slice(got, func(one, two int) bool {
				return strings.Compare(got[one], got[two]) < 0
			})
			want := []string{
				`000001`, `000020`, `000300`,
				`001001`, `001020`, `001300`,
				`002001`, `002020`, `002300`,
			}
			require.Equal(t, want, got)
		},
	)
}
