package main

import (
  "github.com/nsf/termbox-go"
)

type CommandBar struct {
  Position int
  Value []rune
}

func (commandBar CommandBar) Draw() {
  termbox.SetCell(0, commandBar.Position, ':', termbox.ColorWhite, termbox.ColorDefault)

  for i, char := range(commandBar.Value) {
    termbox.SetCell(i + 1, commandBar.Position, char, termbox.ColorWhite, termbox.ColorDefault)
  }
}

