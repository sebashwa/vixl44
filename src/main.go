package main

import (
  "flag"
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
  CommandBar CommandBar
  Cursor Cursor
  SelectedColor termbox.Attribute
  CurrentMode string
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
    app.CommandBar.Draw()
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
        quit := commandModeKeyMapping(event.Ch, event.Key)
        if quit { break loop }
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

func parseFlags() (int, int) {
  var rows, columns int

  flag.IntVar(&rows, "rows", 20, "number of rows on your canvas, 0 means full height")
  flag.IntVar(&columns, "cols", 20, "number of columns on your canvas, 0 means full width")

  flag.Parse()

  return rows, columns
}

func initializeApp() {
  rows, columns := parseFlags()

  canvas := createCanvas(rows, columns)
  palette := createPalette(canvas.Rows, canvas.Columns)
  statusBar := StatusBar{canvas.Rows}
  commandBar := CommandBar{canvas.Rows, make([]rune, 0)}
  cursor := Cursor{}
  selectedColor := termbox.Attribute(256)
  currentMode := modes.NormalMode

  app = AppState{
    canvas,
    palette,
    statusBar,
    commandBar,
    cursor,
    selectedColor,
    currentMode,
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
