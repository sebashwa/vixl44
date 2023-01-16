package keybindings

import (
	"github.com/nsf/termbox-go"
	"github.com/sebashwa/vixl44/actions/cursor"
)

func CursorMovement(Ch rune, Key termbox.Key) {
	switch Ch {
	case '0':
		cursor.JumpToBeginningOfLine()
	case '$':
		cursor.JumpToEndOfLine()
	case 'b':
		cursor.Move('X', -10)
	case 'g':
		cursor.JumpToFirstLine()
	case 'G':
		cursor.JumpToLastLine()
	case 'h':
		cursor.Move('X', -2)
	case 'j':
		cursor.Move('Y', 1)
	case 'k':
		cursor.Move('Y', -1)
	case 'l':
		cursor.Move('X', +2)
	case 'w':
		cursor.Move('X', +10)
	}

	switch Key {
	case termbox.KeyCtrlU:
		cursor.Move('Y', -5)
	case termbox.KeyCtrlD:
		cursor.Move('Y', +5)
	}
}
