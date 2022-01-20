package src

import "fmt"

func drawState(s State) {
	for i := 0; i < 100; i++ {
		fmt.Println("")
	}

	for i, file := range s.dirEntries {
		firstCharacter := " "

		if i == s.currentPosition.y {
			firstCharacter = "*"
		}

		fmt.Printf("%s [ ] %s\n", firstCharacter, file.Name())
	}
}
