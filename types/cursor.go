package types

type Cursor struct {
  Position vertex
  VisualModeFixpoint vertex
}

func rangeLimits(a, b int) (int, int) {
  if a > b {
    return b, a
  }

  return a, b
}

func (cursor Cursor) GetVisualModeArea() (int, int, int, int) {
  position := cursor.Position
  fixpoint := cursor.VisualModeFixpoint

  xMin, xMax := rangeLimits(fixpoint.X, position.X)
  yMin, yMax := rangeLimits(fixpoint.Y, position.Y)

  return xMin, xMax, yMin, yMax
}

