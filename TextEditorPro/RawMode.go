// # RawMode.go
// This file contains the code for enabling and disabling raw mode
// 

package main

import (
	"os"
	"syscall"
	"github.com/pkg/term/termios"
	"golang.org/x/sys/unix"
)

func enableRawMode() (*unix.Termios, error) {
	var originalState unix.Termios
	fd := os.Stdin.Fd()
	if err := termios.Tcgetattr(fd, &originalState); err != nil {
		return nil, err
	}
	newState := originalState
	// Turn off ECHO and ICANON
	newState.Lflag &^= syscall.ECHO | syscall.ICANON

	if err := termios.Tcsetattr(fd, termios.TCSAFLUSH, &newState); err != nil {
		return nil, err
	}
	return &originalState, nil
}

func disableRawMode(originalState *unix.Termios) error {
	fd := os.Stdin.Fd()
	return termios.Tcsetattr(fd, termios.TCSAFLUSH, originalState)
}

func readChar() byte {
	buffer := make([]byte, 1)
	os.Stdin.Read(buffer)
	return buffer[0]
}
