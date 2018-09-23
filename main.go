package main

import (
	"bufio"
	"fmt"
	"os"
	"playment/commands"
	"playment/session"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	currentSession := session.StartSession()
	fmt.Println("$> <Starting your application...>")

	for {
		ps1 := "$> "
		fmt.Printf(ps1)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		splitInput := strings.Split(input, " ")

		switch command := splitInput[0]; command {
		case "cd":
			commands.Cd(currentSession, splitInput[1:])
		case "ls":
			commands.Ls(currentSession, splitInput[1:])
		case "mkdir":
			commands.Mkdir(currentSession, splitInput[1:])
		case "pwd":
			commands.Pwd(currentSession, splitInput[1:])
		case "rm":
			commands.Rm(currentSession, splitInput[1:])
		case "exit":
			commands.Exit(currentSession, splitInput[1:])
		case "session":
			commands.Session(currentSession, splitInput[1:])
		default:
			fmt.Println("ERR: CANNOT RECOGNIZE INPUT.")
		}
	}
}
