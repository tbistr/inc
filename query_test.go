package inc

import (
	"testing"
)

var cands []Candidate = []Candidate{
	{Ptr: 0, Text: "123456"},
	{Ptr: 1, Text: "abcdefg"},
	{Ptr: 2, Text: "123四五六"},
	{Ptr: 3, Text: "あいうえお"},
	{Ptr: 4, Text: "13579"},
	{Ptr: 5, Text: "あかいきうくえけおこ"},
	{Ptr: 6, Text: "ああいいううええおお"},
}

// Assert that each methods don't panic.

func TestEngine_AddQuery(t *testing.T) {
	e := New("123", cands)
	e.AddQuery('四')
	t.Log(e.MatchedIndex())
}

func TestEngine_RmQuery(t *testing.T) {
	e := New("123", cands)
	e.RmQuery()
	t.Log(e.MatchedIndex())
}

func TestEngine_DelQuery(t *testing.T) {
	e := New("123", cands)
	e.DelQuery()
	t.Log(e.MatchedIndex())
}
