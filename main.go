package main

import (
  "github.com/nsf/termbox-go"
)

type Vertex struct {
  X int
  Y int
}

const visualBlockMode = "VISUAL-BLOCK"
const normalMode = "NORMAL"
const canvasRows = 40
const canvasColumns = 80
var colors = []termbox.Attribute {
	termbox.ColorRed,
	termbox.ColorGreen,
	termbox.ColorBlue,
	termbox.ColorYellow,
	termbox.ColorCyan,
	termbox.ColorMagenta,
	termbox.ColorWhite,
	termbox.ColorBlack,
}

var canvas [canvasColumns][canvasRows] termbox.Attribute
var mode = normalMode
var cursor = Vertex{0,0}
var visualModeFixpoint = Vertex{0,0}
var selectedColorIndex = 0

/* DRAWING */

func drawCursor(x, y int) {
  cursorColor := termbox.ColorWhite
  backgroundColor := canvas[x][y]

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

func drawStatusBar() {
  selectedColor := colors[selectedColorIndex]

  for i, character := range mode {
    termbox.SetCell(i, canvasRows, character, selectedColor, termbox.ColorDefault)
  }

  for x := len(mode) + 1; x < canvasColumns; x++ {
    termbox.SetCell(x, canvasRows, ' ', selectedColor, selectedColor)
  }
}

func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

  drawCanvas()
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

func selectColor(diff int) {
  newIndex := selectedColorIndex + diff

  if newIndex < 0 {
    selectedColorIndex = len(colors) - 1
  } else if newIndex >= len(colors) {
    selectedColorIndex = 0
  } else {
    selectedColorIndex = newIndex
  }
  draw()
}

func rangeLimits(a, b int) (int, int) {
  if a > b {
    return b, a
  }

  return a, b
}

func setColor(color termbox.Attribute) {
  if mode == normalMode {
    canvas[cursor.X][cursor.Y] = colors[selectedColorIndex]
    canvas[cursor.X + 1][cursor.Y] = colors[selectedColorIndex]
  } else {
    xMin, xMax := rangeLimits(visualModeFixpoint.X, cursor.X)
    yMin, yMax := rangeLimits(visualModeFixpoint.Y, cursor.Y)

    for x := xMin; x <= xMax; x++ {
      for y := yMin; y <= yMax; y++ {
        canvas[x][y] = color
        canvas[x + 1][y] = color
      }
    }
  }

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

/* KEY-MAPPING */

func pollEvents() {
loop:
	for {
		switch event := termbox.PollEvent(); event.Type {
		case termbox.EventKey:
      switch event.Ch {
      case '0':
        jumpToBeginningOfLine()
      case '$':
        jumpToEndOfLine()
      case 'b':
        moveCursor('X', -10)
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
        selectColor(+1)
      case 'K':
        selectColor(-1)
      case 'l':
        moveCursor('X', +2)
      case 'q':
        break loop
      case 'w':
        moveCursor('X', +10)
      case 'x':
        setColor(termbox.ColorDefault)
      }
      switch event.Key {
      case termbox.KeyCtrlU:
        moveCursor('Y', -5)
      case termbox.KeyCtrlD:
        moveCursor('Y', +5,)
      case termbox.KeyCtrlV:
        switchToVisualBlockMode()
      case termbox.KeyEsc:
        switchToNormalMode()
      case termbox.KeySpace:
        setColor(colors[selectedColorIndex])
      }
		case termbox.EventResize:
      draw()
		}
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
  termbox.HideCursor()

  draw()
  pollEvents()
}
