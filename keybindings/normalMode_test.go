package keybindings

import (
	"github.com/nsf/termbox-go"
	"github.com/sebashwa/iwouldlove"
	"testing"

	"github.com/sebashwa/vixl44/state"
	"github.com/sebashwa/vixl44/types"
)

func TestNormalMode(t *testing.T) {
	idLove, it, before := iwouldlove.Init(t)

	before(func() {
		state.Cursor = types.Cursor{}
		state.Canvas = types.Canvas{
			Values: [][]termbox.Attribute{
				[]termbox.Attribute{termbox.ColorWhite},
				[]termbox.Attribute{termbox.ColorWhite},
			},
			Rows:    1,
			Columns: 2,
		}
	})

	it("fills a canvas pixel with the selected color when pressing Space", func() {
		state.SelectedColor = termbox.ColorRed

		NormalMode(0, termbox.KeySpace)

		idLove(state.Canvas.Values[0][0], "to equal", termbox.ColorRed)
		idLove(state.Canvas.Values[1][0], "to equal", termbox.ColorRed)
	})

	it("fills a canvas pixel with the selected color when pressing Enter", func() {
		state.SelectedColor = termbox.ColorRed

		NormalMode(0, termbox.KeyEnter)

		idLove(state.Canvas.Values[0][0], "to equal", termbox.ColorRed)
		idLove(state.Canvas.Values[1][0], "to equal", termbox.ColorRed)
	})

	it("undoes a step in history when pressing 'u'", func() {
		state.SelectedColor = termbox.ColorYellow
		NormalMode(0, termbox.KeyEnter)
		state.SelectedColor = termbox.ColorRed
		NormalMode(0, termbox.KeyEnter)

		NormalMode('u', *new(termbox.Key))

		idLove(state.Canvas.Values[0][0], "to equal", termbox.ColorYellow)
		idLove(state.Canvas.Values[1][0], "to equal", termbox.ColorYellow)
	})

	it("redoes a step in history when pressing Ctrl-r", func() {
		state.SelectedColor = termbox.ColorYellow
		NormalMode(0, termbox.KeyEnter)
		state.SelectedColor = termbox.ColorRed
		NormalMode(0, termbox.KeyEnter)
		NormalMode('u', *new(termbox.Key))

		NormalMode(0, termbox.KeyCtrlR)

		idLove(state.Canvas.Values[0][0], "to equal", termbox.ColorRed)
		idLove(state.Canvas.Values[1][0], "to equal", termbox.ColorRed)
	})

	it("pastes the yank buffer content when pressing 'p'", func() {
		state.YankBuffer = types.YankBuffer{
			Values: [][]termbox.Attribute{
				[]termbox.Attribute{termbox.ColorRed},
				[]termbox.Attribute{termbox.ColorRed},
			},
		}

		NormalMode('p', *new(termbox.Key))

		idLove(state.Canvas.Values[0][0], "to equal", termbox.ColorRed)
		idLove(state.Canvas.Values[1][0], "to equal", termbox.ColorRed)
	})

	it("flood fills the area with the selected color when pressing 'f'", func() {
		state.SelectedColor = termbox.ColorBlue

		NormalMode('f', *new(termbox.Key))

		idLove(state.Canvas.Values[0][0], "to equal", termbox.ColorBlue)
		idLove(state.Canvas.Values[1][0], "to equal", termbox.ColorBlue)
	})

	it("selects the color under the cursor when pressing 's'", func() {
		state.SelectedColor = termbox.ColorBlue

		NormalMode('s', *new(termbox.Key))

		idLove(state.SelectedColor, "to equal", termbox.ColorWhite)
	})

	it("kills the color under the cursor when pressing 'x'", func() {
		NormalMode('x', *new(termbox.Key))

		idLove(state.Canvas.Values[0][0], "to equal", termbox.ColorDefault)
		idLove(state.Canvas.Values[1][0], "to equal", termbox.ColorDefault)
	})
}
