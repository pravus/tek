package tek_test

import (
	"testing"

	"github.com/pravus/tek"
	"github.com/stretchr/testify/require"
)

func TestGlue(t *testing.T) {
	type test struct {
		s []string
		d string
		e string
	}
	for _, test := range []test{
		{s: nil, d: ``, e: ``},
		{s: []string{}, d: ``, e: ``},
		{s: []string{`pony`}, d: ``, e: `pony`},
		{s: []string{`pony`, `meat`}, d: ``, e: `ponymeat`},
		{s: []string{`pony`, `meat`, `hoof`}, d: ``, e: `ponymeathoof`},
		{s: nil, d: `:`, e: ``},
		{s: []string{}, d: `:`, e: ``},
		{s: []string{`pony`}, d: `:`, e: `pony:`},
		{s: []string{`pony`, `meat`}, d: `:`, e: `pony:meat:`},
		{s: []string{`pony`, `meat`, `hoof`}, d: `:`, e: `pony:meat:hoof:`},
		{s: []string{`:`}, d: `:`, e: `::`},
		{s: []string{`::`}, d: `:`, e: `:::`},
		{s: []string{`:`}, d: `::`, e: `:::`},
	} {
		require.Equal(t, test.e, tek.Glue(test.s, test.d), `glue(%+v, "%s")`, test.s, test.d)
	}
}
