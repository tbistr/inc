package main

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/tbistr/inc"
)

var ss = []string{
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	"Feugiat nisl pretium fusce id velit ut tortor pretium",
	"Magna etiam tempor orci eu lobortis",
	"Tempor nec feugiat nisl pretium",
	"Sodales neque sodales ut etiam sit amet",
	"Est ante in nibh mauris",
	"In aliquam sem fringilla ut morbi tincidunt augue",
	"Sed pulvinar proin gravida hendrerit lectus",
	"Sapien nec sagittis aliquam malesuada bibendum arcu vitae elementum",
	"Id interdum velit laoreet id donec ultrices tincidunt",
	"Sagittis eu volutpat odio facilisis mauris sit amet massa",
	"Ac tortor dignissim convallis aenean et tortor at risus",
	"Pharetra vel turpis nunc eget lorem dolor",
	"Purus ut faucibus pulvinar elementum integer",
	"Massa ultricies mi quis hendrerit dolor magna",
	"Sed pulvinar proin gravida hendrerit lectus a",
	"Sit amet venenatis urna cursus eget nunc scelerisque viverra",
	"Vitae congue eu consequat ac felis donec et",
	"Vel fringilla est ullamcorper eget nulla facilisi etiam dignissim diam",
	"Vel risus commodo viverra maecenas",
	"Cursus eget nunc scelerisque viverra mauris in aliquam",
	"Magna eget est lorem ipsum dolor sit amet",
	"Quam quisque id diam vel quam elementum",
	"Quisque id diam vel quam elementum pulvinar etiam non",
	"Condimentum vitae sapien pellentesque habitant morbi tristique senectus",
	"Vivamus arcu felis bibendum ut tristique et egestas",
	"Cursus metus aliquam eleifend mi in nulla posuere sollicitudin aliquam",
	"Neque aliquam vestibulum morbi blandit cursus risus",
}

var ENGINE *inc.Engine = inc.New("", inc.Strs2Cands(ss))

var DEBUG = tview.NewTextView()

func dumpCands() string {
	idx := ENGINE.MatchedIndex()
	cur := 0
	b := strings.Builder{}
	color := ""
	for i, c := range ENGINE.Cands {
		DEBUG.SetText(fmt.Sprintf("idx: %v, cur: %v", idx, cur))
		if cur < len(idx) && i == idx[cur] {
			b.WriteString("+ ")
			color = "[white]"
			b.WriteString(color)
			cur++
		} else {
			b.WriteString("- ")
			color = "[grey]"
			b.WriteString(color)
		}
		founds := c.FoundRunes()
		last := uint(0)
		for _, f := range founds {
			b.WriteString(c.Text[last:f.Pos])
			b.WriteString("[red]")
			b.WriteString(c.Text[f.Pos : f.Pos+f.Len])
			b.WriteString(color)
			last = f.Pos + f.Len
		}
		b.WriteString(c.Text[last:])
		b.WriteString("\n")
	}
	return b.String()
}

func main() {
	DEBUG.
		SetTitle("debug").
		SetBorder(true)

	app := tview.NewApplication()

	result := tview.NewTextView()
	result.SetTitle("Result").SetBorder(true)
	result.SetDynamicColors(true)
	result.SetText(dumpCands())

	query := tview.NewInputField()
	query.
		SetLabel("query:").
		SetBorder(true)

	query.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyRune:
			ENGINE.AddQuery(event.Rune())
		case tcell.KeyBackspace, tcell.KeyBackspace2:
			ENGINE.RmQuery()
		}
		result.SetText("")
		result.SetText(dumpCands())
		return event
	})

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(query, 3, 1, true).
		// AddItem(DEBUG, 3, 1, false).
		AddItem(result, 0, 1, false)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
