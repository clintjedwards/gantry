// Gantry is a tool for manipulating docker infrastructure over ssh

// It uses ssh tunneling to establish a connection to the docker socket file
// and then drops the user into an interactive terminal with the correct
// docker host env variable. Allowing the local docker engine to communicate with the
// remote docker instance securely.

package main

func main() {

	arguments := argumentParse()
	go establishTunnel(arguments.remoteUsername, arguments.remoteURL, arguments.localPort)
	spawnInteractiveShell(arguments)

}
