package keybindings

import (
  "github.com/nsf/termbox-go"

   "github.com/sebashwa/vixl44/modes"
  commonActions  "github.com/sebashwa/vixl44/actions"
  paintActions   "github.com/sebashwa/vixl44/actions/paint"
)


func VisualBlockMode(Ch rune, Key termbox.Key) {
  switch Ch {
  case 'y':
    commonActions.Copy()
    commonActions.SetMode(modes.NormalMode)
  case 'd', 'x':
    commonActions.Cut()
    commonActions.SetMode(modes.NormalMode)
  }

  switch Key {
  case termbox.KeySpace, termbox.KeyEnter:
    paintActions.FillArea()
    commonActions.SetMode(modes.NormalMode)
  }
}
