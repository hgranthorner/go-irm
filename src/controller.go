package src

import (
	"fmt"
	"os"
	"syscall"
)

var (
	//tabKey       = uintptr(0x09)
	//leftArrowKey = uintptr(0x25)
	upArrowKey   = uintptr(0x26)
	downArrowKey = uintptr(0x28)
	qKey         = uintptr(0x51)
)

type KeyReader interface {
	readKey(keyCode uintptr) uintptr
}

type WindowsKeyReader struct {
	dll              *syscall.LazyDLL
	getAsyncKeyState *syscall.LazyProc
}

func (r WindowsKeyReader) readKey(keyCode uintptr) bool {
	ret, _, _ := r.getAsyncKeyState.Call(keyCode)
	return ret != 0
}

func handleInput(keyChan <-chan byte) {
	coordinate := Coordinate{0, 0}
	currentDirectoryPath, _ := os.Getwd()
	files, _ := os.ReadDir(currentDirectoryPath)

	state := State{
		currentPosition: coordinate,
		dirEntries:      files,
	}

	for true {
		key := <-keyChan
		if key == 'u' {
			fmt.Println("up")
			state.moveCursorUp()
			fmt.Println(state.currentPosition.y)
		}

		if key == 'd' {
			fmt.Println("down")
			state.moveCursorDown()
			fmt.Println(state.currentPosition.y)
		}

		drawState(state)
	}
}

func Init() {
	user32 := syscall.NewLazyDLL("user32.dll")
	procGetAsyncKeyState := user32.NewProc("GetAsyncKeyState")

	reader := WindowsKeyReader{dll: user32, getAsyncKeyState: procGetAsyncKeyState}

	keyChan := make(chan byte)

	go handleInput(keyChan)

	shouldContinue := true
	upPressed := false
	downPressed := false

	for shouldContinue {
		if reader.readKey(qKey) {
			shouldContinue = false
		}

		if reader.readKey(upArrowKey) && upPressed == false {
			upPressed = true
			keyChan <- 'u'
			continue
		}

		if reader.readKey(downArrowKey) && downPressed == false {
			downPressed = true
			keyChan <- 'd'
			continue
		}

		if !reader.readKey(upArrowKey) {
			upPressed = false
		}

		if !reader.readKey(downArrowKey) {
			downPressed = false
		}

	}
}
