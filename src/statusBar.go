package main

import (
  "github.com/nsf/termbox-go"
)

type StatusBar struct {
  Position int
  Hint string
  Error string
  Command string
}

func displayStatus(position int) {
  for i, character := range app.CurrentMode {
    termbox.SetCell(i, position, character, app.SelectedColor, termbox.ColorDefault)
  }

  for x := len(app.CurrentMode) + 1; x < app.Canvas.Columns; x++ {
    termbox.SetCell(x, position, ' ', app.SelectedColor, app.SelectedColor)
  }
}

func displayMessage(message string, position int, color termbox.Attribute) {
  for i, char := range(message) {
    termbox.SetCell(i, position, char, color, termbox.ColorDefault)
  }
}

func (statusBar StatusBar) DrawCommand() {
  termbox.SetCell(0, statusBar.Position, ':', termbox.ColorWhite, termbox.ColorDefault)

  for i, char := range(statusBar.Command) {
    termbox.SetCell(i + 1, statusBar.Position, char, termbox.ColorWhite, termbox.ColorDefault)
  }
}

func (statusBar *StatusBar) Draw() {
  if statusBar.Error != "" {
    displayMessage(statusBar.Error, statusBar.Position, termbox.ColorRed)
  } else if statusBar.Hint != "" {
    displayMessage(statusBar.Hint, statusBar.Position, termbox.ColorGreen)
  } else {
    displayStatus(statusBar.Position)
  }

  statusBar.Hint = ""
  statusBar.Error = ""
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

