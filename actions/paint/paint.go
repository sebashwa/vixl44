package paint

import(
  "github.com/nsf/termbox-go"

  "github.com/sebashwa/vixl44/state"
  "github.com/sebashwa/vixl44/modes"
)

func AdjustColor(diff int) {
  newIndex := int(state.SelectedColor) + diff

  if newIndex < 1 {
    state.SelectedColor = 256
  } else if newIndex > 256 {
    state.SelectedColor = 1
  } else {
    state.SelectedColor = termbox.Attribute(newIndex)
  }
}

func SelectColor() {
  position := state.Cursor.Position

  if state.CurrentMode == modes.PaletteMode {
    state.SelectedColor = state.Palette[position.X][position.Y]
  } else {
    state.SelectedColor = state.Canvas.Values[position.X][position.Y]
  }
}

func FillPixel(color termbox.Attribute) {
  position := state.Cursor.Position

  state.Canvas.Values[position.X][position.Y] = color
  state.Canvas.Values[position.X + 1][position.Y] = color

  state.History.AddCanvasState(state.Canvas.GetValuesCopy())
}

func rangeLimits(a, b int) (int, int) {
  if a > b {
    return b, a
  }

  return a, b
}

func FillArea(color termbox.Attribute) {
  position := state.Cursor.Position
  fixpoint := state.Cursor.VisualModeFixpoint

  xMin, xMax := rangeLimits(fixpoint.X, position.X)
  yMin, yMax := rangeLimits(fixpoint.Y, position.Y)

  for x := xMin; x <= xMax; x++ {
    for y := yMin; y <= yMax; y++ {
      state.Canvas.Values[x][y] = color
      state.Canvas.Values[x + 1][y] = color
    }
  }

  state.History.AddCanvasState(state.Canvas.GetValuesCopy())
}

func floodFill(x, y int, targetColor termbox.Attribute, replacementColor termbox.Attribute) {
  if targetColor == replacementColor { return }
  if state.Canvas.Values[x][y] != targetColor { return }

  state.Canvas.Values[x][y] = replacementColor

  if x > 0 {
    floodFill(x - 1, y, targetColor, replacementColor)
  }
  if x < state.Canvas.Columns - 1 {
    floodFill(x + 1, y, targetColor, replacementColor)
  }
  if y > 0 {
    floodFill(x, y - 1, targetColor, replacementColor)
  }
  if y < state.Canvas.Rows - 1 {
    floodFill(x, y + 1, targetColor, replacementColor)
  }
}

func FloodFill() {
  x := state.Cursor.Position.X
  y := state.Cursor.Position.Y

  floodFill(x, y, state.Canvas.Values[x][y], state.SelectedColor)
  state.History.AddCanvasState(state.Canvas.GetValuesCopy())
}

