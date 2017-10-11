package state

import (
  "github.com/nsf/termbox-go"

  "../types"
)

var Canvas types.Canvas
var Palette types.Palette
var StatusBar types.StatusBar
var Cursor types.Cursor
var SelectedColor termbox.Attribute
var CurrentMode string
var Filename string

