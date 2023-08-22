package tek_test

import (
	"fmt"
	"testing"
	"strings"

	"github.com/pravus/tek"
	"github.com/stretchr/testify/require"
)

var meprErr = fmt.Errorf("told to return an error")

func TestMepr(t *testing.T) {
	{
		v, e := tek.Mepr(nil, func(a any) (any, error) {
			return a, nil
		})
		require.NoError(t, e)
		require.Equal(t, []any{}, v)
	}
	{
		v, e := tek.Mepr([]string{"foo", "bar", "baz"}, func(s string) (string, error) {
			return strings.ToUpper(s), nil
		})
		require.NoError(t, e)
		require.Equal(t, []string{"FOO", "BAR", "BAZ"}, v)
	}
	{
		v, e := tek.Mepr([]int{-1, 0, +1}, func(i int) (string, error) {
			return fmt.Sprintf(`%d`, i), nil
		})
		require.NoError(t, e)
		require.Equal(t, []string{"-1", "0", "1"}, v)
	}
	{
		v, e := tek.Mepr([]any{"foo", "bar", "baz"}, func(a any) (any, error) {
			return strings.ToUpper(a.(string)), nil
		})
		require.NoError(t, e)
		require.Equal(t, []any{"FOO", "BAR", "BAZ"}, v)
	}
	{
		v, e := tek.Mepr([]int{-1, 0, +1}, func(i int) (any, error) {
			return fmt.Sprintf(`%d`, i), nil
		})
		require.NoError(t, e)
		require.Equal(t, []any{"-1", "0", "1"}, v)
	}
	{
		v, e := tek.Mepr([]any{-1, 0, +1}, func(a any) (string, error) {
			return fmt.Sprintf(`%d`, a.(int)), nil
		})
		require.NoError(t, e)
		require.Equal(t, []string{"-1", "0", "1"}, v)
	}
	{
		_, e := tek.Mepr([]any{nil, nil, meprErr, nil}, func(a any) (string, error) {
			if a != nil {
				return ``, a.(error)
			}
			return `MEPR`, nil
		})
		require.Error(t, e)
		require.ErrorIs(t, e, meprErr)
	}
}
