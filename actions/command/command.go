package command

import (
  "strings"
  "io/ioutil"
  "encoding/json"

  "../../state"
  "../../types"
)

func Set(command string) {
  state.StatusBar.Command = command
}

func Append(char rune) {
  state.StatusBar.Command = string(append([]rune(state.StatusBar.Command), char))
}

func Truncate() {
  if len(state.StatusBar.Command) > 0 {
    newLength := len(state.StatusBar.Command) - 1
    state.StatusBar.Command = state.StatusBar.Command[:newLength]
  }
}

func Execute() (bool, string, string) {
  command, arguments := getCommandElements()

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

func getCommandElements() (string, []string) {
  commandElements := strings.Split(state.StatusBar.Command, " ")

  if len(commandElements) > 1 {
    return commandElements[0], commandElements[1:]
  } else {
    return commandElements[0], []string{""}
  }
}

func writeStateToFile(pathToFilename string) (string, string) {
  var filename string

  if pathToFilename != "" {
    filename = pathToFilename
  } else if state.Filename != "" {
    filename = state.Filename
  } else {
    return "", "Error: No filename given"
  }

  json, _ := json.Marshal(types.File{state.Canvas.ConvertToFileCanvas()})
  err := ioutil.WriteFile(filename, json, 0644)

  if err != nil {
    return "", "Error: " + err.Error()
  } else {
    state.Filename = filename
    return "Written file to " + filename, ""
  }
}

