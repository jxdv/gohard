package mods

import (
    "encoding/json"
    "io/ioutil"
    "log"
)

func LoadModules(platform string, admin bool, service string) []Module {
    jsonData, err := ioutil.ReadFile("assets/modules.json")
    if err != nil {
        log.Fatal(err)
    }

    // This holds unmarshalled data
    var config map[string][]Module

    // Load the json data into our struct
    if err := json.Unmarshal(jsonData, &config); err != nil {
        log.Fatal(err)
    }

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