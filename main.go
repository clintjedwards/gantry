// Gantry is a tool for manipulating docker-compose infrastructure using ssh tunneling

//  Idea is to create a tool that simply ssh's into a server
//  creates an ssh tunnel for you and then drops you into a shell with the right
//  docker variables set. This would enable you to run normal docker commands on
//  the remote host and wouldn't have to recreate commands

package main

//import "fmt"

func main() {

    arguments := argument_parse()
    go establish_tunnel(arguments.remote_username, arguments.remote_url, arguments.local_port)
    spawn_interactive_shell(arguments)

}
