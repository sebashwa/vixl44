package main

import (
  "github.com/nsf/termbox-go"
)

func modeKeyMapping(Ch rune, Key termbox.Key) {
  switch Key {
  case termbox.KeyCtrlV:
    app.currentMode = modes.visualBlockMode
  case termbox.KeyEsc:
    app.currentMode = modes.normalMode
  }

  switch Ch {
  case ':':
    app.currentMode = modes.commandMode
  case 'c':
    app.currentMode = modes.colorSelectMode
  }
}
