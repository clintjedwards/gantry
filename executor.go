package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

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
	case "half-deploy":
		command1 := fmt.Sprintf("docker-compose -f docker-compose.production.yml pull")
		command2 := fmt.Sprintf("docker-compose -f docker-compose.production.yml up -d")
		cmd := exec.Command("/bin/sh", "-c", command1, "&&", command2)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("Got error: %s\n", err.Error())
		}
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
