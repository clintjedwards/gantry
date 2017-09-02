// Gantry is a tool for manipulating docker infrastructure over ssh

// It uses ssh tunneling to establish a connection to the docker socket file
// and then drops the user into an interactive terminal with the correct
// docker host env variable. Allowing the local docker engine to communicate with the
// remote docker instance securely.

package main

import (
	"fmt"
	"os"

	prompt "github.com/c-bata/go-prompt"
)

func main() {

	arguments := argumentParse()
	go establishTunnel(arguments.remoteUsername, arguments.remoteURL, arguments.localPort)
	os.Setenv("DOCKER_HOST", "tcp://localhost:"+arguments.localPort)

	fmt.Printf("Gantry: Manage Docker over SSH 🐳\n")
	fmt.Printf("Connected via tcp://localhost:%s to %s as user %s\n", arguments.localPort, arguments.remoteURL, arguments.remoteUsername)
	defer fmt.Println("Exited Docker interactive environment.")

	p := prompt.New(
		executor,
		completer,
		prompt.OptionTitle("Gantry: Manage docker over ssh 🐳"),
		prompt.OptionPrefix(fmt.Sprintf("[%s] ", arguments.remoteURL)),
		prompt.OptionPrefixTextColor(prompt.DarkBlue),
	)

	p.Run()

}
