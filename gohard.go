package main

import (
    "flag"
    "fmt"
    "os"

    "github.com/jxd1337/gohard/util"
    "github.com/jxd1337/gohard/mods"
    "github.com/jxd1337/gohard/ui"
)

func main() {
    allowedServices := []string{"firewall", "user", "network", "sys", "kernel", "filesystem", "ssh"}

    // Parse cli args
    servicePtr := flag.String("service", "", "Specify a service (firewall, user, network, sys, kernel, filesystem, ssh)")
    flag.Parse()

    // Check if service flag is provided
    if *servicePtr == "" {
        fmt.Println("Please provide the service flag.")
        flag.Usage()
        os.Exit(1)
    }

    // Check if provided service is in the list of allowed services
    validService := false
    for _, s := range allowedServices {
        if *servicePtr == s {
            validService = true
            break
        }
    }

    if !validService {
        fmt.Println("Invalid service. Allowed services are:", allowedServices)
        os.Exit(1)
    }

    // Chosen service for hardening
    service := *servicePtr

    assetExists, err := util.PathExists("assets/modules.json")
    util.FatalErr(err)

    if !assetExists {
        fmt.Println("Unable to find modules.json file in default location!")
        os.Exit(1)
    }

    detectedPlatform, err := util.IsLinux()
    util.FatalErr(err)

    isAdmin, err := util.IsAdmin()
    util.FatalErr(err)

    filteredModules := mods.LoadModules(detectedPlatform, isAdmin, service)
    if len(filteredModules) == 0 {
        /*
        If no modules are loaded after initial loading & filtering
        we exit the program -> there are situations when this can happen, for example:
        loading ssh modules without admin privileges
        */
        fmt.Println("No modules loaded! Exiting..")
        os.Exit(1)
    }

    ui.Run()
}