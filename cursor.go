package main

import (
  "github.com/nsf/termbox-go"
)

type Vertex struct {
  X int
  Y int
}

type Cursor struct {
  Position Vertex
  VisualModeFixpoint Vertex
}

func drawCursor(x, y int) {
  cursorColor := termbox.ColorWhite
  var backgroundColor termbox.Attribute

  if app.CurrentMode == modes.ColorSelectMode {
    backgroundColor = app.Palette.Values[x][y]
  } else {
    backgroundColor = app.Canvas.Values[x][y]
  }

  if _, exists := app.Palette.LightColors[backgroundColor]; exists {
    cursorColor = termbox.ColorBlack
  }

  termbox.SetCell(x, y, '[', cursorColor, backgroundColor)
  termbox.SetCell(x + 1, y, ']', cursorColor, backgroundColor)
}

func (cursor Cursor) Draw() {
  drawCursor(cursor.Position.X, cursor.Position.Y)
}

func (cursor Cursor) DrawBox() {
  position := cursor.Position
  fixpoint := cursor.VisualModeFixpoint

  drawCursor(position.X, position.Y)
  drawCursor(fixpoint.X, fixpoint.Y)
  drawCursor(fixpoint.X, position.Y)
  drawCursor(position.X, fixpoint.Y)
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

func (cursor *Cursor) Move(direction rune, diff int) {
  if direction == 'X' {
    newPosition := newPosition(cursor.Position.X, diff, app.Canvas.Columns - 1)
    cursor.Position.X = newPosition
  } else if direction == 'Y' {
    cursor.Position.Y = newPosition(cursor.Position.Y, diff, app.Canvas.Rows)
  }

  if app.CurrentMode != modes.VisualBlockMode {
    cursor.VisualModeFixpoint.X = cursor.Position.X
    cursor.VisualModeFixpoint.Y = cursor.Position.Y
  }
}

func (cursor *Cursor) JumpToEndOfLine() {
  diff := (app.Canvas.Columns - 2) - cursor.Position.X
  cursor.Move('X', diff)
}

func (cursor *Cursor) JumpToBeginningOfLine() {
  cursor.Move('X', -cursor.Position.X)
}

func (cursor *Cursor) JumpToFirstLine() {
  cursor.Move('Y', -cursor.Position.Y)
}

func (cursor *Cursor) JumpToLastLine() {
  diff := (app.Canvas.Rows - 1) - cursor.Position.Y
  cursor.Move('Y', diff)
}

func cursorMovementKeyMapping(Ch rune, Key termbox.Key) {
  switch Ch {
  case '0':
    app.Cursor.JumpToBeginningOfLine()
  case '$':
    app.Cursor.JumpToEndOfLine()
  case 'b':
    app.Cursor.Move('X', -10)
  case 'g':
    app.Cursor.JumpToFirstLine()
  case 'G':
    app.Cursor.JumpToLastLine()
  case 'h':
    app.Cursor.Move('X', -2)
  case 'j':
    app.Cursor.Move('Y', 1)
  case 'k':
    app.Cursor.Move('Y', -1)
  case 'l':
    app.Cursor.Move('X', +2)
  case 'w':
    app.Cursor.Move('X', +10)
  }

  switch Key {
  case termbox.KeyCtrlU:
    app.Cursor.Move('Y', -5)
  case termbox.KeyCtrlD:
    app.Cursor.Move('Y', +5,)
  }
}
