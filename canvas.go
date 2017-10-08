package main

import (
  "github.com/nsf/termbox-go"
)

type Canvas struct {
  Values [][]termbox.Attribute
  Rows int
  Columns int
}

func (canvas Canvas) Draw() {
  for x, column := range canvas.Values {
    for y := range column {
      color := canvas.Values[x][y]
      termbox.SetCell(x, y, ' ', color, color)
    }
  }
}

func (canvas Canvas) ConvertToFileCanvas() [][]termbox.Attribute {
  fileCanvas := make([][]termbox.Attribute, canvas.Rows)

  for y := range(fileCanvas) {
    fileCanvas[y] = make([]termbox.Attribute, canvas.Columns / 2)
  }

  for y := range(fileCanvas) {
    for x := 0; x < canvas.Columns; x += 2 {
      fileCanvas[x / 2][y] = canvas.Values[x][y]
    }
  }

  return fileCanvas
}

func createCanvasFromFileCanvas(fileCanvas [][]termbox.Attribute) Canvas {
  appCanvas := createCanvas(len(fileCanvas[0]), len(fileCanvas))

  for x, column := range(fileCanvas) {
    for y := range(column) {
      appCanvas.Values[x * 2][y] = fileCanvas[x][y]
      appCanvas.Values[(x * 2) + 1][y] = fileCanvas[x][y]
    }
  }

  return appCanvas
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
