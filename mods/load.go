package mods

import (
    "encoding/json"
    "io/ioutil"
    "fmt"

    "github.com/jxd1337/gohard/util"
)

func LoadModules(platform string, admin bool, service string) []Module {
    if !admin {
        fmt.Println("gohard running without admin privileges - modules which require admin won't be loaded")
    }

    jsonData, err := ioutil.ReadFile("assets/modules.json")
    util.FatalErr(err)


    // Store unmarshalled data
    var config map[string][]Module

    // Load the json data into our struct
    err = json.Unmarshal(jsonData, &config)
    util.FatalErr(err)

    // Empty list which will hold filtered modules
    var chosenModules []Module

    // Filter modules based on OS && admin privileges
    for _, module := range config[service] {
        if platform == "linux" && module.TargetOS == "linux" && (!module.RequireSuperuser || admin) {
            chosenModules = append(chosenModules, module)
        } else if platform == "win" && module.TargetOS == "win" {
            chosenModules = append(chosenModules, module)
        }
    }

    return chosenModules
}