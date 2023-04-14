package inc

import (
	"testing"

	"golang.org/x/exp/slices"
)

var testCands = []string{
	"123456",
	"abcdefg",
	"123四五六",
	"あいうえお",
}

func TestEngine_AddQuery(t *testing.T) {
	for name, tt := range map[string]struct {
		initial string
		add     rune
		matched []int
	}{
		"Simple1":       {"", '3', []int{0, 2}},
		"Simple2":       {"", 'う', []int{3}},
		"Simple3":       {"", 'a', []int{1}},
		"Some initial1": {"12", '3', []int{0, 2}},
		"Some initial2": {"12", '四', []int{2}},
	} {
		t.Run(name, func(t *testing.T) {
			e := New(tt.initial, Strs2Cands(testCands))
			e.AddQuery(tt.add)
			got := e.MatchedIndex()
			if !slices.Equal(got, tt.matched) {
				t.Errorf("AddQuery() got matched index %v, want %v", got, tt.matched)
			}
		})
	}
}

func TestEngine_RmQuery(t *testing.T) {
	for name, tt := range map[string]struct {
		initial string
		matched []int
	}{
		"Simple1":     {"1234", []int{0, 2}},
		"Simple2":     {"あいう", []int{3}},
		"Empty":       {"", []int{0, 1, 2, 3}},
		"Non ASCII 1": {"四五", []int{2}},
		"Non ASCII 2": {"あ胃", []int{3}},
	} {
		t.Run(name, func(t *testing.T) {
			e := New(tt.initial, Strs2Cands(testCands))
			e.RmQuery()
			got := e.MatchedIndex()
			if !slices.Equal(got, tt.matched) {
				t.Errorf("RmQuery() got matched index %v, want %v", got, tt.matched)
			}
		})
	}
}

func TestEngine_DelQuery(t *testing.T) {
	for name, tt := range map[string]struct {
		initial string
	}{
		"Simple1": {"1234"},
		"Simple2": {"あいう"},
		"Empty":   {""},
		"Long":    {"123456789123456789123456789123456789123456789123456789123456789123456789123456789123456789123456789123456789"},
	} {
		t.Run(name, func(t *testing.T) {
			e := New(tt.initial, Strs2Cands(testCands))
			e.DelQuery()
			got := e.MatchedIndex()
			if !slices.Equal(got, []int{0, 1, 2, 3}) {
				t.Errorf("DelQuery() got matched index %v, want %v", got, []int{0, 1, 2, 3})
			}
		})
	}
}

func TestEngine_SwapQuery(t *testing.T) {
	for name, tt := range map[string]struct {
		initial, swap string
	}{
		"Efficient swap":   {"123", "12*456"},
		"All rune swapped": {"123", "654321"},
		"Empty query":      {"123", ""},
	} {
		t.Run(name, func(t *testing.T) {
			e := New(tt.initial, Strs2Cands(testCands))
			e.SwapQuery(tt.swap)
			got := e.Query()
			if got != tt.swap {
				t.Errorf("SwapQuery() got query %s, want %s", got, tt.swap)
			}
		})
	}
}
