// Controls configuration file loading for gantry

package main

import "fmt"
import "encoding/json"
import "os"
import "io/ioutil"


type Config struct {
    Name string `json:"name"`
    Server_url string `json:"server_url"`
}

func load_config() map[string]string {

    raw, err := ioutil.ReadFile("./environments.json")
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    var config []Config
    json.Unmarshal(raw, &config)

    config_map := make(map[string]string)

    for _, item := range config {
        config_map[item.Name] = item.Server_url
    }

    return config_map

}
