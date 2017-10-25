package state

import (
  "github.com/nsf/termbox-go"

  "github.com/sebashwa/vixl44/types"
)

var Canvas types.Canvas
var Palette types.Palette
var StatusBar types.StatusBar
var Cursor types.Cursor
var SelectedColor termbox.Attribute
var CurrentMode types.Mode
var Filename string
var History types.History
var YankBuffer types.YankBuffer
