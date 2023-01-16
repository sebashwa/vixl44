package main

import (
	"encoding/json"
	"flag"
	"github.com/nsf/termbox-go"
	"log"
	"os"

	"github.com/sebashwa/vixl44/drawing"
	"github.com/sebashwa/vixl44/factory"
	"github.com/sebashwa/vixl44/keybindings"
	"github.com/sebashwa/vixl44/modes"
	"github.com/sebashwa/vixl44/state"
	"github.com/sebashwa/vixl44/types"
)

func draw() {
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		log.Println(err)
		return
	}

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

	err = termbox.Flush()
	if err != nil {
		log.Println(err)
		return
	}
}

func pollEvents() {
loop:
	for {
		switch event := termbox.PollEvent(); event.Type {
		case termbox.EventKey:
			if state.CurrentMode == modes.CommandMode {
				shouldQuit := keybindings.CommandMode(event.Ch, event.Key)
				if shouldQuit {
					break loop
				}
			} else {
				switch state.CurrentMode {
				case modes.VisualBlockMode:
					keybindings.VisualBlockMode(event.Ch, event.Key)
				case modes.PaletteMode:
					keybindings.PaletteMode(event.Ch, event.Key)
				case modes.NormalMode:
					keybindings.NormalMode(event.Ch, event.Key)
				}

				keybindings.CursorMovement(event.Ch, event.Key)
				keybindings.Common(event.Ch)
				keybindings.ModeSelection(event.Ch, event.Key)
			}

			draw()
		case termbox.EventResize:
			draw()
		}
	}
}

func parseArguments() (string, int, int) {
	var rows, columns int

	for _, value := range []string{"rows", "r"} {
		flag.IntVar(&rows, value, 20, "number of rows, default is 20, 0 means full height, ignored if name of existing file given")
	}

	for _, value := range []string{"cols", "c"} {
		flag.IntVar(&columns, value, 20, "number of columns, default is 20, 0 means full width, ignored if name of existing file given")
	}

	flag.Parse()

	return flag.Arg(0), rows, columns
}

func openOrCreateCanvas(filename string, columns, rows int) types.Canvas {
	if _, err := os.Stat(filename); err == nil {
		if data, err := os.ReadFile(filename); err == nil {
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
	filename, canvasRows, canvasColumns := parseArguments()

	state.Canvas = openOrCreateCanvas(filename, canvasColumns, canvasRows)
	state.Palette = factory.CreatePalette(state.Canvas.Columns, state.Canvas.Rows)
	state.StatusBar = types.StatusBar{
		Position: state.Canvas.Rows,
		Hint:     "",
		Error:    "",
		Command:  "",
	}
	state.Cursor = types.Cursor{}
	state.SelectedColor = termbox.Attribute(4)
	state.CurrentMode = modes.NormalMode
	state.Filename = filename

	state.History = types.History{}
	state.History.AddCanvasState(state.Canvas.GetValuesCopy())
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
