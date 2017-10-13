package main

import (
  "os"
  "flag"
  "io/ioutil"
  "encoding/json"
  "github.com/nsf/termbox-go"

  "github.com/sebashwa/vixl44/drawing"
  "github.com/sebashwa/vixl44/keybindings"
  "github.com/sebashwa/vixl44/state"
  "github.com/sebashwa/vixl44/modes"
  "github.com/sebashwa/vixl44/types"
  "github.com/sebashwa/vixl44/factory"
)

func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

  if state.CurrentMode == modes.PaletteMode {
    drawing.DrawPalette()
  } else {
    drawing.DrawCanvas()
  }

  if state.CurrentMode == modes.VisualBlockMode {
    drawing.DrawVisualBlockCursor()
  } else {
    drawing.DrawCursor()
  }

  if state.CurrentMode == modes.CommandMode {
    drawing.DrawCommand()
  } else {
    drawing.DrawStatusBar()
  }

	termbox.Flush()
}

func pollEvents() {
loop:
	for {
		switch event := termbox.PollEvent(); event.Type {
		case termbox.EventKey:
      if state.CurrentMode == modes.CommandMode {
        shouldQuit := keybindings.CommandMode(event.Ch, event.Key)
        if shouldQuit { break loop }
      } else {
        keybindings.CursorMovement(event.Ch, event.Key)
        keybindings.Common(event.Ch)
        keybindings.ModeSelection(event.Ch, event.Key)

        switch state.CurrentMode {
        case modes.VisualBlockMode:
          keybindings.VisualBlockMode(event.Ch, event.Key)
        case modes.PaletteMode:
          keybindings.PaletteMode(event.Ch, event.Key)
        case modes.NormalMode:
          keybindings.NormalMode(event.Ch, event.Key)
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

func openOrCreateCanvas(filename string, columns, rows int) types.Canvas {
  if _, err := os.Stat(filename); err == nil {
    if data, err := ioutil.ReadFile(filename); err == nil {
      var file types.File
      if err := json.Unmarshal(data, &file); err != nil {
        panic(err)
      }

      return factory.CreateCanvasFromFileCanvas(file.Canvas)
    } else {
      panic(err)
    }
  } else {
    return factory.CreateCanvas(columns, rows)
  }
}

func initializeApp() {
  filename, canvasRows, canvasColumns := parseFlags()

  modes.NormalMode = "NORMAL"
  modes.VisualBlockMode = "VISUAL-BLOCK"
  modes.PaletteMode = "PALETTE"
  modes.CommandMode = "COMMAND"

  state.Canvas = openOrCreateCanvas(filename, canvasRows, canvasColumns)
  state.Palette = factory.CreatePalette(state.Canvas.Rows, state.Canvas.Columns)
  state.StatusBar = types.StatusBar{state.Canvas.Rows, "", "", ""}
  state.Cursor = types.Cursor{}
  state.SelectedColor = termbox.Attribute(4)
  state.CurrentMode = modes.NormalMode
  state.Filename = filename
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
  termbox.SetOutputMode(termbox.Output256)

  initializeApp()

	defer termbox.Close()
  termbox.HideCursor()

  draw()
  pollEvents()
}
