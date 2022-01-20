package src

import (
	"syscall"
)

var (
	tabKey       = uintptr(0x09)
	upArrowKey   = uintptr(0x26)
	downArrowKey = uintptr(0x28)
	qKey         = uintptr(0x51)
)

type KeyReader interface {
	readKey(keyCode uintptr) bool
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
	state := initializeState()

	drawState(state)

	// Sometimes random stuff shows up in the channel on initialization.
	for len(keyChan) > 0 {
		<-keyChan
	}

	for true {
		key := <-keyChan
		if key == 'u' {
			state.moveCursorUp()
		}

		if key == 'd' {
			state.moveCursorDown()
		}

		drawState(state)
	}
}

func testKey(reader KeyReader, key uintptr, pressed *bool, val byte, ch chan byte) {
	keyPressed := reader.readKey(key)

	if keyPressed && *pressed == false {
		*pressed = true
		ch <- val
	}

	if !keyPressed {
		*pressed = false
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
	tabPressed := false

	for shouldContinue {
		if reader.readKey(qKey) {
			shouldContinue = false
		}

		testKey(reader, upArrowKey, &upPressed, 'u', keyChan)
		testKey(reader, downArrowKey, &downPressed, 'd', keyChan)
		testKey(reader, tabKey, &tabPressed, 't', keyChan)
	}
}
