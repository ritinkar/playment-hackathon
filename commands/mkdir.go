package commands

import (
	"fmt"
	"playment/session"
)

// Mkdir : Creates children of the current working directory
func Mkdir(currentSession *session.Session, args []string) {
	if len(args) < 1 {
		fmt.Println("ERR: TOO FEW ARGS. MKDIR TAKES AT LEAST ONE ARGUMENT")
	} else {
		atLeastOneDirCreated, errorDirs, invalidDirNames := currentSession.AddNodes(args)
		if len(errorDirs) != 0 && len(invalidDirNames) != 0 && atLeastOneDirCreated {
			fmt.Println("ERR: SOME OF THE DIRECTORIES ALREADY EXIST. SOME OF THE DIRECTORY NAMES HAVE A '/'.THE REST HAVE BEEN CREATED")
		} else if len(errorDirs) != 0 && len(invalidDirNames) == 0 && atLeastOneDirCreated {
			fmt.Println("ERR: SOME OF THE DIRECTORIES ALREADY EXIST. THE REST HAVE BEEN CREATED")
		} else if len(errorDirs) != 0 && len(invalidDirNames) == 0 && !atLeastOneDirCreated {
			fmt.Println("ERR: DIRECTORY ALREADY EXISTS")
		} else if len(errorDirs) == 0 && len(invalidDirNames) != 0 && atLeastOneDirCreated {
			fmt.Println("ERR: SOME OF THE DIRECTORY NAMES HAVE A '/'. THE REST HAVE BEEN CREATED")
		} else if len(errorDirs) == 0 && len(invalidDirNames) != 0 && !atLeastOneDirCreated {
			fmt.Println("ERR: DIRECTORY NAME HAS A '/'. ")
		} else if len(errorDirs) == 0 && len(invalidDirNames) == 0 && atLeastOneDirCreated {
			fmt.Println("SUCC: CREATED")
		}

	}
}
