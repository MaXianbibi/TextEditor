package main

import "fmt"

// import "fmt"

func endProgram() {
	disableRawMode(Get_raw_state().originalState)

}

func main() {
	terminalSetup()

	// defer garantit que la fonction endProgram sera appelée
	defer endProgram()


	

	for {
		c := readChar()
		if c == 'q' {
			break
		}
		fmt.Print(string(c))
	}
}
