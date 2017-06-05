// Gantry is a tool for manipulating docker-compose infrastructure using ssh

//  Idea is to create a tool that simply ssh's into a server
//  creates an ssh tunnel for you and then drops you into a shell with the right
//  docker variables set. This would enable you to run normal docker commands on
//  the remote host and wouldn't have to recreate commands

package main

import "fmt"

func main() {

    connection_url_slice := argument_parse()
    connect_to_remote(connection_url_slice[0], connection_url_slice[1])

    fmt.Println(connection_url_slice)

}
