package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/tbistr/inc"
	"github.com/tbistr/inc/ui"
	"golang.org/x/term"
)

const say = `------------------------------------------------
| You need to pipe some input to this program! |
------------------------------------------------
      \   ^__^
       \  (oo)\_______
          (__)\       )\/\
              ||----w |
              ||     ||
`

func main() {
	// Check if stdin is a terminal or piped.
	if term.IsTerminal(int(os.Stdin.Fd())) {
		fmt.Print(say)
		return
	}
	stdin, _ := io.ReadAll(os.Stdin)
	cands := strings.Split(string(stdin), "\n")

	e := inc.New("", inc.Strs2Cands(cands))
	ui.RunSelector(e)

	for _, c := range e.Matched() {
		fmt.Println(c)
	}
}
