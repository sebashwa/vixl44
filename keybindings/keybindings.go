package keybindings

import (
  "github.com/nsf/termbox-go"

  "github.com/sebashwa/vixl44/modes"
  "github.com/sebashwa/vixl44/state"
  commonActions  "github.com/sebashwa/vixl44/actions"
  paintActions   "github.com/sebashwa/vixl44/actions/paint"
  commandActions "github.com/sebashwa/vixl44/actions/command"
)

func NormalMode(Ch rune, Key termbox.Key) {
  switch Ch {
  case 'x':
    paintActions.FillPixel(termbox.ColorDefault)
  case 's':
    paintActions.SelectColor()
  case 'f':
    paintActions.FloodFill()
  case 'u':
    commonActions.Undo()
  case 'p':
    commonActions.Paste()
  }

  switch Key {
  case termbox.KeyCtrlR:
    commonActions.Redo()
  case termbox.KeySpace, termbox.KeyEnter:
    paintActions.FillPixel(state.SelectedColor)
  }
}

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
    paintActions.FillArea(state.SelectedColor)
    commonActions.SetMode(modes.NormalMode)
  }
}

func CommandMode(Ch rune, Key termbox.Key) bool {
  if Ch != 0 {
    commandActions.Append(Ch)
  }

  switch Key {
  case termbox.KeyBackspace, termbox.KeyBackspace2:
    commandActions.Truncate()
  case termbox.KeySpace:
    commandActions.Append(' ')
  case termbox.KeyEnter:
    shouldQuit, hint, err := commandActions.Execute()

    if err != nil {
      commonActions.SetError(err.Error())
    } else if shouldQuit {
      return true
    } else if hint != "" {
      commonActions.SetHint(hint)
    }

    commandActions.Set("")
    commonActions.SetMode(modes.NormalMode)
  default:
    ModeSelection('0', Key)
  }

  return false
}

func ModeSelection(Ch rune, Key termbox.Key) {
  switch Key {
  case termbox.KeyCtrlV:
    commonActions.SetMode(modes.VisualBlockMode)
  case termbox.KeyEsc:
    commonActions.SetMode(modes.NormalMode)
  }

  switch Ch {
  case ':':
    commonActions.SetMode(modes.CommandMode)
  case 'c':
    commonActions.SetMode(modes.PaletteMode)
  }
}

func Common(Ch rune) {
  switch Ch {
  case 'J':
    paintActions.AdjustColor(+1)
  case 'K':
    paintActions.AdjustColor(-1)
  }
}

