package mods

import (
    "fmt"
)

func LoadModules(platform string, admin bool, service string) {
    if platform == "win" {
        // Load windows modules
        fmt.Println("Loading windows modules...")
    } else {
        // Load Linux modules
        fmt.Println("Loading linux modules...")
    }
}