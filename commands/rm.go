package commands

import (
	"fmt"
	"playment/session"
)

// Rm : Removes current working directories
func Rm(currentSession *session.Session, args []string) {
	if len(args) < 1 {
		fmt.Println("ERR: TOO FEW ARGS. RM TAKES AT LEAST ONE ARGUMENT")
	} else {
		atLeastOneDirDeleted, errorDirs := currentSession.RemoveNodes(args)
		if len(errorDirs) != 0 && atLeastOneDirDeleted {
			fmt.Println("ERR: SOME OF THE DIRECTORIES PATHS ARE WRONG. THE REST HAVE BEEN DELETED")
		} else if len(errorDirs) != 0 && !atLeastOneDirDeleted {
			fmt.Println("ERR: THE DIRECTORIES PATHS ARE WRONG")
		} else {
			fmt.Println("SUCC: THE DIRECTORIES HAVE BEEN DELETED")
		}
	}
}
