// Provides functions to handle ssh tunneling

package main

import "fmt"
import "net"
import "os"
import "io"

import "golang.org/x/crypto/ssh"
import "golang.org/x/crypto/ssh/agent"

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
		listener.Close()
		panic(err)
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
		fmt.Printf("Failed to dial: %s", err)
		panic("Exiting")
	}

	return remoteConnection

}

//Establish socket connection on remote machine
func connectToRemoteSocket(remoteConnection *ssh.Client) *net.Conn {

	socketConnection, err := remoteConnection.Dial("unix", "/var/run/docker.sock")
	if err != nil {
		fmt.Printf("Failed to dial: %s", err)
		panic("Exiting")
	}

	return &socketConnection
}

//Allow two way communication between two connections
func copyConnectionData(writer, reader net.Conn) {
	_, err := io.Copy(writer, reader)
	if err != nil {
		fmt.Printf("io.Copy error: %s", err)
	}
}

//Establish local, remote, socket connections and then allow communication
func establishTunnel(username string, hostname string, localPort string) {
	listener := establishLocalListener(localPort)

	for {
		localConnection, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		remoteConnection := connectToRemote(username, hostname)
		socketConnection := connectToRemoteSocket(remoteConnection)

		go copyConnectionData(localConnection, *socketConnection)
		go copyConnectionData(*socketConnection, localConnection)

	}

}
