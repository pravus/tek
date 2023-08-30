package tek_test

import (
	"encoding/json"
	"testing"

	"github.com/pravus/tek"
	"github.com/stretchr/testify/require"
)

func TestP(t *testing.T) {
	type test struct {
		P any    `json:"p,omitempty"`
		E string `json:"e"`
	}
	for _, test := range []test{
		{P: nil, E: `null`},
		{P: true, E: `true`},
		{P: false, E: `false`},
		{P: "p", E: `"p"`},
		{P: 1e-1, E: `0.1`},
		{P: []any{nil}, E: "[\n  null\n]"},
		{P: test{}, E: "{\n  \"e\": \"\"\n}"},
		{P: test{E: "E"}, E: "{\n  \"e\": \"E\"\n}"},
		{P: test{P: test{}}, E: "{\n  \"p\": {\n    \"e\": \"\"\n  },\n  \"e\": \"\"\n}"},
		{P: test{P: test{E: "E"}}, E: "{\n  \"p\": {\n    \"e\": \"E\"\n  },\n  \"e\": \"\"\n}"},
	} {
		require.Equal(t, test.E, tek.P(test.P, json.MarshalIndent), `p(%+v)`, test.P)
	}
}
