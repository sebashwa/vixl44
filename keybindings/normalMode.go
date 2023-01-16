package keybindings

import (
	"github.com/nsf/termbox-go"
	commonActions "github.com/sebashwa/vixl44/actions"
	paintActions "github.com/sebashwa/vixl44/actions/paint"
)

func NormalMode(Ch rune, Key termbox.Key) {
	switch Ch {
	case 'x':
		paintActions.KillPixel()
	case 's':
		paintActions.SelectColor()
	case 'f':
		paintActions.FloodFill()
	case 'u':
		commonActions.Undo()
	case 'p':
		commonActions.Paste()
	}

	switch Key {
	case termbox.KeyCtrlR:
		commonActions.Redo()
	case termbox.KeySpace, termbox.KeyEnter:
		paintActions.FillPixel()
	}
}
