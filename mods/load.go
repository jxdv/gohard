package mods

import (
    "encoding/json"
    "io/ioutil"
    "fmt"
    "log"
)

func LoadModules(platform string, admin bool, service string) {
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

    chosenModules := config[service]

    // Testing to see if they are loaded correctly
    for _, module := range chosenModules {
		fmt.Println("Module Name:", module.Name)
		fmt.Println("Description:", module.Description)
		fmt.Println("Command:", module.Command)
		fmt.Println("Require Superuser:", module.RequireSuperuser)
		fmt.Println("Require Restart:", module.RequireRestart)
		fmt.Println("Target OS:", module.TargetOS)
    }

    if platform == "win" {
        // Load all windows modules since we can't check for admin
        fmt.Println("Loading windows modules...")
    } else {
        fmt.Println("Loading linux modules...")

        if !admin {
            // Modules which require admin won't be displayed
        } else {
            // Modules which require admin will be displayed
        }
    }
}