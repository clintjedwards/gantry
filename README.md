# Gantry: Manage docker through ssh

Gantry is a command line tool to manage docker using the local docker engine over ssh. It sets up an ssh tunnel to the docker instance specified and then drops you into an interactive shell where you can run docker commands locally.

### Prerequisites

* Golang
* OSX

## Installation

```
go get -u github.com/clintjedwards/gantry
```

## Usage
```
~|⇒ gantry --help
usage: gantry [<flags>] <server_url> [<port>]

Flags:
  --help  Show context-sensitive help (also try --help-long and --help-man).

Args:
  <server_url>  Server to connect to in format: user@example.com
  [<port>]      Local port to bind to. Will bind to 9876 if not specified


```

## Example
![example](https://i.imgur.com/7sIl4Pz.gif)

## Troubleshooting
* Make sure the ssh key you're using is loaded via `ssh-add` (ssh agent)
* Make sure the user has permissions to log into the remote machine and talk to the remote docker socket

## Problem | Solution
Looking for an easy way to manage remote docker instances I quickly boiled down my choices to the recommended docker machine. While feature filled the docker machine tool had some short comings I couldn't reconcile.

1) It didn't have great support for docker instances that were already setup.(or at least the support was very confusing)
2) It required exposing the docker process to the outside and possibly the generation of client certs. Wayyyyyy too much work.

Since my primary way of managing remote instances was ssh, it would be easy to just create a ssh tunnel and talk to the docker socket in that manner. Instead of setting up the tunnel manually each time I figured it was easier to wrap it in go and be able to add features as I see fit.

## Authors

* **Clint Edwards** - [Github](https://github.com/clintjedwards)
