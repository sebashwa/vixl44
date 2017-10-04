package main

import (
  "github.com/nsf/termbox-go"
)

type Canvas struct {
  values [][]termbox.Attribute
  rows int
  columns int
}

func (canvas Canvas) draw() {
  for x, column := range canvas.values {
    for y := range column {
      color := canvas.values[x][y]
      termbox.SetCell(x, y, ' ', color, color)
    }
  }
}

func createCanvas(rows, columns int) Canvas {
  rows, columns = adjustCanvasSize(rows, columns)

  values := make([][]termbox.Attribute, columns)
  for x := range(values) {
    values[x] = make([]termbox.Attribute, rows)
  }

  return Canvas{values, rows, columns}
}

func adjustCanvasSize(rows, columns int) (int, int) {
  if rows < 0 { rows = rows * -1 }
  if columns < 0 { columns = columns * -1 }

  terminalWidth, terminalHeight := termbox.Size()
  if columns == 0 { columns = terminalWidth / 2 }
  if rows == 0 { rows = terminalHeight - 1 }

  columns = columns * 2

  return rows, columns
}
