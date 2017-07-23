//We set up the proper environment for the user then drop into an interactive shell
// author: http://technosophos.com/2014/07/11/start-an-interactive-shell-from-within-go.html

package main

import (
	"fmt"
	"log"
	"os"
)

func spawnInteractiveShell(arguments Arguments) {

	cwd, err := os.Getwd()
	if err != nil {
		log.Print(err)
	}

	os.Setenv("DOCKER_HOST", "tcp://localhost:"+arguments.localPort)

	// Transfer stdin, stdout, and stderr to the new process
	// and also set target directory for the shell to start in.
	pa := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Dir:   cwd,
	}

	fmt.Printf("Connected via tcp://localhost:%s to %s as user %s\n", arguments.localPort, arguments.remoteURL, arguments.remoteUsername)
	fmt.Println("Starting dockerized interactive shell 🐳")

	proc, err := os.StartProcess("/bin/bash", []string{"bash", "--rcfile", "./bashrc"}, &pa)
	if err != nil {
		log.Print(err)
	}

	// Wait until user exits the shell
	_, err = proc.Wait()
	if err != nil {
		log.Print(err)
	}

	// Alert the user that they have exited shell
	fmt.Printf("Exited dockerized shell\n")
}
