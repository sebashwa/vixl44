package factory

import (
  "github.com/nsf/termbox-go"

  "github.com/sebashwa/vixl44/types"
)

func CreatePalette(columns, rows int) types.Palette {
  colorIndex := 1

  values := make([][]termbox.Attribute, columns)
  for x := range(values) {
    values[x] = make([]termbox.Attribute, rows)
  }

loop:
  for y := 0; y < rows; y += 1 {
    for x := 0; x < columns; x += 2 {
      if colorIndex > rows * columns || colorIndex > 256 {
        break loop
      }
      values[x][y] = termbox.Attribute(colorIndex)
      values[x + 1][y] = termbox.Attribute(colorIndex)
      colorIndex += 1
    }
  }

  return values
}

func CreateCanvas(columns, rows int) types.Canvas {
  columns, rows = adjustCanvasSize(columns, rows)

  values := make([][]termbox.Attribute, columns)
  for x := range(values) {
    values[x] = make([]termbox.Attribute, rows)
  }

  return types.Canvas{Values: values, Rows: rows, Columns: columns}
}

func CreateCanvasFromFileCanvas(fileCanvas [][]termbox.Attribute) types.Canvas {
  appCanvas := CreateCanvas(len(fileCanvas), len(fileCanvas[0]))

  for x, column := range(fileCanvas) {
    for y := range(column) {
      appCanvas.Values[x * 2][y] = fileCanvas[x][y]
      appCanvas.Values[(x * 2) + 1][y] = fileCanvas[x][y]
    }
  }

  return appCanvas
}

func adjustCanvasSize(columns, rows int) (int, int) {
  if rows < 0 { rows = rows * -1 }
  if columns < 0 { columns = columns * -1 }

  terminalWidth, terminalHeight := termbox.Size()
  if columns == 0 { columns = terminalWidth / 2 }
  if rows == 0 { rows = terminalHeight - 1 }

  columns = columns * 2

  return columns, rows
}
