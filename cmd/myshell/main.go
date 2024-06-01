package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		res := strings.TrimSpace(input)
		if res == "exit 0" {
			os.Exit(0)
		} else {
			fmt.Fprintf(os.Stdout, "%s: command not found\n", res)
		}
	}

}
