package keybindings

import (
  "github.com/nsf/termbox-go"

  "github.com/sebashwa/vixl44/modes"
  "github.com/sebashwa/vixl44/state"
  commonActions  "github.com/sebashwa/vixl44/actions"
  cursorActions  "github.com/sebashwa/vixl44/actions/cursor"
  paintActions   "github.com/sebashwa/vixl44/actions/paint"
  commandActions "github.com/sebashwa/vixl44/actions/command"
)

func CursorMovement(Ch rune, Key termbox.Key) {
  switch Ch {
  case '0':
    cursorActions.JumpToBeginningOfLine()
  case '$':
    cursorActions.JumpToEndOfLine()
  case 'b':
    cursorActions.Move('X', -10)
  case 'g':
    cursorActions.JumpToFirstLine()
  case 'G':
    cursorActions.JumpToLastLine()
  case 'h':
    cursorActions.Move('X', -2)
  case 'j':
    cursorActions.Move('Y', 1)
  case 'k':
    cursorActions.Move('Y', -1)
  case 'l':
    cursorActions.Move('X', +2)
  case 'w':
    cursorActions.Move('X', +10)
  }

  switch Key {
  case termbox.KeyCtrlU:
    cursorActions.Move('Y', -5)
  case termbox.KeyCtrlD:
    cursorActions.Move('Y', +5,)
  }
}

func PaletteMode(Ch rune, Key termbox.Key) {
  switch Ch {
  case 'q':
    commonActions.SetMode(modes.NormalMode)
  }
  switch Key {
  case termbox.KeySpace, termbox.KeyEnter:
    paintActions.SelectColor()
    commonActions.SetMode(modes.NormalMode)
  }
}

func NormalMode(Ch rune, Key termbox.Key) {
  switch Ch {
  case 'x':
    paintActions.FillPixel(termbox.ColorDefault)
  case 's':
    paintActions.SelectColor()
  }
  switch Key {
  case termbox.KeySpace, termbox.KeyEnter:
    paintActions.FillPixel(state.SelectedColor)
  }
}

func VisualBlockMode(Ch rune, Key termbox.Key) {
  switch Ch {
  case 'x':
    paintActions.FillArea(termbox.ColorDefault)
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
    shouldQuit, hint, errMsg := commandActions.Execute()

    if errMsg != "" {
      commonActions.SetError(errMsg)
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

