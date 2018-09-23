package commands

import (
	"fmt"
	"playment/session"
)

// Ls : Prints all the children of the current directory
func Ls(currentSession *session.Session, args []string) {
	if len(args) > 0 {
		fmt.Println("ERR: TOO MANY ARGS. LS TAKES NO ARGUMENTS")
	}

	fmt.Printf("DIRS:")
	for _, directory := range currentSession.GetChildrenDirectories() {
		fmt.Printf(" %v", directory)
	}
	fmt.Printf("\n")
}
