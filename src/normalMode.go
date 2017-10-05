package main

import (
  "github.com/nsf/termbox-go"
)

func fillPixel(color termbox.Attribute) {
  app.Canvas.Values[cursor.X][cursor.Y] = color
  app.Canvas.Values[cursor.X + 1][cursor.Y] = color
}

func normalModeKeyMapping(Ch rune, Key termbox.Key) {
  switch Ch {
  case 'x':
    fillPixel(termbox.ColorDefault)
  case 's':
    app.SelectedColor = app.Canvas.Values[cursor.X][cursor.Y]
  }
  switch Key {
  case termbox.KeySpace, termbox.KeyEnter:
    fillPixel(app.SelectedColor)
  }
}
