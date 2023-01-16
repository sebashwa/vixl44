package keybindings

import (
	"github.com/nsf/termbox-go"
	"github.com/sebashwa/iwouldlove"
	"testing"

	"github.com/sebashwa/vixl44/state"
	"github.com/sebashwa/vixl44/types"
)

func TestCursorMovement(t *testing.T) {
	idLove, it, before := iwouldlove.Init(t)

	before(func() {
		state.Canvas = types.Canvas{
			Rows:    40,
			Columns: 40,
		}

		state.Cursor = types.Cursor{
			Position: types.Vertex{
				X: 15,
				Y: 15,
			},
		}
	})

	it("moves the cursor two columns to the left when pressing 'h'", func() {
		CursorMovement('h', *new(termbox.Key))
		idLove(state.Cursor.Position.X, "to equal", 13)
	})

	it("moves the cursor two columns to the right when pressing 'l'", func() {
		CursorMovement('l', *new(termbox.Key))
		idLove(state.Cursor.Position.X, "to equal", 17)
	})

	it("moves the cursor one row to the bottom when pressing 'j'", func() {
		CursorMovement('j', *new(termbox.Key))
		idLove(state.Cursor.Position.Y, "to equal", 16)
	})

	it("moves the cursor one row to the top when pressing 'k'", func() {
		CursorMovement('k', *new(termbox.Key))
		idLove(state.Cursor.Position.Y, "to equal", 14)
	})

	it("jumps to the beginning of line when pressing '0'", func() {
		CursorMovement('0', *new(termbox.Key))
		idLove(state.Cursor.Position.X, "to equal", 0)
	})

	it("jumps to the end of line when pressing '$'", func() {
		CursorMovement('$', *new(termbox.Key))
		idLove(state.Cursor.Position.X, "to equal", 38)
	})

	it("jumps to the first row when pressing 'g'", func() {
		CursorMovement('g', *new(termbox.Key))
		idLove(state.Cursor.Position.Y, "to equal", 0)
	})

	it("jumps to the last row when pressing 'G'", func() {
		CursorMovement('G', *new(termbox.Key))
		idLove(state.Cursor.Position.Y, "to equal", 39)
	})

	it("jumps 10 columns to the right when pressing 'w'", func() {
		CursorMovement('w', *new(termbox.Key))
		idLove(state.Cursor.Position.X, "to equal", 25)
	})

	it("jumps 10 columns to the left when pressing 'b'", func() {
		CursorMovement('b', *new(termbox.Key))
		idLove(state.Cursor.Position.X, "to equal", 5)
	})

	it("jumps 5 rows to the bottom when pressing 'Ctrl-d'", func() {
		CursorMovement(0, termbox.KeyCtrlD)
		idLove(state.Cursor.Position.Y, "to equal", 20)
	})

	it("jumps 5 rows to the top when pressing 'Ctrl-u'", func() {
		CursorMovement(0, termbox.KeyCtrlU)
		idLove(state.Cursor.Position.Y, "to equal", 10)
	})
}
