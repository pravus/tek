package tek_test

import (
	"testing"

	"github.com/pravus/tek"
	"github.com/stretchr/testify/require"
)

func testF(t *testing.T, f func(), e bool) {
	v := tek.F(f)
	require.Equal(t, v, e)
}

func TestF(t *testing.T) {
	f := false
	if v := f; v {
		t.Error(`trapped in the ante logic error`)
	} else if tek.F(func() {
		testF(t, func() {
			f = true
		}, v)
	}) {
		t.Error(`trapped in a maze of logic errors`)
	} else if tek.F(nil) {
		t.Error(`trapped in a maze of logic errors`)
	} else if err := error(nil); err != nil {
		t.Error(`trapped in the post logic error`)
	} else {
		require.Equal(t, v, false)
	}
	require.Equal(t, f, true)
}
