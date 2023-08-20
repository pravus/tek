package tek_test

import (
	"testing"

	"github.com/pravus/tek"
	"github.com/stretchr/testify/require"
)

func testPtr[T any](t *testing.T, v T) {
	p := tek.Ptr(v)
	require.Equal(t, v, *p)
	require.IsType(t, v, *p)
}

func TestPtr(t *testing.T) {
	testPtr(t, int(10))
	testPtr(t, int8(21))
	testPtr(t, int16(32))
	testPtr(t, int32(43))
	testPtr(t, int64(54))

	testPtr(t, float32(1.32))
	testPtr(t, float64(2.34))

	testPtr(t, complex64(1+2i))
	testPtr(t, complex128(2+3i))

	testPtr(t, byte('g'))
	testPtr(t, rune('o'))

	testPtr(t, `go`)

	testPtr(t, true)
	testPtr(t, false)

	{
		v := struct {
			i int
			s string
			b bool
		}{42, `go`, true}
		p := tek.Ptr(v)
		require.Equal(t, v, *p)
		require.IsType(t, v, *p)
		require.Equal(t, v.i, p.i)
		require.IsType(t, v.i, p.i)
		require.Equal(t, v.s, p.s)
		require.IsType(t, v.s, p.s)
		require.Equal(t, v.b, p.b)
		require.IsType(t, v.b, p.b)

		p.i = 24
		p.s = `og`
		p.b = false
		require.NotEqual(t, v, *p)
		require.NotEqual(t, v.i, p.i)
		require.NotEqual(t, v.s, p.s)
		require.NotEqual(t, v.b, p.b)
	}
}
