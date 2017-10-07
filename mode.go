package main

import (
  "github.com/nsf/termbox-go"
)

func modeKeyMapping(Ch rune, Key termbox.Key) {
  switch Key {
  case termbox.KeyCtrlV:
    app.CurrentMode = modes.VisualBlockMode
  case termbox.KeyEsc:
    app.CurrentMode = modes.NormalMode
  }

  switch Ch {
  case ':':
    app.CurrentMode = modes.CommandMode
  case 'c':
    app.CurrentMode = modes.ColorSelectMode
  }
}
