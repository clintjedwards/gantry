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
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    connection, err := ssh.Dial("tcp", hostname + ":" + "22", sshConfig)
    if err != nil {
    	fmt.Printf("Failed to dial: %s", err)
        return
    }

    session, err := connection.NewSession()
    if err != nil {
    	connection.Close()
        fmt.Println(err)
    }

    out, err := session.CombinedOutput("pwd")
    if err != nil {
        connection.Close()
        fmt.Println(err)
    }

    fmt.Println(string(out))
    connection.Close()
}
