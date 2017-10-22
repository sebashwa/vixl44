package actions

import (
  "github.com/sebashwa/vixl44/state"
)

func SetMode(mode string) {
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
