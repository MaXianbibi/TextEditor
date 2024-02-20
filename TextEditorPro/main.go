package main

import "fmt"

// import "fmt"

func main() {
	
	origalState := terminalSetup()
	defer disableRawMode(origalState)


	for {
		c := readChar()
		if c == 'q' {
			break
		}
		fmt.Print(string(c))
	}
}
