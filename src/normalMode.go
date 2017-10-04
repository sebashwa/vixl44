package main

import (
  "github.com/nsf/termbox-go"
)

func fillPixel(color termbox.Attribute) {
  app.canvas.values[cursor.X][cursor.Y] = color
  app.canvas.values[cursor.X + 1][cursor.Y] = color
}

func normalModeKeyMapping(Ch rune, Key termbox.Key) {
  switch Ch {
  case 'x':
    fillPixel(termbox.ColorDefault)
  case 's':
    app.selectedColor = app.canvas.values[cursor.X][cursor.Y]
  }
  switch Key {
  case termbox.KeySpace, termbox.KeyEnter:
    fillPixel(app.selectedColor)
  }
}
