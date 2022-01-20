package src

import (
	"io/fs"
)

type Coordinate struct {
	x int
	y int
}

type State struct {
	currentPosition Coordinate
	dirEntries      []fs.DirEntry
}

func (s *State) moveCursorUp() {
	s.currentPosition.y = s.currentPosition.y - 1
}

func (s *State) moveCursorDown() {
	s.currentPosition.y = s.currentPosition.y + 1
}
