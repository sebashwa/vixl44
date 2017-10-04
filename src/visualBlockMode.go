package main

import (
  "github.com/nsf/termbox-go"
)

func rangeLimits(a, b int) (int, int) {
  if a > b {
    return b, a
  }

  return a, b
}

func fillArea(color termbox.Attribute) {
  xMin, xMax := rangeLimits(visualModeFixpoint.X, cursor.X)
  yMin, yMax := rangeLimits(visualModeFixpoint.Y, cursor.Y)

  for x := xMin; x <= xMax; x++ {
    for y := yMin; y <= yMax; y++ {
      app.canvas.values[x][y] = color
      app.canvas.values[x + 1][y] = color
    }
  }

  app.currentMode = modes.normalMode
}

func visualBlockModeKeyMapping(Ch rune, Key termbox.Key) {
  switch Ch {
  case 'x':
    fillArea(termbox.ColorDefault)
  }
  switch Key {
  case termbox.KeySpace, termbox.KeyEnter:
    fillArea(app.selectedColor)
  }
}
