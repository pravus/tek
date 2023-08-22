package tek_test

import (
	"encoding/json"
	"testing"

	"github.com/pravus/tek"
	"github.com/stretchr/testify/require"
)

type U struct {
	A []byte  `json:"z"`
	B bool    `json:"b"`
	F float64 `json:"f"`
	I int     `json:"i"`
	R rune    `json:"r"`
	S string  `json:"s"`
}

func TestU(t *testing.T) {
	for _, c := range []struct {
		s string
		v *U
		e string
	}{
		{s: `{}`, v: &U{}, e: ``},
		{s: `{"b": true}`, v: &U{B: true}, e: ``},
		{s: `{"b": false}`, v: &U{B: false}, e: ``},
		{s: `{"f": 117e-2}`, v: &U{F: 117e-2}, e: ``},
		{s: `{"i": 117}`, v: &U{I: 117}, e: ``},
		{s: `{"r": 117}`, v: &U{R: 'u'}, e: ``},
		{s: `{"s": "u"}`, v: &U{S: `u`}, e: ``},
		{s: `{"z": "dQ=="}`, v: &U{A: []byte(`u`)}, e: ``},
		{s: `}{`, v: nil, e: "invalid character '}' looking for beginning of value"},
	} {
		if v, e := tek.U([]byte(c.s), U{}, json.Unmarshal); e != nil {
			if c.e == `` {
				t.Errorf(`expected no error but received %s`, e)
			}
			require.ErrorContains(t, e, c.e)
		} else {
			if c.e != `` {
				t.Errorf(`expected error but received none`)
			}
			require.Equal(t, v, c.v)
		}
	}
}
