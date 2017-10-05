package main

import (
  "github.com/nsf/termbox-go"
)

type StatusBar struct {
  position int
}

func (statusBar StatusBar) Draw() {
  for i, character := range app.CurrentMode {
    termbox.SetCell(i, statusBar.position, character, app.SelectedColor, termbox.ColorDefault)
  }

  for x := len(app.CurrentMode) + 1; x < app.Canvas.Columns; x++ {
    termbox.SetCell(x, statusBar.position, ' ', app.SelectedColor, app.SelectedColor)
  }
}

func adjustColor(diff int) {
  newIndex := int(app.SelectedColor) + diff

  if newIndex < 1 {
    app.SelectedColor = 256
  } else if newIndex > 256 {
    app.SelectedColor = 1
  } else {
    app.SelectedColor = termbox.Attribute(newIndex)
  }

  draw()
}

func statusBarKeyMapping(Ch rune) {
  switch Ch {
  case 'J':
    adjustColor(+1)
  case 'K':
    adjustColor(-1)
  }
}

