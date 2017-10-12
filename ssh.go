// Provides functions to handle ssh tunneling

package main

import (
	"io"
	"log"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

//SSHAgent leverages the local ssh-agent process to authenticate
func SSHAgent() ssh.AuthMethod {

	if sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		return ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers)
	}
	return nil
}

//Opens local port on host machine and listens for connections
func establishLocalListener(port string) net.Listener {

	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Print(err)
		listener.Close()
	}

	return listener
}

//Establish ssh connection to remote machine
func connectToRemote(username string, hostname string) *ssh.Client {

	sshConfig := &ssh.ClientConfig{
		User:            username,
		Auth:            []ssh.AuthMethod{SSHAgent()},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	remoteConnection, err := ssh.Dial("tcp", hostname+":"+"22", sshConfig)
	if err != nil {
		log.Fatalf("remote connectioned failed: please check valid ssh key is loaded and access to the server is permitted\n%v", err)
	}

	return remoteConnection

}

//Establish socket connection on remote machine
func connectToSocket(remoteConnection *ssh.Client, socketPath string) *net.Conn {

	socketConnection, err := remoteConnection.Dial("unix", socketPath)
	if err != nil {
		log.Fatal("Cannot connect to socket successfully. Check permissions and make sure process is running.")
	}

	return &socketConnection
}

//Allow two way communication between two connections
func copyConnectionData(writer, reader net.Conn) {
	_, err := io.Copy(writer, reader)
	if err != nil {
		log.Print(err)
	}
}

//Establish local, remote, socket connections and then allow communication
func establishTunnel(username string, hostname string, localPort string, done chan bool) {
	listener := establishLocalListener(localPort)
	remoteConnection := connectToRemote(username, hostname)
	socketConnection := connectToSocket(remoteConnection, "/var/run/docker.sock")

	done <- true

	defer remoteConnection.Close()
	defer listener.Close()

	for {
		localConnection, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		go copyConnectionData(localConnection, *socketConnection)
		go copyConnectionData(*socketConnection, localConnection)

		socketConnection = connectToSocket(remoteConnection, "/var/run/docker.sock")

	}

}
