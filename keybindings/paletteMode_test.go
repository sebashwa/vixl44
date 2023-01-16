package keybindings

import (
	"github.com/nsf/termbox-go"
	"github.com/sebashwa/iwouldlove"
	"testing"

	"github.com/sebashwa/vixl44/modes"
	"github.com/sebashwa/vixl44/state"
	"github.com/sebashwa/vixl44/types"
)

func TestPaletteMode(t *testing.T) {
	idLove, it, before := iwouldlove.Init(t)

	before(func() {
		state.CurrentMode = modes.PaletteMode
		state.Cursor = types.Cursor{}
		state.Palette = types.Palette{
			[]termbox.Attribute{termbox.Attribute(3)},
		}
	})

	it("returns to normal mode when pressing 'q'", func() {
		PaletteMode('q', *new(termbox.Key))

		idLove(state.CurrentMode, "to equal", modes.NormalMode)
	})

	it("selects the color under the cursor and returns to normal mode when pressing Space", func() {
		state.SelectedColor = termbox.Attribute(0)

		PaletteMode(0, termbox.KeySpace)

		idLove(state.SelectedColor, "to equal", termbox.Attribute(3))
		idLove(state.CurrentMode, "to equal", modes.NormalMode)
	})

	it("selects the color under the cursor and returns to normal mode when pressing Enter", func() {
		state.SelectedColor = termbox.Attribute(0)

		PaletteMode(0, termbox.KeyEnter)

		idLove(state.SelectedColor, "to equal", termbox.Attribute(3))
		idLove(state.CurrentMode, "to equal", modes.NormalMode)
	})
}
