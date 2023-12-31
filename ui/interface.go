package ui

import (
    "strings"
    "strconv"
    "bufio"
    "fmt"
    "os"
    "path/filepath"

    "github.com/jxdv/gohard/mods"
    "github.com/jxdv/gohard/util"
)

func Run(modules []mods.Module) {
    assetPath := filepath.Join("assets", "modules.json")

    // Check for .json file, which contains hardening modules
    assetExists, err := util.PathExists(assetPath)
    util.FatalErr(err)

    if !assetExists {
        fmt.Println("Unable to find modules.json file in default location!")
        os.Exit(1)
    }

    separator := strings.Repeat("-", 50)
    for idx, module := range modules {
        if len(modules) > 1 {
            fmt.Println(separator)
        }
        fmt.Printf("Module ID: %d | %s\n", idx + 1, module.Name)
        fmt.Printf("Description: %s\n", module.Description)
    }

    reader := bufio.NewReader(os.Stdin)
    fmt.Printf("Enter module ID or interval ('exit' to terminate gohard): ")

    userInput, _ := reader.ReadString('\n')
    userInput = strings.TrimSpace(userInput)

    if strings.ToLower(userInput) == "exit" {
        fmt.Println("Bye!")
        os.Exit(0)
    }

    // User didn't enter anything
    if userInput == "" {
        fmt.Println("No modules selected.")
        return
    }

    selectedModules := parseModuleSelection(userInput, modules)
    if len(selectedModules) == 0 {
        fmt.Println("Invalid module selection.")
        return
    }

    for _, module := range selectedModules {
        fmt.Printf("Executing %s\n", module.Name)
        util.ExecCmd(module.Command)
    }

    fmt.Println("Done.")
}

func parseModuleSelection(input string, modules []mods.Module) []mods.Module {
    var selectedModules []mods.Module

    // This will select all the available modules
    if input == "-" {
        return modules
    }

    // Check if it's a range selection
    if strings.Contains(input, "-") {
        rangeParts := strings.Split(input, "-")
        if len (rangeParts) == 2 {
            start, startErr := strconv.Atoi(rangeParts[0])
            end, endErr := strconv.Atoi(rangeParts[1])

            // Valid range selection
            if startErr == nil && endErr == nil && start <= end && start >= 1 && end <= len(modules) {
                for i := start; i <= end; i++ {
                    selectedModules = append(selectedModules, modules[i-1])
                }
            } else {
                fmt.Println("Range is invalid!")
                os.Exit(1)
            }
        }
    } else {
        moduleID, err := strconv.Atoi(input)
        if err == nil && moduleID >= 1 && moduleID <= len(modules) {
            selectedModules = append(selectedModules, modules[moduleID-1])
        } else {
            fmt.Println("Wrong module ID selected!")
            os.Exit(1)
        }
    }

    return selectedModules
}
