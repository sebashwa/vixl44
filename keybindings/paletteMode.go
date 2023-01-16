package keybindings

import (
	"github.com/nsf/termbox-go"
	commonActions "github.com/sebashwa/vixl44/actions"
	paintActions "github.com/sebashwa/vixl44/actions/paint"
	"github.com/sebashwa/vixl44/modes"
)

func PaletteMode(Ch rune, Key termbox.Key) {
	switch Ch {
	case 'q':
		commonActions.SetMode(modes.NormalMode)
	}
	switch Key {
	case termbox.KeySpace, termbox.KeyEnter:
		paintActions.SelectColor()
		commonActions.SetMode(modes.NormalMode)
	}
}
