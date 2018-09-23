package commands

import (
	"fmt"
	"playment/session"
)

// Session : Prints all the children of the current directory
func Session(currentSession *session.Session, args []string) {
	if len(args) < 1 {
		fmt.Println("ERR: TOO FEW ARGS. SESSION TAKES EXACTLY ONE ARGUMENT")
	} else if len(args) > 1 {
		fmt.Println("ERR: TOO MANY ARGS. SESSION TAKES EXACTLY ONE ARGUMENT")
	} else if args[0] != "clear" {
		fmt.Println("ERR: UNRECOGNIZED ARG : ", args[0])
	} else {
		currentSession.Clear()
		fmt.Println("SUCC: CLEARED: RESET TO ROOT")
	}

}
