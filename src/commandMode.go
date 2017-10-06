package main

import (
  "strings"
  "io/ioutil"
  "encoding/json"
  "github.com/nsf/termbox-go"
)

func appendToCommand(char rune) {
  app.StatusBar.Command = string(append([]rune(app.StatusBar.Command), char))
}

func truncateCommand() {
  if len(app.StatusBar.Command) > 0 {
    newLength := len(app.StatusBar.Command) - 1
    app.StatusBar.Command = app.StatusBar.Command[:newLength]
  }
}

func writeStateToFile(pathToFilename string) (string, string) {
  var filename string

  if pathToFilename != "" {
    filename = pathToFilename
  } else if app.Filename != "" {
    filename = app.Filename
  } else {
    return "", "Error: No filename given"
  }

  json, _ := json.Marshal(app.Canvas)
  err := ioutil.WriteFile(filename, json, 0644)

  if err != nil {
    return "", "Error: " + err.Error()
  } else {
    app.Filename = filename
    return "Written file to " + filename, ""
  }
}

func executeCommand() (bool, string, string) {
  commandElements := strings.Split(app.StatusBar.Command, " ")

  command := commandElements[0]
  arguments := []string{""}

  if len(commandElements) > 1 {
    arguments = commandElements[1:]
  }

  switch command {
  case "q", "qu", "qui", "quit":
    return true, "", ""
  case "w", "wr", "wri", "writ", "write":
    hint, errMsg := writeStateToFile(arguments[0])
    return false, hint, errMsg
  case "wq":
    hint, errMsg := writeStateToFile(arguments[0])
    return true, hint, errMsg
  default:
    return false, "", ""
  }
}

func commandModeKeyMapping(Ch rune, Key termbox.Key) bool {
  if Ch != 0 {
    appendToCommand(Ch)
  }

  switch Key {
  case termbox.KeyBackspace, termbox.KeyBackspace2:
    truncateCommand()
  case termbox.KeySpace:
    appendToCommand(' ')
  case termbox.KeyEnter:
    shouldQuit, hint, errMsg := executeCommand()

    if shouldQuit {
      return true
    } else if errMsg != "" {
      app.StatusBar.Error = errMsg
    } else if hint != "" {
      app.StatusBar.Hint = hint
    }

    app.CurrentMode = modes.NormalMode
    app.StatusBar.Command = ""
  default:
    modeKeyMapping('0', Key)
  }

  return false
}
