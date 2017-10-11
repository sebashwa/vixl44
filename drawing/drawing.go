package drawing

import (
  "github.com/nsf/termbox-go"

  "../state"
  "../modes"
)

func DrawCanvas() {
  for x, column := range state.Canvas.Values {
    for y := range column {
      color := state.Canvas.Values[x][y]
      termbox.SetCell(x, y, ' ', color, color)
    }
  }
}

func DrawPalette() {
  for x, column := range state.Palette.Values {
    for y := range column {
      color := termbox.Attribute(state.Palette.Values[x][y])
      termbox.SetCell(x, y, ' ', color, color)
    }
  }
}

func DrawStatusBar() {
  if state.StatusBar.Error != "" {
    displayMessage(state.StatusBar.Error, state.StatusBar.Position, termbox.ColorRed)
  } else if state.StatusBar.Hint != "" {
    displayMessage(state.StatusBar.Hint, state.StatusBar.Position, termbox.ColorGreen)
  } else {
    displayStatus(state.StatusBar.Position)
  }

  state.StatusBar.Hint = ""
  state.StatusBar.Error = ""
}

func displayStatus(position int) {
  for i, character := range state.CurrentMode {
    foregroundColor := state.SelectedColor
    if state.SelectedColor == termbox.Attribute(1) {
      foregroundColor = termbox.ColorWhite
    }

    termbox.SetCell(i, position, character, foregroundColor, termbox.ColorDefault)
  }

  for x := len(state.CurrentMode) + 1; x < state.Canvas.Columns; x++ {
    termbox.SetCell(x, position, ' ', state.SelectedColor, state.SelectedColor)
  }
}

func displayMessage(message string, position int, color termbox.Attribute) {
  for i, char := range(message) {
    termbox.SetCell(i, position, char, color, termbox.ColorDefault)
  }
}

func DrawCommand() {
  termbox.SetCell(0, state.StatusBar.Position, ':', termbox.ColorWhite, termbox.ColorDefault)

  for i, char := range(state.StatusBar.Command) {
    termbox.SetCell(i + 1, state.StatusBar.Position, char, termbox.ColorWhite, termbox.ColorDefault)
  }
}

func DrawCursor() {
  drawCursor(state.Cursor.Position.X, state.Cursor.Position.Y)
}

func DrawVisualBlockCursor() {
  position := state.Cursor.Position
  fixpoint := state.Cursor.VisualModeFixpoint

  drawCursor(position.X, position.Y)
  drawCursor(fixpoint.X, fixpoint.Y)
  drawCursor(fixpoint.X, position.Y)
  drawCursor(position.X, fixpoint.Y)
}

func drawCursor(x, y int) {
  cursorColor := termbox.ColorWhite
  var backgroundColor termbox.Attribute

  if state.CurrentMode == modes.PaletteMode {
    backgroundColor = state.Palette.Values[x][y]
  } else {
    backgroundColor = state.Canvas.Values[x][y]
  }

  if _, exists := state.Palette.LightColors[backgroundColor]; exists {
    cursorColor = termbox.ColorBlack
  }

  termbox.SetCell(x, y, '[', cursorColor, backgroundColor)
  termbox.SetCell(x + 1, y, ']', cursorColor, backgroundColor)
}

