package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 { //checking if there is more than one argument
		switch strings.ToLower(os.Args[1]) {
		case "install":
			RegisterProtocol() // Resgitring the URL protocol
			break
		case "uninstall":
			UnregisterProtocol() //Unregistring
			break
		case "-open":
			fmt.Printf("link : %s", string(os.Args[2])) // showing the link sent by the browser
			break
		default:
			fmt.Printf("Available commands : install or uninstall")

		}

	}
	fmt.Scanf("Waiting for key press...")
}
