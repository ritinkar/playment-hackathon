package commands

import (
	"fmt"
	"os"
	"playment/session"
)

// Exit : exits the program
func Exit(session *session.Session, args []string) {
	if len(args) > 0 {
		fmt.Println("ERR: TOO MANY ARGS. EXIT TAKES NO ARGUMENTS")
	} else {
		os.Exit(1)

	}

}
