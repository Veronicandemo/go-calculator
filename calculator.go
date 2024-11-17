package main

import (
	"fmt"
	"github.com/Veronicandemo/calculator/compute"
	"io"
	"os"
	"strconv"
	"strings"
)

import (
	"golang.org/x/crypto/ssh/terminal"
)

// Stores the state of the terminal before making it raw
var regularState *terminal.State

func main() {
	fmt.Println("Go Calculator Shell Mode")
	fmt.Println("Type 'exit' to quit.")

	if len(os.Args) > 1 {
		input := strings.Replace(strings.Join(os.Args[1:], ""), " ", "", -1)
		res, err := compute.Evaluate(input)
		if err != nil {
			return
		}
		fmt.Printf("%s\n", strconv.FormatFloat(res, 'G', -1, 64))
		return
	}

	var err error
	//Change terminal into raw mode so that the input can be processed one key at a time
	regularState, err = terminal.MakeRaw(0)
	//If there is an error the program terminates with an error message
	if err != nil {
		panic(err)
	}

	//Ensures that the terminal state is restored to its original state when the program exits
	defer terminal.Restore(0, regularState)

	//Creates a new terminal that reads from os.Stdin and uses > as prompt
	term := terminal.NewTerminal(os.Stdin, "> ")
	term.AutoCompleteCallback = handleKey
	// loop that will read input from the user until the program exits
	for {
		text, err := term.ReadLine()
		if err != nil {
			if err == io.EOF {
				// Quit without error on Ctrl^D
				exit()
			}
			panic(err)
		}

		text = strings.Replace(text, " ", "", -1)
		if text == "exit" || text == "quit" {
			break
		}

		res, err := compute.Evaluate(text)
		if err != nil {
			term.Write([]byte(fmt.Sprintln("Error: " + err.Error())))
			continue
		}
		term.Write([]byte(fmt.Sprintln(strconv.FormatFloat(res, 'G', -1, 64))))
	}
}

// Handles key events in the terminal
func handleKey(line string, pos int, key rune) (newLine string, newPos int, ok bool) {
	if key == '\x03' {
		// Quit without error on Ctrl^C
		exit()
	}
	return "", 0, false
}

func exit() {
	terminal.Restore(0, regularState)
	fmt.Println()
	os.Exit(0)
}
