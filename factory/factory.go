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
  for y := range(values) {
    for x := 0; x < columns; x += 2 {
      if colorIndex > 256 {
        break loop
      }
      values[x][y] = termbox.Attribute(colorIndex)
      values[x + 1][y] = termbox.Attribute(colorIndex)
      colorIndex += 1
    }
  }

  return types.Palette{Values: values, LightColors: getLightColors()}
}

func getLightColors() map[termbox.Attribute]struct{} {
  colors := make(map[termbox.Attribute]struct{})
  colorNumbers := []int{
    3,4,7,8,15,16,47,48,49,50,51,52,
    83,84,85,86,87,88,120,121,122,123,124,
    155,156,157,158,159,160,185,186,187,188,
    189,190,191,192,193,194,195,196,220,221,
    222,223,224,225,226,227,228,229,230,231,232,
    251,252,253,254,255,256,
  }

  for _, colorNumber := range(colorNumbers) {
    colors[termbox.Attribute(colorNumber)] = struct{}{}
  }

  return colors
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
