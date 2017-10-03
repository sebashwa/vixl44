package main

import (
  "flag"
  "github.com/nsf/termbox-go"
)

type Vertex struct {
  X int
  Y int
}

const visualBlockMode = "VISUAL-BLOCK"
const normalMode = "NORMAL"
const colorSelectMode = "COLOR-SELECT"
const commandMode = "COMMAND"

var canvasRows int
var canvasColumns int
const canvasRowsDefault = 20
const canvasColumnsDefault = 20

var canvas [][]termbox.Attribute
var palette [][]termbox.Attribute

var mode = normalMode
var cursor = Vertex{0,0}
var visualModeFixpoint = Vertex{0,0}
var selectedColor = termbox.Attribute(256)
var command []rune

/* DRAWING */

func drawCursor(x, y int) {
  cursorColor := termbox.ColorWhite
  var backgroundColor termbox.Attribute

  if mode == colorSelectMode {
    backgroundColor = termbox.Attribute(palette[x][y])
  } else {
    backgroundColor = canvas[x][y]
  }

  if backgroundColor == termbox.ColorWhite {
    cursorColor = termbox.ColorBlack
  }

  termbox.SetCell(x, y, '[', cursorColor, backgroundColor)
  termbox.SetCell(x + 1, y, ']', cursorColor, backgroundColor)
}

func drawCanvas() {
  for x, column := range canvas {
    for y := range column {
      color := canvas[x][y]
      termbox.SetCell(x, y, ' ', color, color)
    }
  }
}

func drawPalette() {
  for x, column := range palette {
    for y := range column {
      color := termbox.Attribute(palette[x][y])
      termbox.SetCell(x, y, ' ', color, color)
    }
  }
}

func drawStatusBar() {
  if mode == commandMode {
    termbox.SetCell(0, canvasRows, ':', termbox.ColorWhite, termbox.ColorDefault)
    for i, char := range(command) {
      termbox.SetCell(i + 1, canvasRows, char, termbox.ColorWhite, termbox.ColorDefault)
    }
  } else {
    selectedColor := termbox.Attribute(selectedColor)

    for i, character := range mode {
      termbox.SetCell(i, canvasRows, character, selectedColor, termbox.ColorDefault)
    }

    for x := len(mode) + 1; x < canvasColumns; x++ {
      termbox.SetCell(x, canvasRows, ' ', selectedColor, selectedColor)
    }
  }
}

func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

  if mode == colorSelectMode {
    drawPalette()
  } else {
    drawCanvas()
  }
  drawCursor(cursor.X, cursor.Y)
  if mode == visualBlockMode {
    drawCursor(visualModeFixpoint.X, visualModeFixpoint.Y)
    drawCursor(visualModeFixpoint.X, cursor.Y)
    drawCursor(cursor.X, visualModeFixpoint.Y)
  }
  drawStatusBar()

	termbox.Flush()
}

/* ACTIONS */

func newPosition(oldPosition int, diff int, limit int) int {
  newPosition := oldPosition + diff

  switch {
  case newPosition >= limit:
    return limit - 1
  case newPosition < 0:
    return 0
  default:
    return newPosition
  }
}

func moveCursor(direction rune, diff int) {
  if direction == 'X' {
    cursor.X = newPosition(cursor.X, diff, canvasColumns - 1)
  } else if direction == 'Y' {
    cursor.Y = newPosition(cursor.Y, diff, canvasRows)
  }

  draw()
}

func jumpToEndOfLine() {
  cursor.X = canvasColumns - 2
  draw()
}

func jumpToBeginningOfLine() {
  cursor.X = 0
  draw()
}

func jumpToFirstLine() {
  cursor.Y = 0
  draw()
}

func jumpToLastLine() {
  cursor.Y = canvasRows - 1
  draw()
}

func adjustColor(diff int) {
  newIndex := int(selectedColor) + diff

  if newIndex < 1 {
    selectedColor = 256
  } else if newIndex > 256 {
    selectedColor = 1
  } else {
    selectedColor = termbox.Attribute(newIndex)
  }

  draw()
}

func selectColor(color termbox.Attribute) {
  selectedColor = color
  mode = normalMode
  draw()
}

func rangeLimits(a, b int) (int, int) {
  if a > b {
    return b, a
  }

  return a, b
}

func fillPixel(color termbox.Attribute) {
  canvas[cursor.X][cursor.Y] = color
  canvas[cursor.X + 1][cursor.Y] = color

  draw()
}

func fillArea(color termbox.Attribute) {
  xMin, xMax := rangeLimits(visualModeFixpoint.X, cursor.X)
  yMin, yMax := rangeLimits(visualModeFixpoint.Y, cursor.Y)

  for x := xMin; x <= xMax; x++ {
    for y := yMin; y <= yMax; y++ {
      canvas[x][y] = color
      canvas[x + 1][y] = color
    }
  }

  mode = normalMode
  draw()
}

