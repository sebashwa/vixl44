package types

import (
  "fmt"
  "strings"
  "github.com/nsf/termbox-go"

  "github.com/sebashwa/vixl44/colors"
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

func (canvas Canvas) ConvertToSvg() string {
  fileCanvas := canvas.ConvertToFileCanvas()
  template := `<?xml version="1.0" standalone="no"?>
<svg viewBox="0 0 %v %v" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" xml:space="preserve" width="%vpx" height="%vpx">
  %s
</svg>
`

  var rects []string
  for x, col := range(fileCanvas) {
    for y := range(col) {

      rect := `<rect x="%v" y="%v" style="fill: %s;" width="1" height="1" />`
      fill := colors.MappingToHex[int(fileCanvas[x][y])]
      rects = append(rects, fmt.Sprintf(rect, x, y, fill))
    }
  }
  viewBoxX := len(fileCanvas)
  viewBoxY := len(fileCanvas[0])

  return fmt.Sprintf(
    template,
    viewBoxX,
    viewBoxY,
    viewBoxX * 10,
    viewBoxY * 10,
    strings.Join(rects, `
  `),
  )
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

