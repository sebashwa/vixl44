package main

import (
  "github.com/nsf/termbox-go"
)

type Palette struct {
  Values [][]termbox.Attribute
}

func (palette Palette) Draw() {
  for x, column := range palette.Values {
    for y := range column {
      color := termbox.Attribute(palette.Values[x][y])
      termbox.SetCell(x, y, ' ', color, color)
    }
  }
}

func createPalette(rows, columns int) Palette {
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

  return Palette{values}
}
