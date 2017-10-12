package actions

import (
  "github.com/sebashwa/vixl44/state"
)

func SetMode(mode string) {
  state.CurrentMode = mode
}

func SetError(message string) {
  state.StatusBar.Error = message
}

func SetHint(message string) {
  state.StatusBar.Hint = message
}

func SetCommand(command string) {
  state.StatusBar.Command = command
}
