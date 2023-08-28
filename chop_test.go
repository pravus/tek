package tek_test

import (
	"testing"

	"github.com/pravus/tek"
	"github.com/stretchr/testify/require"
)

func TestChop(t *testing.T) {
	type test struct {
		s string
		d string
		e []string
	}
	for _, test := range []test{
		{s: ``, d: ``, e: []string{}},
		{s: ``, d: `:`, e: []string{}},
		{s: `pork`, d: `:`, e: []string{`pork`}},
		{s: `:pork`, d: `:`, e: []string{`pork`}},
		{s: `::pork`, d: `:`, e: []string{`pork`}},
		{s: `pork:`, d: `:`, e: []string{`pork`}},
		{s: `pork::`, d: `:`, e: []string{`pork`}},
		{s: `:pork:`, d: `:`, e: []string{`pork`}},
		{s: `::pork:`, d: `:`, e: []string{`pork`}},
		{s: `:pork::`, d: `:`, e: []string{`pork`}},
		{s: `::pork::`, d: `:`, e: []string{`pork`}},
		{s: `pork:bone`, d: `:`, e: []string{`pork`, `bone`}},
		{s: `pork::bone`, d: `:`, e: []string{`pork`, `bone`}},
		{s: `:pork:bone`, d: `:`, e: []string{`pork`, `bone`}},
		{s: `::pork:bone`, d: `:`, e: []string{`pork`, `bone`}},
		{s: `:pork::bone`, d: `:`, e: []string{`pork`, `bone`}},
		{s: `::pork::bone`, d: `:`, e: []string{`pork`, `bone`}},
		{s: `pork:bone:`, d: `:`, e: []string{`pork`, `bone`}},
		{s: `pork::bone:`, d: `:`, e: []string{`pork`, `bone`}},
		{s: `pork:bone::`, d: `:`, e: []string{`pork`, `bone`}},
		{s: `pork::bone::`, d: `:`, e: []string{`pork`, `bone`}},
		{s: `:pork:bone:`, d: `:`, e: []string{`pork`, `bone`}},
		{s: `::pork:bone:`, d: `:`, e: []string{`pork`, `bone`}},
		{s: `:pork::bone:`, d: `:`, e: []string{`pork`, `bone`}},
		{s: `:pork:bone::`, d: `:`, e: []string{`pork`, `bone`}},
		{s: `::pork::bone:`, d: `:`, e: []string{`pork`, `bone`}},
		{s: `:pork::bone::`, d: `:`, e: []string{`pork`, `bone`}},
		{s: `::pork:bone::`, d: `:`, e: []string{`pork`, `bone`}},
		{s: `::pork::bone::`, d: `:`, e: []string{`pork`, `bone`}},
		{s: `pork:bone:meat`, d: `:`, e: []string{`pork`, `bone`, `meat`}},
		{s: `:`, d: `pork`, e: []string{`:`}},
		{s: `:`, d: `:`, e: []string{}},
		{s: `::`, d: `:`, e: []string{}},
		{s: `:::`, d: `:`, e: []string{}},
		{s: `::`, d: `:::`, e: []string{`::`}},
	} {
		require.Equal(t, test.e, tek.Chop(test.s, test.d), `chop("%s", "%s")`, test.s, test.d)
	}
}
