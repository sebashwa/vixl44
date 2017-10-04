package main

import (
  "flag"
  "github.com/nsf/termbox-go"
)

type Modes struct {
  normalMode string
  visualBlockMode string
  colorSelectMode string
  commandMode string
}

type AppState struct {
  canvas Canvas
  palette Palette
  statusBar StatusBar
  commandBar CommandBar
  selectedColor termbox.Attribute
  currentMode string
}

var modes Modes
var app AppState

func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

  if app.currentMode == modes.colorSelectMode {
    app.palette.draw()
  } else {
    app.canvas.draw()
  }

  if app.currentMode == modes.visualBlockMode {
    drawVisualBlockCursor()
  } else {
    drawNormalCursor()
  }

  if app.currentMode == modes.commandMode {
    app.commandBar.draw()
  } else {
    app.statusBar.draw()
  }

	termbox.Flush()
}

func pollEvents() {
loop:
	for {
		switch event := termbox.PollEvent(); event.Type {
		case termbox.EventKey:
      if app.currentMode == modes.commandMode {
        quit := commandModeKeyMapping(event.Ch, event.Key)
        if quit { break loop }
      } else {
        cursorMovementKeyMapping(event.Ch, event.Key)
        statusBarKeyMapping(event.Ch)
        modeKeyMapping(event.Ch, event.Key)

        switch app.currentMode {
        case modes.visualBlockMode:
          visualBlockModeKeyMapping(event.Ch, event.Key)
        case modes.colorSelectMode:
          colorSelectModeKeyMapping(event.Ch, event.Key)
        case modes.normalMode:
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
  palette := createPalette(canvas.rows, canvas.columns)
  statusBar := StatusBar{canvas.rows}
  commandBar := CommandBar{canvas.rows, make([]rune, 0)}
  selectedColor := termbox.Attribute(256)
  currentMode := modes.normalMode

  app = AppState{
    canvas,
    palette,
    statusBar,
    commandBar,
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
