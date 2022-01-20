package src

import (
	"io/fs"
	"os"
)

type Node struct {
	file   fs.DirEntry
	marked bool
	open   bool
}

type State struct {
	y     int
	nodes []Node
}

func initializeState() State {
	currentDirectoryPath, _ := os.Getwd()
	files, _ := os.ReadDir(currentDirectoryPath)
	nodes := make([]Node, len(files))

	for i, file := range files {
		nodes[i] = Node{
			file:   file,
			marked: false,
			open:   false,
		}
	}

	state := State{
		y:     0,
		nodes: nodes,
	}
	return state
}

func (s *State) moveCursorUp() {
	newY := s.y - 1
	if newY >= 0 {
		s.y = newY
	}
}

func (s *State) moveCursorDown() {
	newY := s.y + 1
	if newY < len(s.nodes) {
		s.y = newY
	}
}
