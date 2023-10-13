package ui

import (
    "flag"
    "fmt"
    "os"
)

func ParseArgs() string {
    allowedServices := []string{"firewall", "user", "network", "sys", "kernel", "filesystem", "ssh"}

    // Parse CLI args
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
        fmt.Printf("Invalid service. Use one of these: %v\n", allowedServices)
        os.Exit(1)
    }

    // Chosen service for hardening
    service := *servicePtr

    return service
}