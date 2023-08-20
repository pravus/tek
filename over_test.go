package tek_test

import (
	"fmt"
	"testing"

	"github.com/pravus/tek"
	"github.com/stretchr/testify/require"
)

var overErr = fmt.Errorf("told to return an error")

func testOver[T any](t *testing.T, s []T, b func(T) error, e error) {
	v := tek.Over(s, b)
	require.Equal(t, v, e)
}

func testOverAny(value any) error {
	return overErr
}

func testOverBool(value bool) error {
	if !value {
		return overErr
	}
	return nil
}

func testOverByte(value byte) error {
	if value == byte(0x2d) {
		return overErr
	}
	return nil
}

func testOverComplex128(value complex128) error {
	if value == -1i {
		return overErr
	}
	return nil
}

func testOverFloat64(value float64) error {
	if value == 1e-1 {
		return overErr
	}
	return nil
}

func testOverInt(value int) error {
	if value == 0xdead_beef {
		return overErr
	}
	return nil
}

func testOverRune(value rune) error {
	if value == '\u002d' {
		return overErr
	}
	return nil
}

func testOverString(value string) error {
	if value == `-` {
		return overErr
	}
	return nil
}

func TestOver(t *testing.T) {
	testOver(t, nil, testOverAny, nil)
	testOver(t, []any{}, testOverAny, nil)
	testOver(t, []bool{true}, testOverBool, nil)
	testOver(t, []bool{false}, testOverBool, overErr)
	testOver(t, []byte(`+`), testOverByte, nil)
	testOver(t, []byte(`-`), testOverByte, overErr)
	testOver(t, []complex128{+1i}, testOverComplex128, nil)
	testOver(t, []complex128{-1i}, testOverComplex128, overErr)
	testOver(t, []float64{0e+1}, testOverFloat64, nil)
	testOver(t, []float64{1e-1}, testOverFloat64, overErr)
	testOver(t, []int{0x0000_0000}, testOverInt, nil)
	testOver(t, []int{0xdead_beef}, testOverInt, overErr)
	testOver(t, []rune{'+'}, testOverRune, nil)
	testOver(t, []rune{'-'}, testOverRune, overErr)
	testOver(t, []string{`+`}, testOverString, nil)
	testOver(t, []string{`-`}, testOverString, overErr)
}
