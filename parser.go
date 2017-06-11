// Controls command line argument parsing for gantry

package main

import "strings"
import "gopkg.in/alecthomas/kingpin.v2"


type Arguments struct{
    remote_username string
    remote_url string
    local_port string
}

func argument_parse() Arguments{

    var server_url = kingpin.Arg("server_url", "Server to connect to in format: user@example.com").Required().String()
    var port = kingpin.Arg("port", "Local port to bind to. Will bind to 9876 if not specified").Default("9876").String()

    kingpin.Parse()

    arguments := Arguments{}

    server_url_slice := strings.Split(*server_url, "@")

    arguments.remote_username = server_url_slice[0]
    arguments.remote_url = server_url_slice[1]
    arguments.local_port = *port

    return arguments
}
