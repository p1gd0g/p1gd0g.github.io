package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		command, _, _ := reader.ReadLine()
		handle(string(command))
	}
}

func handle(command string) {
	switch command {
	case "123":
		fmt.Println("you input 123")
	case "234":
		fmt.Println("you don't input 123")

	}
}
