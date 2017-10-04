package main

import (
  "github.com/nsf/termbox-go"
)

type CommandBar struct {
  position int
  value []rune
}

func (commandBar CommandBar) draw() {
  termbox.SetCell(0, commandBar.position, ':', termbox.ColorWhite, termbox.ColorDefault)

  for i, char := range(commandBar.value) {
    termbox.SetCell(i + 1, commandBar.position, char, termbox.ColorWhite, termbox.ColorDefault)
  }
}

