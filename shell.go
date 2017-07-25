//We set up the proper environment for the user then drop into an interactive shell

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func createTmpBashrc() *os.File {

	content := []byte(bashrcContent)

	tmpfile, err := ioutil.TempFile("", "bashrc")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	return tmpfile
}

func spawnInteractiveShell(arguments Arguments) {

	cwd, err := os.Getwd()
	if err != nil {
		log.Print(err)
	}

	os.Setenv("DOCKER_HOST", "tcp://localhost:"+arguments.localPort)
	os.Setenv("GANTRY_HOST", fmt.Sprintf("%s@%s", arguments.remoteUsername, arguments.remoteURL))

	// Transfer stdin, stdout, and stderr to the new process
	// and also set target directory for the shell to start in.
	pa := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Dir:   cwd,
	}

	fmt.Printf("Connected via tcp://localhost:%s to %s as user %s\n", arguments.localPort, arguments.remoteURL, arguments.remoteUsername)
	fmt.Println("Starting dockerized interactive shell 🐳")

	rcfile := createTmpBashrc()

	proc, err := os.StartProcess("/bin/bash", []string{"bash", "--rcfile", rcfile.Name()}, &pa)
	if err != nil {
		log.Print(err)
	}

	defer os.Remove(rcfile.Name())

	// Wait until user exits the shell
	_, err = proc.Wait()
	if err != nil {
		log.Print(err)
	}

	defer fmt.Printf("Exited dockerized shell\n")
}
