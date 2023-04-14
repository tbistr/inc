package ui

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/gdamore/tcell/v2"
	"github.com/tbistr/inc"
)

var (
	defStyle  = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	emphStyle = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorRed)
)

// RunSlector runs the default selector UI.
func RunSelector(e *inc.Engine) {
	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	// s.EnableMouse()
	s.EnablePaste()
	s.Clear()

	fin := func() {
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer fin()

	// Event loop
	for {
		s.Clear()
		printQuery(s, e)
		i := 0
		for _, c := range e.Cands {
			if c.Matched() {
				printCand(s, c, i+1)
				i++
			}
		}
		s.Show()

		// Process event
		switch ev := s.PollEvent().(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				fin()
				if err := syscall.Kill(os.Getpid(), syscall.SIGINT); err != nil {
					fmt.Println(err)
				}
			} else if ev.Key() == tcell.KeyCtrlL {
				s.Sync()
			} else if ev.Key() == tcell.KeyRune {
				e.AddQuery(ev.Rune())
			} else if ev.Key() == tcell.KeyBackspace || ev.Key() == tcell.KeyBackspace2 || ev.Key() == tcell.KeyDelete {
				e.RmQuery()
			} else if ev.Key() == tcell.KeyEnter {
				return
			}

			// case *tcell.EventMouse:
			// 	x, y := ev.Position()

			// 	switch ev.Buttons() {
			// 	case tcell.Button1, tcell.Button2:
			// 		if ox < 0 {
			// 			ox, oy = x, y // record location when click started
			// 		}

			// 	case tcell.ButtonNone:
			// 		if ox >= 0 {
			// 			label := fmt.Sprintf("%d,%d to %d,%d", ox, oy, x, y)
			// 			drawBox(s, ox, oy, x, y, emphStyle, label)
			// 			ox, oy = -1, -1
			// 		}
		}
	}
}
