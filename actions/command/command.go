package command

import (
  "encoding/json"
  "errors"
  "io/ioutil"
  "strings"
  "strconv"

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
    hint, err := exportStateToImage(arguments[0])
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
    return "", errors.New("No filename given")
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
    return "", errors.New(err.Error())
  } else {
    state.Filename = filename
    return "Written file to " + filename, nil
  }
}

func exportStateToImage(filename string) (string, error) {
  if filename == "" {
    return "", errors.New("No filename given")
  }

  var err error
  var buf []byte

  filenameElements := strings.Split(filename, ".")
  extension := filenameElements[len(filenameElements) -1]

  switch extension {
  case "svg":
    buf, err = state.Canvas.ConvertToSvg()
  case "png":
    scaleFactor := 1
    filenameAndScaleFactor := strings.Split(filenameElements[0], "x")
    scaleFactorString := filenameAndScaleFactor[len(filenameAndScaleFactor) - 1]

    if parsedFactor, err := strconv.Atoi(scaleFactorString); err == nil {
      scaleFactor = parsedFactor
    }

    buf, err = state.Canvas.ConvertToPng(scaleFactor)
  default:
    err = errors.New("Add .svg or .png as extension")
  }

  if err != nil {
    return "", errors.New(err.Error())
  }

  err = ioutil.WriteFile(filename, buf, 0644)

  if err != nil {
    return "", errors.New(err.Error())
  }

  return "Exported file to " + filename, nil
}
