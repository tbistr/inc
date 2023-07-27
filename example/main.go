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
	if len(cands) != 0 && cands[len(cands)-1] == "" {
		cands = cands[:len(cands)-1]
	}

	e := inc.New("", inc.Strs2Cands(cands))
	canceled, idx, _ := ui.RunSelector(e)
	if canceled {
		return
	}

	fmt.Println(e.Candidates()[idx])
}
