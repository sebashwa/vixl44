package types

import (
  "github.com/nsf/termbox-go"
)

type Palette [][]termbox.Attribute

type StatusBar struct {
  Position int
  Hint string
  Error string
  Command string
}

type vertex struct {
  X int
  Y int
}

type Cursor struct {
  Position vertex
  VisualModeFixpoint vertex
}

type File struct {
  Canvas [][]termbox.Attribute `json:"canvas"`
}

