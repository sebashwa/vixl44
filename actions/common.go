package actions

import (
  "github.com/nsf/termbox-go"

  "github.com/sebashwa/vixl44/state"
  "github.com/sebashwa/vixl44/types"
  "github.com/sebashwa/vixl44/actions/paint"
)

func SetMode(mode types.Mode) {
  state.CurrentMode = mode
}

func SetError(message string) {
  state.StatusBar.Error = message
}

func SetHint(message string) {
  state.StatusBar.Hint = message
}

func updateCanvasFromHistory(err error) {
  state.Canvas.Values = state.History.GetCurrentCanvasValuesCopy()
  if err != nil {
    state.StatusBar.Error = err.Error()
  }
}

func Undo() {
  updateCanvasFromHistory(state.History.Undo())
}

func Redo() {
  updateCanvasFromHistory(state.History.Redo())
}

func Copy() {
  xMin, xMax, yMin, yMax := state.Cursor.GetVisualModeArea()
  state.YankBuffer.Set(xMin, xMax, yMin, yMax, state.Canvas.Values)
}

func Cut() {
  Copy()
  paint.FillArea(termbox.ColorDefault)
}

func Paste() {
  position := state.Cursor.Position

loop:
  for x, column := range(state.YankBuffer.Values) {
    if (x + position.X >= state.Canvas.Columns) { break loop }

    for y, color := range(column) {
      if (y + position.Y < state.Canvas.Rows && color != termbox.Attribute(0)) {
        state.Canvas.Values[x + position.X][y + position.Y] = color
      }
    }
  }

  state.History.AddCanvasState(state.Canvas.GetValuesCopy())
}
