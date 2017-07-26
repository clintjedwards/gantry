// Controls command line argument parsing for gantry

package main

import (
	"log"
	"os/user"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

// Arguments represents storage object for command-line parsed variables.
type Arguments struct {
	remoteUsername string
	remoteURL      string
	localPort      string
}

func argumentParse() Arguments {

	var serverURL = kingpin.Arg("serverURL", "Server to connect to in format: user@example.com or example.com").Required().String()
	var port = kingpin.Arg("port", "Local port to bind to. Will bind to 9876 if not specified").Default("9876").String()

	kingpin.Parse()

	arguments := Arguments{}

	if strings.Contains(*serverURL, "@") {

		serverURLSlice := strings.Split(*serverURL, "@")

		arguments.remoteUsername = serverURLSlice[0]
		arguments.remoteURL = serverURLSlice[1]

	} else {
		user, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}

		arguments.remoteUsername = user.Username
		arguments.remoteURL = *serverURL
	}

	arguments.localPort = *port

	return arguments
}
