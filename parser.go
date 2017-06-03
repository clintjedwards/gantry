// Controls argument parsing for gantry program

package main

import "fmt"
import "gopkg.in/alecthomas/kingpin.v2"

func argument_parse() {

    var action = kingpin.Arg("action", "Action to perform on environment.").Required().String()
    var environment = kingpin.Arg("environment", "Environment to perform action in. Environment names are pulled from json file.").Required().String()

    kingpin.Parse()
    fmt.Printf("%s, %s\n", *action, *environment)

}
