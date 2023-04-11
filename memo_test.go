package inc

import (
	"testing"

	"golang.org/x/exp/slices"
)

func Test_matchWithMemo(t *testing.T) {
	for name, tt := range map[string]struct {
		query  []rune
		target string
		memo   memo
	}{
		"empty": {
			[]rune{},
			"abc",
			memo{true, []FoundRune{}},
		},
		"matched": {
			[]rune{'a', 'b', 'c'},
			"aaabbbccc",
			memo{true, []FoundRune{
				{0, 1}, {3, 1}, {6, 1},
			}},
		},
		"not matched": {
			[]rune{'a', 'b', 'c'},
			"aaabbbddd",
			memo{false, []FoundRune{
				{0, 1}, {3, 1},
			}},
		},
		"multi byte char": {
			[]rune{'一', '二', '五'},
			"一二三四五六",
			memo{true, []FoundRune{
				{0, 3}, {3, 3}, {12, 3},
			}},
		},
		"mixed multi byte char": {
			[]rune{'2', '三', '五'},
			"123一二三四五六",
			memo{true, []FoundRune{
				{1, 1}, {9, 3}, {15, 3},
			}},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got := matchWithMemo(tt.query, tt.target)
			if got.matched != tt.memo.matched {
				t.Errorf("matchWithMemo() got.matched = %v, want %v", got.matched, tt.memo.matched)
			}
			if !slices.Equal(got.founds, tt.memo.founds) {
				t.Errorf("matchWithMemo() got.founds = %v, want %v", got.founds, tt.memo.founds)
			}
		})
	}
}
