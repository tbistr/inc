# inc

inc is an incremental text search library.

<img src="example/demo.gif" width="400">

## install

`go get -u github.com/tbistr/inc`

## Usage

For more information, see [godoc](https://pkg.go.dev/github.com/tbistr/inc).

```golang
// initialize with the initial query and the target strings
// Strs2Cands converts []string to []inc.Cand
target := []string{"foobar", "hogehuga", "bazbar"}
e := inc.New("initial query", inc.Strs2Cands(target))

e.DelQuery() // delete the query
e.AddQuery('o') // add 'o' to the query
fmt.Println(e.MatchedString())
e.RmQuery() // remove the last rune from the query
fmt.Println(e.MatchedString())

// if you want to ignore the case, use inc.IgnoreCase()
e = inc.New("initial query", nil, inc.IgnoreCase())

// you can get some matched object or pointer
e := inc.New("", []inc.Candidate{
    {Ptr: &a, Text: "abc"},
    {Ptr: &b, Text: "def"},
    {Ptr: &c, Text: "ghq"},
})
ms := e.MatchedPtr()
```

## What is incremental search?

For the following text,

```golang
[]string{"foobar", "hogehuga", "bazbar"}
```

query `"ob"` matches `"foobar"` and `"bazbar"`.  
Then add `'a'` to the query, it matches `"foobar"` only.  
Remove `'a'` and add `'r'`, it matches `"foobar"` and `"bazbar"` again.

Like this, incremental search checks

- if the runes in the query are included in the text;
- if the order of the runes is the same as the text.

But doesn't care about substrings between the runes.
