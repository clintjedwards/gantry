//Provides functions to create local port,
//                        ssh into remote server,
//                        connect to docker socket and
//                        spawn terminal that has the proper docker env variables to talk to daemon

package main

import "fmt"
import "net"
import "os"
import "io"

import "golang.org/x/crypto/ssh"
import "golang.org/x/crypto/ssh/agent"

func SSHAgent() ssh.AuthMethod {

    if sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		return ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers)
	}
	return nil
}

func establish_local_listner(port string) net.Listener{

    listener, err := net.Listen("tcp", "localhost:" + port)
    if err != nil {
        listener.Close()
        panic(err)
    }

    return listener
}

func connect_to_remote(username string, hostname string) *ssh.Client{

    sshConfig := &ssh.ClientConfig{
    	User: username,
    	Auth: []ssh.AuthMethod{SSHAgent()},
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    remote_connection, err := ssh.Dial("tcp", hostname + ":" + "22", sshConfig)
    if err != nil {
    	fmt.Printf("Failed to dial: %s", err)
        panic("Exiting")
    }

    return remote_connection

}

func connect_to_remote_socket(remote_connection *ssh.Client) *net.Conn{

    socket_connection, err := remote_connection.Dial("unix", "/var/run/docker.sock")
    if err != nil {
    	fmt.Printf("Failed to dial: %s", err)
        panic("Exiting")
    }

    return &socket_connection
}

func copy_connection_data(writer, reader net.Conn){
    _, err:= io.Copy(writer, reader)
    if err != nil {
        fmt.Printf("io.Copy error: %s", err)
    }
}

func establish_tunnel(username string, hostname string, local_port string) {
    listener := establish_local_listner(local_port)

    for {
		local_connection, err := listener.Accept()
		if err != nil {
			panic(err)
		}

        remote_connection := connect_to_remote(username, hostname)
        socket_connection := connect_to_remote_socket(remote_connection)


    	go copy_connection_data(local_connection, *socket_connection)
        go copy_connection_data(*socket_connection, local_connection)

	}

}
