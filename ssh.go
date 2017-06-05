//Provides functions to create local port,
//                        ssh into remote server,
//                        connect to docker socket and
//                        spawn terminal that has the proper docker env variables to talk to daemon

package main

import "fmt"
import "net"
import "os"

import "golang.org/x/crypto/ssh"
import "golang.org/x/crypto/ssh/agent"

func SSHAgent() ssh.AuthMethod {

    if sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		return ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers)
	}
	return nil

}


func connect_to_remote(username string, hostname string) {

    sshConfig := &ssh.ClientConfig{
    	User: username,
    	Auth: []ssh.AuthMethod{SSHAgent()},
    }

    connection, err := ssh.Dial("tcp", hostname + ":" + "22", sshConfig)
    if err != nil {
    	fmt.Errorf("Failed to dial: %s", err)
    }

    fmt.Println(connection)
}
