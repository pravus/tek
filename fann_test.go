package tek_test

import (
	"fmt"
	"strconv"
	"sync"
	"testing"

	"github.com/pravus/tek"
	"github.com/stretchr/testify/require"
)

func testFanN[I, O any](n int, i []I, f func(int, I) O, t func(int, []O)) {
	fann := tek.FanN(f, n, 1, 1)
	go func() {
		for _, v := range i {
			fann.I <- v
		}
		close(fann.I)
	}()
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i, o := range fann.O {
		go func(i int, o <-chan O) {
			got := []O{}
			for v := range o {
				got = append(got, v)
			}
			wg.Done()
			t(i, got)
		}(i, o)
	}
	wg.Wait()
}

func TestFanN(t *testing.T) {
	testFanN(3,
		[]int{1, 20, 300},
		func(n, v int) string {
			return fmt.Sprintf(`%03d%03d`, n, v)
		},
		func(n int, got []string) {
			want := []string{
				`00` + strconv.Itoa(n) + `001`,
				`00` + strconv.Itoa(n) + `020`,
				`00` + strconv.Itoa(n) + `300`,
			}
			require.Equal(t, want, got)
		},
	)
}
