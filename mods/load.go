package mods

import (
    "fmt"
)

func LoadModules(platform string, admin bool, service string) {
    if platform == "win" {
        // Load all windows modules since we can't check for admin
        fmt.Println("Loading windows modules...")
    } else {
        fmt.Println("Loading linux modules...")

        if !admin {
            fmt.Println("Won't load modules that require admin")
        } else {
            fmt.Println("Will read modules that require admin")
        }
    }
}