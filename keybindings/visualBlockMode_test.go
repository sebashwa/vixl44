package keybindings

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/sebashwa/iwouldlove"
	"testing"

	"github.com/sebashwa/vixl44/modes"
	"github.com/sebashwa/vixl44/state"
	"github.com/sebashwa/vixl44/types"
)

func TestVisualBlockMode(t *testing.T) {
	idLove, it, before := iwouldlove.Init(t)

	before(func() {
		state.CurrentMode = modes.VisualBlockMode
		state.Canvas = types.Canvas{
			Values: [][]termbox.Attribute{
				{termbox.ColorWhite, termbox.ColorWhite},
				{termbox.ColorWhite, termbox.ColorWhite},
				{termbox.ColorWhite, termbox.ColorWhite},
				{termbox.ColorWhite, termbox.ColorWhite},
			},
			Rows:    2,
			Columns: 4,
		}
		state.Cursor = types.Cursor{
			Position: types.Vertex{
				X: 0,
				Y: 0,
			},
			VisualModeFixpoint: types.Vertex{
				X: 2,
				Y: 1,
			},
		}
	})

	it("copies the area and returns to normal mode when pressing 'y'", func() {
		VisualBlockMode('y', *new(termbox.Key))

		idLove(state.YankBuffer.Values, "to equal", state.Canvas.Values)
		idLove(state.CurrentMode, "to equal", modes.NormalMode)
	})

	for _, deleteKey := range []rune{'d', 'x'} {
		it(fmt.Sprintf("cuts the area and returns to normal mode when pressing '%v'", string(deleteKey)), func() {
			expectedYankBuffer := state.Canvas.GetValuesCopy()
			expectedCanvas := [][]termbox.Attribute{
				{termbox.ColorDefault, termbox.ColorDefault},
				{termbox.ColorDefault, termbox.ColorDefault},
				{termbox.ColorDefault, termbox.ColorDefault},
				{termbox.ColorDefault, termbox.ColorDefault},
			}

			VisualBlockMode(deleteKey, *new(termbox.Key))

			idLove(state.YankBuffer.Values, "to equal", expectedYankBuffer)
			idLove(state.Canvas.Values, "to equal", expectedCanvas)
			idLove(state.CurrentMode, "to equal", modes.NormalMode)
		})
	}

	it("fills the area with the selected color and returns to normal mode when pressing Enter", func() {
		state.SelectedColor = termbox.ColorRed
		expectedCanvas := [][]termbox.Attribute{
			{termbox.ColorRed, termbox.ColorRed},
			{termbox.ColorRed, termbox.ColorRed},
			{termbox.ColorRed, termbox.ColorRed},
			{termbox.ColorRed, termbox.ColorRed},
		}

		VisualBlockMode(0, termbox.KeyEnter)

		idLove(state.Canvas.Values, "to equal", expectedCanvas)
		idLove(state.CurrentMode, "to equal", modes.NormalMode)
	})

	it("fills the area with the selected color and returns to normal mode when pressing Space", func() {
		state.SelectedColor = termbox.ColorRed
		expectedCanvas := [][]termbox.Attribute{
			{termbox.ColorRed, termbox.ColorRed},
			{termbox.ColorRed, termbox.ColorRed},
			{termbox.ColorRed, termbox.ColorRed},
			{termbox.ColorRed, termbox.ColorRed},
		}

		VisualBlockMode(0, termbox.KeySpace)

		idLove(state.Canvas.Values, "to equal", expectedCanvas)
		idLove(state.CurrentMode, "to equal", modes.NormalMode)
	})
}
