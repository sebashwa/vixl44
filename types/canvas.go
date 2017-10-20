package types

import (
	"fmt"
	"image"
	"strings"

	"github.com/nsf/termbox-go"

	"bytes"
	"image/png"

	"image/color"

	"github.com/lucasb-eyer/go-colorful"
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

func (canvas Canvas) ConvertToSvg() string {
	fileCanvas := canvas.ConvertToFileCanvas()
	template := `<?xml version="1.0" standalone="no"?>
<svg viewBox="0 0 %v %v" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" xml:space="preserve" width="%vpx" height="%vpx">
  %s
</svg>
`
	var rects []string
	for x, col := range fileCanvas {
		for y := range col {

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
		viewBoxX*10,
		viewBoxY*10,
		strings.Join(rects, `
  `),
	)
}

func (canvas Canvas) ConvertToPng() ([]byte, error) {
	fileCanvas := canvas.ConvertToFileCanvas()
	viewBoxX := len(fileCanvas)
	viewBoxY := len(fileCanvas[0])
	m := image.NewRGBA(image.Rect(0, 0, viewBoxX, viewBoxY))

	for x, col := range fileCanvas {
		for y := range col {
			var err error
			var clr color.Color

			cur := int(fileCanvas[x][y])

			// Ignore transparent pixels
			if cur > 0 {
				clr, err = colorful.Hex(colors.MappingToHex[int(fileCanvas[x][y])])

				if err != nil {
					return []byte{}, err
				}

				m.Set(x, y, clr)
			}
		}
	}

	buf := new(bytes.Buffer)
	err := png.Encode(buf, m)
	if err != nil {
		return []byte{}, err
	}

	return buf.Bytes(), nil
}
