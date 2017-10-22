package types

import (
  "bytes"
  "fmt"
  "image"
  "image/png"
  "strings"
  "github.com/nsf/termbox-go"

  "github.com/sebashwa/vixl44/colors"
)

type Canvas struct {
  Values  [][]termbox.Attribute
  Rows    int
  Columns int
}

func (canvas Canvas) GetValuesCopy() [][]termbox.Attribute {
  values := make([][]termbox.Attribute, len(canvas.Values))
  copy(values, canvas.Values)

  for i := range values {
    column := make([]termbox.Attribute, len(canvas.Values[i]))
    copy(column, canvas.Values[i])
    values[i] = column
  }

  return values
}

func (canvas Canvas) ConvertToFileCanvas() [][]termbox.Attribute {
  fileCanvas := make([][]termbox.Attribute, canvas.Columns/2)

  for x := range fileCanvas {
    fileCanvas[x] = make([]termbox.Attribute, canvas.Rows)
  }

  for y := range fileCanvas[0] {
    for x := 0; x < canvas.Columns; x += 2 {
      fileCanvas[x/2][y] = canvas.Values[x][y]
    }
  }

  return fileCanvas
}

func (canvas Canvas) ConvertToSvg() ([]byte, error) {
  fileCanvas := canvas.ConvertToFileCanvas()
  template := `<?xml version="1.0" standalone="no"?>
<svg viewBox="0 0 %v %v" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" xml:space="preserve" width="%vpx" height="%vpx">
  %s
</svg>
`
  var rects []string
  for x, col := range fileCanvas {
    for y := range col {

      currentColor := fileCanvas[x][y]

      if currentColor > 0 {
        rect := `<rect x="%v" y="%v" style="fill: %s;" width="1" height="1" />`
        fill, err := colors.MapTermboxColorToColor(currentColor)

        if err != nil {
          return []byte{}, err
        }

        rects = append(rects, fmt.Sprintf(rect, x, y, fill.Hex()))
      }
    }
  }
  viewBoxX := len(fileCanvas)
  viewBoxY := len(fileCanvas[0])

  svg := fmt.Sprintf(
    template,
    viewBoxX,
    viewBoxY,
    viewBoxX*10,
    viewBoxY*10,
    strings.Join(rects, `
  `),
  )

  return []byte(svg), nil
}

func (canvas Canvas) ConvertToPng(scaleFactor int) ([]byte, error) {
  fileCanvas := canvas.ConvertToFileCanvas()
  viewBoxX := len(fileCanvas) * scaleFactor
  viewBoxY := len(fileCanvas[0]) * scaleFactor
  rgbaCanvas := image.NewRGBA(image.Rect(0, 0, viewBoxX, viewBoxY))

  for x, column := range fileCanvas {
    for y := range column {
      currentColor := fileCanvas[x][y]

      if currentColor > 0 {
        fill, err := colors.MapTermboxColorToColor(currentColor)

        if err != nil {
          return []byte{}, err
        }

        for i := 0; i < scaleFactor; i++ {
          for j := 0; j < scaleFactor; j++ {
            rgbaCanvas.Set(x * scaleFactor + i, y * scaleFactor + j, fill)
          }
        }
      }
    }
  }

  buf := new(bytes.Buffer)
  err := png.Encode(buf, rgbaCanvas)

  if err != nil {
    return []byte{}, err
  }

  return buf.Bytes(), nil
}

func (canvas Canvas) ConvertToANSI() []byte {
  var buffer bytes.Buffer

  for y := 0; y < canvas.Rows; y++ {
    for x := 0; x < canvas.Columns; x++ {
      var stringValue string
      fill := canvas.Values[x][y]

      if fill == 0 {
        stringValue = "\x1b[0m "
      } else {
        stringValue = fmt.Sprintf("\033[48;5;%vm ", fill - 1)
      }

      buffer.WriteString(stringValue)
    }
    buffer.WriteString("\n")
  }

  return buffer.Bytes()
}
