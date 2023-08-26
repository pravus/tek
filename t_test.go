package tek_test

import (
	"testing"

	"github.com/pravus/tek"
	"github.com/stretchr/testify/require"
)

func testT[T any](t *testing.T, vt, vf T) {
	ft := func() T { return vt }
	ff := func() T { return vf }
	require.Equal(t, tek.T(true, ft, ff), vt)
	require.Equal(t, tek.T(false, ft, ff), vf)
}

func TestT(t *testing.T) {
	testT(t, true, false)
	testT(t, []byte(`t`), []byte(`f`))
	testT(t, 'y', 'n')
	testT(t, `y`, `n`)
	testT(t, 0, 1)
	testT(t, 0e0, 1e0)
	testT(t, 0i, 1i)
}
