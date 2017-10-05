package main

import (
  "github.com/nsf/termbox-go"
)

func selectColor(color termbox.Attribute) {
  app.SelectedColor = color
  app.CurrentMode = modes.NormalMode
}

func colorSelectModeKeyMapping(Ch rune, Key termbox.Key) {
  switch Ch {
  case 'q':
    app.CurrentMode = modes.NormalMode
  }
  switch Key {
  case termbox.KeySpace, termbox.KeyEnter:
    selectColor(app.Palette.Values[cursor.X][cursor.Y])
  }
}
