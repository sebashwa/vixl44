package cursor

import (
  "github.com/sebashwa/vixl44/modes"
  "github.com/sebashwa/vixl44/state"
)

func newPosition(oldPosition int, diff int, limit int) int {
  newPosition := oldPosition + diff

  switch {
  case newPosition >= limit:
    return limit - 1
  case newPosition < 0:
    return 0
  default:
    return newPosition
  }
}

func Move(direction rune, diff int) {
  if direction == 'X' {
    newPosition := newPosition(state.Cursor.Position.X, diff, state.Canvas.Columns - 1)
    state.Cursor.Position.X = newPosition
  } else if direction == 'Y' {
    state.Cursor.Position.Y = newPosition(state.Cursor.Position.Y, diff, state.Canvas.Rows)
  }

  if state.CurrentMode != modes.VisualBlockMode {
    state.Cursor.VisualModeFixpoint.X = state.Cursor.Position.X
    state.Cursor.VisualModeFixpoint.Y = state.Cursor.Position.Y
  }
}

func JumpToEndOfLine() {
  diff := (state.Canvas.Columns - 2) - state.Cursor.Position.X
  Move('X', diff)
}

func JumpToBeginningOfLine() {
  Move('X', -state.Cursor.Position.X)
}

func JumpToFirstLine() {
  Move('Y', -state.Cursor.Position.Y)
}

func JumpToLastLine() {
  diff := (state.Canvas.Rows - 1) - state.Cursor.Position.Y
  Move('Y', diff)
}