func appendToCommand(char rune) {
  command = append(command, char)
  draw()
}

func truncateCommand() {
  if len(command) > 0 {
    command = command[:len(command) - 1]
    draw()
  }
}

func executeCommand() bool {
  switch string(command) {
  case "q":
    return true
  default:
    return false
  }
}

func switchToCommandMode() {
  mode = commandMode
  draw()
}

func switchToNormalMode() {
  mode = normalMode
  draw()
}

func switchToVisualBlockMode() {
  mode = visualBlockMode
  visualModeFixpoint.X = cursor.X
  visualModeFixpoint.Y = cursor.Y
  draw()
}

func switchToColorSelectMode() {
  mode = colorSelectMode
  draw()
}
/* KEY-MAPPING */

func pollEvents() {
loop:
	for {
		switch event := termbox.PollEvent(); event.Type {
		case termbox.EventKey:
      if mode == commandMode {
        if event.Ch != 0 {
          appendToCommand(event.Ch)
        }
      } else {
        switch event.Ch {
        case '0':
          jumpToBeginningOfLine()
        case '$':
          jumpToEndOfLine()
        case ':':
          switchToCommandMode()
        case 'b':
          moveCursor('X', -10)
        case 'c':
          switchToColorSelectMode()
        case 'g':
          jumpToFirstLine()
        case 'G':
          jumpToLastLine()
        case 'h':
          moveCursor('X', -2)
        case 'j':
          moveCursor('Y', 1)
        case 'k':
          moveCursor('Y', -1)
        case 'J':
          adjustColor(+1)
        case 'K':
          adjustColor(-1)
        case 'l':
          moveCursor('X', +2)
        case 's':
          selectColor(canvas[cursor.X][cursor.Y])
        case 'w':
          moveCursor('X', +10)
        case 'x':
          if mode == normalMode {
            fillPixel(termbox.ColorDefault)
          } else if mode == visualBlockMode {
            fillArea(termbox.ColorDefault)
          }
        }
      }
      switch event.Key {
      case termbox.KeyBackspace, termbox.KeyBackspace2:
        truncateCommand()
      case termbox.KeyCtrlU:
        moveCursor('Y', -5)
      case termbox.KeyCtrlD:
        moveCursor('Y', +5,)
      case termbox.KeyCtrlV:
        switchToVisualBlockMode()
      case termbox.KeyEnter:
        err := executeCommand()
        if err { break loop }
      case termbox.KeyEsc:
        switchToNormalMode()
      case termbox.KeySpace:
        if mode == normalMode {
          fillPixel(termbox.Attribute(selectedColor))
        } else if mode == visualBlockMode {
          fillArea(termbox.Attribute(selectedColor))
        } else if mode == colorSelectMode {
          selectColor(palette[cursor.X][cursor.Y])
        }
      }
		case termbox.EventResize:
      draw()
		}
	}
}

func parseFlags() {
  flag.IntVar(&canvasRows, "rows", canvasRowsDefault, "number of rows on your canvas, 0 means full height")
  flag.IntVar(&canvasColumns, "cols", canvasColumnsDefault, "number of columns on your canvas, 0 means full width")

  flag.Parse()
}

func calculateCanvasSize() {
  if canvasRows < 0 { canvasRows = canvasRowsDefault }
  if canvasColumns < 0 { canvasColumns = canvasColumnsDefault }

  terminalWidth, terminalHeight := termbox.Size()
  if canvasColumns == 0 { canvasColumns = terminalWidth / 2 }
  if canvasRows == 0 { canvasRows = terminalHeight - 1 }

  canvasColumns = canvasColumns * 2
}

func createCanvas() {
  canvas = make([][]termbox.Attribute, canvasColumns)
  for x := 0; x < canvasColumns; x++ {
    canvas[x] = make([]termbox.Attribute, canvasRows)
  }
}

func createPalette() {
  colorIndex := 1

  palette = make([][]termbox.Attribute, canvasColumns)
  for x := 0; x < canvasColumns; x++ {
    palette[x] = make([]termbox.Attribute, canvasRows)
  }

loop:
  for y := 0; y < canvasRows; y++ {
    for x := 0; x < canvasColumns; x += 2 {
      if colorIndex > 256 {
        break loop
      }
      palette[x][y] = termbox.Attribute(colorIndex)
      palette[x + 1][y] = termbox.Attribute(colorIndex)
      colorIndex += 1
    }
  }
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
  termbox.SetOutputMode(termbox.Output256)

  parseFlags()
  calculateCanvasSize()
  createCanvas()
  createPalette()

	defer termbox.Close()
  termbox.HideCursor()

  draw()
  pollEvents()
}
