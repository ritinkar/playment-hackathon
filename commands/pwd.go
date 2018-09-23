package commands

import (
	"fmt"
	"playment/session"
)

// Pwd : Prints current working directory
func Pwd(currentSession *session.Session, args []string) {
	if len(args) > 0 {
		fmt.Println("ERR: TOO MANY ARGS. PWD TAKES NO ARGUMENTS")
	} else {
		fmt.Println("PATH: ", currentSession.GetCwd())
	}
}
