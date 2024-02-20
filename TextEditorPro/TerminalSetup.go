package main

import (
	"golang.org/x/sys/unix"
)

func terminalSetup()  *unix.Termios  {
	origalState, err := enableRawMode()
	if err != nil {
		panic(err)
	}
	return origalState
}