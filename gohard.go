package main

import (
    "flag"
    "fmt"
    "log"
    "os"

    "github.com/jxd1337/gohard/util"
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

    service := *servicePtr
    fmt.Println(service)

    isAdmin, err := util.IsAdmin()
    if err != nil {
        log.Fatal(err)
    }

    if !isAdmin {
        fmt.Println("gohard running without admin privileges, modules which require admin won't be displayed")
    }
}