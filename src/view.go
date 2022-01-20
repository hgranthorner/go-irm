package src

import "fmt"

func drawState(s State) {
	for i := 0; i < 100; i++ {
		fmt.Println("")
	}

	for i, node := range s.nodes {
		if node.file != nil {
			firstCharacter := " "

			if i == s.y {
				firstCharacter = "*"
			}

			symbol := "[ ]"

			if node.file.IsDir() {
				if node.open {
					symbol = " v "
				} else {
					symbol = " > "
				}
			}

			fmt.Printf("%s %s %s\n", firstCharacter, symbol, node.file.Name())
		}
	}
}
