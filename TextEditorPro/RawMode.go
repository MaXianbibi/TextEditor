// # RawMode.go
// This file contains the code for enabling and disabling raw mode
//

package main

import (
	"github.com/pkg/term/termios"
	"golang.org/x/sys/unix"
	"os"
	"syscall"
)

type raw_state struct {
	raw_state *unix.Termios
	originalState *unix.Termios
}

var instance *raw_state

func Get_raw_state() *raw_state {
    if instance == nil {
        instance = &raw_state{
            // initialisez vos champs ici, si nécessaire
        }
    }
    return instance
}



func enableRawMode() (*unix.Termios, error) {
	var originalState unix.Termios
	fd := os.Stdin.Fd()
	if err := termios.Tcgetattr(fd, &originalState); err != nil {
		return nil, err
	}
	newState := originalState
	// Turn off ECHO and ICANON
	newState.Iflag &^= syscall.IXON | syscall.ICRNL | syscall.BRKINT | syscall.INPCK
	newState.Oflag &^= syscall.OPOST
	newState.Lflag &^= syscall.ECHO | syscall.ICANON | syscall.ISIG | syscall.IEXTEN
	newState.Cflag |= syscall.CS8
	// Set the minimum number of bytes of input needed before read() can return
	newState.Cc[syscall.VMIN] = 0
	// Set the maximum amount of time to wait before read() returns
	newState.Cc[syscall.VTIME] = 1

	singleton := Get_raw_state()
	singleton.raw_state = &newState
	singleton.originalState = &originalState

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
	n, err := os.Stdin.Read(buffer)

	if n != 1 || err != nil {
		panic(err)
	}
	return buffer[0]
}
