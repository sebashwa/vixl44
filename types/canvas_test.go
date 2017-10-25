package types

import (
  "testing"
  "github.com/nsf/termbox-go"
  "github.com/sebashwa/iwouldlove"
)

func TestCanvasConvertToANSI(t *testing.T) {
  idLove, it, _ := iwouldlove.Init(t)

  it("returns a byte slice of the current canvas encoded in ANSI", func() {
    canvasValues := [][]termbox.Attribute{
      []termbox.Attribute { termbox.Attribute(0), termbox.Attribute(1) },
      []termbox.Attribute { termbox.Attribute(2), termbox.Attribute(3) },
    }

    canvas := Canvas{
      Values: canvasValues,
      Columns: 2,
      Rows: 2,
    }

    ansi := canvas.ConvertToANSI()
    expectedString := "\033[0m \033[48;5;1m \033[0m\n\033[48;5;0m \033[48;5;2m \033[0m\n"

    idLove(ansi, "to equal", []byte(expectedString))
  })
}
