package types

import (
  "testing"
  "github.com/nsf/termbox-go"
  "github.com/sebashwa/iwouldlove"
)

var errorLocations = make(map[int]struct{})

var whiteCanvas = [][]termbox.Attribute{
  []termbox.Attribute{
    termbox.ColorWhite,
  },
}

var blackCanvas = [][]termbox.Attribute{
  []termbox.Attribute{
    termbox.ColorBlack,
  },
}

func TestAddState(t *testing.T) {
  idLove, describe, it := iwouldlove.Init(t)

  describe("Adding a canvas state to history", func () {
    it("prepends the canvas values to the history values", func () {
      history := History{
        Values: [][][]termbox.Attribute{
          whiteCanvas,
        },
      }

      history.AddCanvasState(blackCanvas)

      expectedHistory := History{
        Values: [][][]termbox.Attribute{
          blackCanvas,
          whiteCanvas,
        },
      }

      idLove(history, "to equal", expectedHistory)
    })

    it("uses the current position as the start of a new history", func () {
      history := History{
        Values: [][][]termbox.Attribute{
          blackCanvas,
          blackCanvas,
          blackCanvas,
          blackCanvas,
        },
        Position: 2,
      }

      history.AddCanvasState(whiteCanvas)

      expectedHistory := History{
        Values: [][][]termbox.Attribute{
          whiteCanvas,
          blackCanvas,
          blackCanvas,
        },
        Position: 0,
      }

      idLove(history, "to equal", expectedHistory)
    })
  })
}

func TestUndo(t *testing.T) {
  idLove, describe, it := iwouldlove.Init(t)

  describe("Undoing a change", func() {
    it("increases the history position", func() {
      history := History{
        Values: [][][]termbox.Attribute{
          whiteCanvas,
          whiteCanvas,
        },
        Position: 0,
      }

      history.Undo()

      idLove(history.Position, "to equal", 1)
    })

    it("returns an error if the history's oldest state is reached", func () {
      history := History{
        Values: [][][]termbox.Attribute{
          whiteCanvas,
          whiteCanvas,
        },
        Position: 0,
      }

      err := history.Undo()
      idLove(err, "to equal", nil)

      err = history.Undo()
      idLove(err, "to not equal", nil)
    })
  })
}

func TestRedo(t *testing.T) {
  idLove, describe, it := iwouldlove.Init(t)

  describe("Redoing a change", func () {
    it("decreases the position in history", func () {
      history := History{
        Values: [][][]termbox.Attribute{
          whiteCanvas,
          whiteCanvas,
        },
        Position: 1,
      }

      history.Redo()

      idLove(history.Position, "to equal", 0)
    })

    it("returns an error if already at newest state of history", func () {
      history := History{
        Values: [][][]termbox.Attribute{
          whiteCanvas,
        },
        Position: 1,
      }

      err := history.Redo()
      idLove(err, "to equal", nil)

      err = history.Redo()
      idLove(err, "to not equal", nil)
    })
  })
}
