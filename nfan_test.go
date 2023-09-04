package tek_test

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/pravus/tek"
	"github.com/stretchr/testify/require"
)

func testNFan[I, O any](n int, i []I, f func(int, I) O, t func([]O)) {
	nfan := tek.NFan(f, n, 1, 1)
	for _, c := range nfan.I {
		go func(c chan<- I) {
			for _, v := range i {
				c <- v
			}
			close(c)
		}(c)
	}
	o := []O{}
	for v := range nfan.O {
		o = append(o, v)
	}
	t(o)
}

func TestNFan(t *testing.T) {
	testNFan(3,
		[]int{1, 20, 300},
		func(n, v int) string {
			return fmt.Sprintf(`%03d%03d`, n, v)
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
