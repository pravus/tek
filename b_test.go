package tek_test

import (
	"testing"

	"github.com/pravus/tek"
	"github.com/stretchr/testify/require"
)

func testB[T any](t *testing.T, vt, vf T) {
	require.Equal(t, tek.B(true, vt, vf), vt)
	require.Equal(t, tek.B(false, vt, vf), vf)
}

func TestB(t *testing.T) {
	testB(t, true, false)
	testB(t, []byte(`t`), []byte(`f`))
	testB(t, 'y', 'n')
	testB(t, `y`, `n`)
	testB(t, 0, 1)
	testB(t, 0e0, 1e0)
	testB(t, 0i, 1i)
}
