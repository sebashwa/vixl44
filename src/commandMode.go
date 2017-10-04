package main

import (
  "github.com/nsf/termbox-go"
)

func appendToCommand(char rune) bool {
  app.commandBar.value = append(app.commandBar.value, char)
  return false
}

func truncateCommand() bool {
  if len(app.commandBar.value) > 0 {
    newLength := len(app.commandBar.value) - 1
    app.commandBar.value = app.commandBar.value[:newLength]
  }
  return false
}

func executeCommand() bool {
  switch string(app.commandBar.value) {
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
