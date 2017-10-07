package main

import (
  "github.com/nsf/termbox-go"
)

func fillPixel(color termbox.Attribute) {
  position := app.Cursor.Position

  app.Canvas.Values[position.X][position.Y] = color
  app.Canvas.Values[position.X + 1][position.Y] = color
}

func normalModeKeyMapping(Ch rune, Key termbox.Key) {
  switch Ch {
  case 'x':
    fillPixel(termbox.ColorDefault)
  case 's':
    position := app.Cursor.Position

    app.SelectedColor = app.Canvas.Values[position.X][position.Y]
  }
  switch Key {
  case termbox.KeySpace, termbox.KeyEnter:
    fillPixel(app.SelectedColor)
  }
}
