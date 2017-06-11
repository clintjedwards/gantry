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
```
gantry|master⚡ ⇒ gantry ubuntu@utility.clintjedwards.com
Connected via tcp://localhost:9876 to utility.clintjedwards.com as user ubuntu
Starting dockerized interactive shell
gantry|master⚡ ⇒ docker info
...
Kernel Version: 4.8.0-53-generic
Operating System: Ubuntu 16.04.2 LTS
OSType: linux
Architecture: x86_64
...
gantry|master⚡ ⇒ docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS                NAMES
f618d7c677f3        gantry_test         "bash"              44 hours ago        Up 44 hours         0.0.0.0:22->22/tcp   cranky_perlman
```

## Problem | Solution
Looking for an easy way to manage remote docker instances I quickly boiled down my choices to the recommended docker machine. While feature filled the docker machine tool had some short comings I couldn't reconcile.

1) It didn't have great support for docker instances that were already setup.(or at least the support was very confusing)
2) It required exposing the docker process to the outside and possibly the generation of client certs. Wayyyyyy too much work.

Since my primary way of managing remote instances was ssh, it would be easy to just create a ssh tunnel and talk to the docker socket in that manner. Instead of setting up the tunnel manually each time I figured it was easier to wrap it in go and be able to add features as I see fit.

## Authors

* **Clint Edwards** - [Github](https://github.com/clintjedwards)
