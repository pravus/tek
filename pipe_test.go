package tek_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/pravus/tek"
	"github.com/stretchr/testify/require"
)

func testPipe[I, O any](t *testing.T, i []I, o []O, f func(I) O) {
	pipe := tek.Pipe(f, 1, 1)
	go func() {
		for _, v := range i {
			pipe.I <- v
		}
		close(pipe.I)
	}()
	got := []O{}
	for v := range pipe.O {
		got = append(got, v)
	}
	require.Equal(t, o, got)
}

func TestPipe(t *testing.T) {
	testPipe(t,
		[]int{1, 20, 300},
		[]string{`001`, `020`, `300`},
		func(i int) string {
			return fmt.Sprintf(`%03d`, i)
		},
	)
	type result struct {
		o   int
		err string
	}
	testPipe(t,
		[]string{
			`001`,
			`-20`,
			`-0-`,
		},
		[]result{
			{o: 1},
			{o: -20},
			{err: `strconv.Atoi: parsing "-0-": invalid syntax`},
		},
		func(i string) result {
			if v, err := strconv.Atoi(i); err != nil {
				return result{err: err.Error()}
			} else {
				return result{o: v}
			}
		},
	)
}
