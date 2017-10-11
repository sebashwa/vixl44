package types

import (
  "github.com/nsf/termbox-go"
)

type Canvas struct {
  Values [][]termbox.Attribute
  Rows int
  Columns int
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

type Palette struct {
  Values [][]termbox.Attribute
  LightColors map[termbox.Attribute]struct{}
}

type StatusBar struct {
  Position int
  Hint string
  Error string
  Command string
}

type vertex struct {
  X int
  Y int
}

type Cursor struct {
  Position vertex
  VisualModeFixpoint vertex
}

type File struct {
  Canvas [][]termbox.Attribute `json:"canvas"`
}

