package types

import (
	"errors"
	"github.com/nsf/termbox-go"
)

type History struct {
	Values   [][][]termbox.Attribute
	Position int
}

func (history *History) AddCanvasState(canvasState [][]termbox.Attribute) {
	historyTail := history.Values[history.Position:]

	history.Values = append([][][]termbox.Attribute{canvasState}, historyTail...)
	history.Position = 0
}

func (history *History) Undo() error {
	newPosition := history.Position + 1
	if newPosition > len(history.Values)-1 {
		return errors.New("reached oldest canvas state")
	} else {
		history.Position = newPosition
		return nil
	}
}

func (history *History) Redo() error {
	newPosition := history.Position - 1
	if newPosition < 0 {
		return errors.New("reached newest canvas state")
	} else {
		history.Position = newPosition
		return nil
	}
}

func (history History) GetCurrentCanvasValuesCopy() [][]termbox.Attribute {
	canvas := Canvas{
		Values: history.Values[history.Position],
	}

	return canvas.GetValuesCopy()
}
