package inc

import (
	"testing"
)

func TestMatch(t *testing.T) {
	for name, tt := range map[string]struct {
		query, body string
		want        bool
	}{
		"empty query": {
			query: "",
			body:  "foo",
			want:  true,
		},
		"single char query": {
			query: "f",
			body:  "foo",
			want:  true,
		},
		"multi char query1": {
			query: "fo",
			body:  "foo",
			want:  true,
		},
		"multi char query2": {
			query: "oa",
			body:  "foobar",
			want:  true,
		},
		"not match": {
			query: "bar",
			body:  "foo",
			want:  false,
		},
	} {
		t.Run(name, func(t *testing.T) {
			if got := Match(tt.query, tt.body); got != tt.want {
				t.Errorf("Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
