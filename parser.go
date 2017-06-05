// Controls argument parsing for gantry program

package main

import "strings"
import "gopkg.in/alecthomas/kingpin.v2"

func argument_parse() []string{

    var server_url = kingpin.Arg("server_url", "Server to connect to in format: user@example.com").Required().String()

    kingpin.Parse()

    server_url_slice := strings.Split(*server_url, "@")

    return server_url_slice
}
