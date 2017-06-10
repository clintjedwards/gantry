//We set up the proper environment for the user
// then drop into an interactive shell
// author: http://technosophos.com/2014/07/11/start-an-interactive-shell-from-within-go.html

package main

import "os"
import "os/user"
import "fmt"


func spawn_interactive_shell(arguments Arguments) {

    current_user, err := user.Current()
    if err != nil {
        panic(err)
    }

    cwd, err := os.Getwd()
    if err != nil {
        panic(err)
    }

    os.Setenv("DOCKER_HOST", "tcp://localhost:" + arguments.local_port)

    // Transfer stdin, stdout, and stderr to the new process
    // and also set target directory for the shell to start in.
    pa := os.ProcAttr {
        Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
        Dir: cwd,
    }

    fmt.Printf("Connected via tcp://localhost:%s to %s as user %s\n", arguments.local_port, arguments.remote_url, arguments.remote_username)
    //fmt.Println("Connected via tcp://localhost:" + arguments.local_port + " to " + arguments.remote_username + ":" + arguments.remote_url)
    fmt.Println("Starting dockerized interactive shell")

    // Start up a new shell.
    // Note that we supply "login" twice.
    // -fpl means "don't prompt for PW and pass through environment."
    proc, err := os.StartProcess("/usr/bin/login", []string{"login", "-fpl", current_user.Username}, &pa)
    if err != nil {
        panic(err)
    }

    // Wait until user exits the shell
    state, err := proc.Wait()
    if err != nil {
        panic(err)
    }

    // Alert the user that they have exited shell
    fmt.Printf("Exited dockerized shell: %s\n", state.String())
}
