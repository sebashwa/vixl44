package main

import (
  "github.com/nsf/termbox-go"
)

type Vertex struct {
  X int
  Y int
}

var cursor = Vertex{0,0}
var visualModeFixpoint = Vertex{0,0}

func drawCursor(x, y int) {
  cursorColor := termbox.ColorWhite
  var backgroundColor termbox.Attribute

  if app.CurrentMode == modes.ColorSelectMode {
    backgroundColor = app.Palette.Values[x][y]
  } else {
    backgroundColor = app.Canvas.Values[x][y]
  }

  if backgroundColor == termbox.ColorWhite {
    cursorColor = termbox.ColorBlack
  }

  termbox.SetCell(x, y, '[', cursorColor, backgroundColor)
  termbox.SetCell(x + 1, y, ']', cursorColor, backgroundColor)
}

func drawNormalCursor() {
  drawCursor(cursor.X, cursor.Y)
}

func drawVisualBlockCursor() {
  drawCursor(cursor.X, cursor.Y)
  drawCursor(visualModeFixpoint.X, visualModeFixpoint.Y)
  drawCursor(visualModeFixpoint.X, cursor.Y)
  drawCursor(cursor.X, visualModeFixpoint.Y)
}


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
    cursor.X = newPosition(cursor.X, diff, app.Canvas.Columns - 1)

  } else if direction == 'Y' {
    cursor.Y = newPosition(cursor.Y, diff, app.Canvas.Rows)
  }

  if app.CurrentMode != modes.VisualBlockMode {
    visualModeFixpoint.X = cursor.X
    visualModeFixpoint.Y = cursor.Y
  }
}

func jumpToEndOfLine() {
  cursor.X = app.Canvas.Columns - 2
}

func jumpToBeginningOfLine() {
  cursor.X = 0
}

func jumpToFirstLine() {
  cursor.Y = 0
}

func jumpToLastLine() {
  cursor.Y = app.Canvas.Rows - 1
}

func cursorMovementKeyMapping(Ch rune, Key termbox.Key) {
  switch Ch {
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
  case 'l':
    moveCursor('X', +2)
  case 'w':
    moveCursor('X', +10)
  }

  switch Key {
  case termbox.KeyCtrlU:
    moveCursor('Y', -5)
  case termbox.KeyCtrlD:
    moveCursor('Y', +5,)
  }
}
