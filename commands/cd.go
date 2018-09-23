package commands

import (
	"fmt"
	"playment/session"
)

// Cd : Receives a session and args and then changes current directory
func Cd(currentSession *session.Session, args []string) {
	if len(args) > 1 {
		fmt.Println("ERR: TOO MANY ARGS. CD TAKES ONLY ONE ARGUMENT")
	} else if len(args) < 1 {
		fmt.Println("ERR: TOO FEW ARGS. CD TAKES AT LEAST ONE ARGUMENT")
	} else {
		err := currentSession.SwitchCurrentDir(args[0])
		if err {
			fmt.Println("ERR: INVALID PATH")
		} else {
			fmt.Println("SUCC: REACHED")

		}
	}
}
