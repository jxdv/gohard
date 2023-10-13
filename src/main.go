package main

import (
    "fmt"
    "os"

    "github.com/jxd1337/gohard/src/util"
    "github.com/jxd1337/gohard/src/mods"
    "github.com/jxd1337/gohard/src/ui"
)

func main() {
    service := ui.ParseArgs()

    // Check for supported platform
    detectedPlatform, err := util.IsLinux()
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }

    // Check for admin privileges
    isAdmin := util.IsAdmin()

    modules := mods.LoadModules(detectedPlatform, isAdmin, service)
    if len(modules) == 0 {
        /*
        If no modules are loaded after initial loading & filtering
        we exit the program -> there are situations when this can happen, for example:
        loading ssh modules without admin privileges
        */
        fmt.Println("No modules loaded! Exiting..")
        os.Exit(1)
    }

    ui.Run(modules)
}