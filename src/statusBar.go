package main

import (
  "github.com/nsf/termbox-go"
)

type StatusBar struct {
  position int
}

func (statusBar StatusBar) draw() {
  for i, character := range app.currentMode {
    termbox.SetCell(i, statusBar.position, character, app.selectedColor, termbox.ColorDefault)
  }

  for x := len(app.currentMode) + 1; x < app.canvas.columns; x++ {
    termbox.SetCell(x, statusBar.position, ' ', app.selectedColor, app.selectedColor)
  }
}

func adjustColor(diff int) {
  newIndex := int(app.selectedColor) + diff

  if newIndex < 1 {
    app.selectedColor = 256
  } else if newIndex > 256 {
    app.selectedColor = 1
  } else {
    app.selectedColor = termbox.Attribute(newIndex)
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

