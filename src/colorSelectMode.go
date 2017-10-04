package main

import (
  "github.com/nsf/termbox-go"
)

func selectColor(color termbox.Attribute) {
  app.selectedColor = color
  app.currentMode = modes.normalMode
}

func colorSelectModeKeyMapping(Ch rune, Key termbox.Key) {
  switch Ch {
  case 'q':
    app.currentMode = modes.normalMode
  }
  switch Key {
  case termbox.KeySpace, termbox.KeyEnter:
    selectColor(app.palette.values[cursor.X][cursor.Y])
  }
}
