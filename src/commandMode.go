package main

import (
  "github.com/nsf/termbox-go"
)

func appendToCommand(char rune) bool {
  app.CommandBar.Value = append(app.CommandBar.Value, char)
  return false
}

func truncateCommand() bool {
  if len(app.CommandBar.Value) > 0 {
    newLength := len(app.CommandBar.Value) - 1
    app.CommandBar.Value = app.CommandBar.Value[:newLength]
  }
  return false
}

func executeCommand() bool {
  switch string(app.CommandBar.Value) {
  case "q":
    return true
  default:
    return false
  }
}

func commandModeKeyMapping(Ch rune, Key termbox.Key) bool {
  if Ch != 0 {
    return appendToCommand(Ch)
  }

  switch Key {
  case termbox.KeyBackspace, termbox.KeyBackspace2:
    return truncateCommand()
  case termbox.KeyEnter:
    return executeCommand()
  default:
    modeKeyMapping('0', Key)
    return false
  }
}
