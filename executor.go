package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func fullDeploy() {
	command1 := fmt.Sprintf("docker-compose -f docker-compose.production.yml pull")
	command2 := fmt.Sprintf("docker-compose -f docker-compose.production.yml down")
	command3 := fmt.Sprintf("docker-compose -f docker-compose.production.yml up -d")
	cmd := exec.Command("/bin/sh", "-c", command1, "&&", command2, "&&", command3)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}
}

func halfDeploy() {
	command1 := fmt.Sprintf("docker-compose -f docker-compose.production.yml pull")
	command2 := fmt.Sprintf("docker-compose -f docker-compose.production.yml up -d")
	cmd := exec.Command("/bin/sh", "-c", command1, "&&", command2)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}
}

func executor(input string) {
	input = strings.TrimSpace(input)

	switch input {
	case "":
		return
	case "quit":
		os.Exit(0)
		return
	case "exit":
		os.Exit(0)
		return
	case "full-deploy":
		fullDeploy()
	case "half-deploy":
		halfDeploy()
	default:
		cmd := exec.Command("/bin/sh", "-c", input)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("Got error: %s\n", err.Error())
		}
	}

	return
}
