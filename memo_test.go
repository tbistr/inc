package inc

import (
	"testing"

	"golang.org/x/exp/slices"
)

func Test_matchWithMemo(t *testing.T) {
	for name, tt := range map[string]struct {
		query  []rune
		target string
		match  bool
		pos    []uint
		len    []uint
	}{
		"empty": {
			query:  []rune{},
			target: "abc",
			match:  true,
			pos:    []uint{},
			len:    []uint{},
		},
		"matched": {
			query:  []rune{'a', 'b', 'c'},
			target: "aaabbbccc",
			match:  true,
			pos:    []uint{0, 3, 6},
			len:    []uint{1, 1, 1},
		},
		"not matched": {
			query:  []rune{'a', 'b', 'c'},
			target: "aaabbbddd",
			match:  false,
			pos:    []uint{0, 3},
			len:    []uint{1, 1},
		},
		"multi byte char": {
			query:  []rune{'一', '二', '五'},
			target: "一二三四五六",
			match:  true,
			pos:    []uint{0, 3, 12},
			len:    []uint{3, 3, 3},
		},
		"mixed multi byte char": {
			query:  []rune{'2', '三', '五'},
			target: "123一二三四五六",
			match:  true,
			pos:    []uint{1, 9, 15},
			len:    []uint{1, 3, 3},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got1, got2, got3 := matchWithMemo(tt.query, tt.target)
			if got1 != tt.match {
				t.Errorf("matchWithMemo() got1 = %v, want %v", got1, tt.match)
			}
			if !slices.Equal(got2, tt.pos) {
				t.Errorf("matchWithMemo() got2 = %v, want %v", got2, tt.pos)
			}
			if !slices.Equal(got3, tt.len) {
				t.Errorf("matchWithMemo() got3 = %v, want %v", got2, tt.len)
			}
		})
	}
}
