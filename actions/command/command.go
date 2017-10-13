package command

import (
  "strings"
  "io/ioutil"
  "encoding/json"
  "errors"

  "github.com/sebashwa/vixl44/state"
  "github.com/sebashwa/vixl44/types"
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

func Execute() (bool, string, error) {
  command, arguments := getCommandElements()

  switch command {
  case "q", "qu", "qui", "quit":
    return true, "", nil
  case "w", "wr", "wri", "writ", "write":
    hint, err := writeStateToFile(arguments[0])
    return false, hint, err
  case "wq":
    hint, err := writeStateToFile(arguments[0])
    return true, hint, err
  case "exp", "expo", "expor", "export":
    hint, err := exportStateToSvg(arguments[0])
    return false, hint, err
  default:
    return false, "", nil
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

func getFilename(filename string) (string, error) {
  if filename != "" {
    return filename, nil
  } else if state.Filename != "" {
    return state.Filename, nil
  } else {
    return "", errors.New("Error: No filename given")
  }
}

func writeStateToFile(pathToFilename string) (string, error) {
  filename, err := getFilename(pathToFilename)

  if err != nil {
    return "", err
  }

  json, _ := json.Marshal(types.File{Canvas: state.Canvas.ConvertToFileCanvas()})
  err = ioutil.WriteFile(filename, json, 0644)

  if err != nil {
    return "", errors.New("Error: " + err.Error())
  } else {
    state.Filename = filename
    return "Written file to " + filename, nil
  }
}

func exportStateToSvg(filename string) (string, error) {
  if filename == "" {
    return "", errors.New("Error: No filename given")
  }

  err := ioutil.WriteFile(filename, []byte(state.Canvas.ConvertToSvg()), 0644)

  if err != nil {
    return "", errors.New("Error: " + err.Error())
  } else {
    return "Exported file to " + filename, nil
  }
}

