package main

import (
  "github.com/nsf/termbox-go"
)

type Palette struct {
  Values [][]termbox.Attribute
  LightColors map[termbox.Attribute]struct{}
}

func (palette Palette) Draw() {
  for x, column := range palette.Values {
    for y := range column {
      color := termbox.Attribute(palette.Values[x][y])
      termbox.SetCell(x, y, ' ', color, color)
    }
  }
}

func createLightColors() map[termbox.Attribute]struct{} {
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

  return Palette{values, createLightColors()}
}
