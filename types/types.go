package types

import "github.com/nsf/termbox-go"

type Palette [][]termbox.Attribute
type Mode string

type StatusBar struct {
	Position int
	Hint     string
	Error    string
	Command  string
}

type Vertex struct {
	X int
	Y int
}

type File struct {
	Canvas [][]termbox.Attribute `json:"canvas"`
}
