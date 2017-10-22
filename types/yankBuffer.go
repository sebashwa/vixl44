package types

import (
  "github.com/nsf/termbox-go"
)

type YankBuffer struct {
  Values [][]termbox.Attribute
}

func (yankBuffer *YankBuffer) Set(xMin, xMax, yMin, yMax int, values [][]termbox.Attribute) {
  newContent := make([][]termbox.Attribute, xMax - xMin + 2)

  for x := xMin; x <= xMax + 1; x++ {
    column := make([]termbox.Attribute, yMax - yMin + 1)
    newContent[x - xMin] = column

    for y := yMin; y <= yMax; y++ {
      newContent[x - xMin][y - yMin] = values[x][y]
    }
  }

  yankBuffer.Values = newContent
}
