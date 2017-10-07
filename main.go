package main

import (
  "os"
  "flag"
  "io/ioutil"
  "encoding/json"
  "github.com/nsf/termbox-go"
)

type Modes struct {
  NormalMode string
  VisualBlockMode string
  ColorSelectMode string
  CommandMode string
}

type AppState struct {
  Canvas Canvas
  Palette Palette
  StatusBar StatusBar
  Cursor Cursor
  SelectedColor termbox.Attribute
  CurrentMode string
  Filename string
}

var modes Modes
var app AppState

func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

  if app.CurrentMode == modes.ColorSelectMode {
    app.Palette.Draw()
  } else {
    app.Canvas.Draw()
  }

  if app.CurrentMode == modes.VisualBlockMode {
    app.Cursor.DrawBox()
  } else {
    app.Cursor.Draw()
  }

  if app.CurrentMode == modes.CommandMode {
    app.StatusBar.DrawCommand()
  } else {
    app.StatusBar.Draw()
  }

	termbox.Flush()
}

func pollEvents() {
loop:
	for {
		switch event := termbox.PollEvent(); event.Type {
		case termbox.EventKey:
      if app.CurrentMode == modes.CommandMode {
        shouldQuit := commandModeKeyMapping(event.Ch, event.Key)
        if shouldQuit { break loop }
      } else {
        cursorMovementKeyMapping(event.Ch, event.Key)
        statusBarKeyMapping(event.Ch)
        modeKeyMapping(event.Ch, event.Key)

        switch app.CurrentMode {
        case modes.VisualBlockMode:
          visualBlockModeKeyMapping(event.Ch, event.Key)
        case modes.ColorSelectMode:
          colorSelectModeKeyMapping(event.Ch, event.Key)
        case modes.NormalMode:
          normalModeKeyMapping(event.Ch, event.Key)
        }
      }

      draw()
		case termbox.EventResize:
      draw()
		}
	}
}

func parseFlags() (string, int, int) {
  var rows, columns int
  filename := ""

  if len(os.Args[1:]) > 0 {
    firstArg := os.Args[1]

    if rune(firstArg[0]) != '-' {
      filename = firstArg
    }
  }

  for _, value := range([]string{"rows", "r"}) {
    flag.IntVar(&rows, value, 20, "number of rows, 0 means full height, ignored if filename given")
  }

  for _, value := range([]string{"cols", "c"}) {
  flag.IntVar(&columns, value, 20, "number of columns, 0 means full width, ignored if filename given")
  }

  for _, value := range([]string{"f", "filename"}) {
    flag.StringVar(&filename, value, filename, "the name of your file")
  }

  flag.Parse()

  return filename, rows, columns
}

func openOrCreateCanvas(filename string, columns, rows int) Canvas {
  var canvas Canvas

  if _, err := os.Stat(filename); err == nil {
    if data, err := ioutil.ReadFile(filename); err == nil {
      if err := json.Unmarshal(data, &canvas); err != nil {
        panic(err)
      }
    } else {
      panic(err)
    }
  } else {
    canvas = createCanvas(columns, rows)
  }

  return canvas
}

func initializeApp() {
  filename, canvasRows, canvasColumns := parseFlags()

  canvas := openOrCreateCanvas(filename, canvasRows, canvasColumns)
  palette := createPalette(canvas.Rows, canvas.Columns)
  statusBar := StatusBar{canvas.Rows, "", "", ""}
  cursor := Cursor{}
  selectedColor := termbox.Attribute(1)
  currentMode := modes.NormalMode

  app = AppState{
    canvas,
    palette,
    statusBar,
    cursor,
    selectedColor,
    currentMode,
    filename,
  }
}

func setModes() {
  allModes := Modes{
    "NORMAL",
    "VISUAL-BLOCK",
    "COLOR-SELECT",
    "COMMAND",
  }

  modes = allModes
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
  termbox.SetOutputMode(termbox.Output256)

  setModes()
  initializeApp()

	defer termbox.Close()
  termbox.HideCursor()

  draw()
  pollEvents()
}
