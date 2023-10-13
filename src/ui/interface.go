package ui

import (
    "strings"
    "strconv"
    "bufio"
    "fmt"
    "os"

    "github.com/jxd1337/gohard/src/mods"
    "github.com/jxd1337/gohard/src/util"
)

func Run(modules []mods.Module) {
    // Check for .json file, which contains hardening modules
    assetExists, err := util.PathExists("assets/modules.json")
    util.FatalErr(err)

    if !assetExists {
        fmt.Println("Unable to find modules.json file in default location!")
        os.Exit(1)
    }

    fmt.Println("Type 'exit' to stop gohard")

    separator := strings.Repeat("-", 80)
    for idx, module := range modules {
        fmt.Println(separator)
        fmt.Printf("Module ID -> %d\n", idx + 1)
        fmt.Printf("Hardening module -> %s\n", module.Name)
        fmt.Printf("Description -> %s\n", module.Description)
        fmt.Printf("Command -> %s\n", module.Command)
    }

    reader := bufio.NewReader(os.Stdin)
    fmt.Printf("Enter module ID or interval: ")

    userInput, _ := reader.ReadString('\n')
    userInput = strings.TrimSpace(userInput)

    if strings.ToLower(userInput) == "exit" {
        fmt.Println("Exiting gohard..")
        os.Exit(0)
    }

    if userInput == "" {
        // User didn't enter anything
        fmt.Println("No modules selected..")
        return
    }

    selectedModules := parseModuleSelection(userInput, modules)
    if len(selectedModules) == 0 {
        fmt.Println("Invalid module selection..")
        return
    }

    fmt.Println("Executing selected modules:")
    for _, module := range selectedModules {
        fmt.Println(module.Name, "->", module.Description)
        util.ExecCmd(module.Command)
    }
}

func parseModuleSelection(input string, modules []mods.Module) []mods.Module {
    var selectedModules []mods.Module

    if input == "-" {
        return modules
    }

    // Check if it's a range selection
    if strings.Contains(input, "-") {
        rangeParts := strings.Split(input, "-")
        if len (rangeParts) == 2 {
            start, err1 := strconv.Atoi(rangeParts[0])
            end, err2 := strconv.Atoi(rangeParts[1])

            if err1 == nil && err2 == nil && start <= end && start >= 1 && end <= len(modules) {
                // Valid range selection
                for i := start; i <= end; i++ {
                    selectedModules = append(selectedModules, modules[i-1])
                }
            } else {
                fmt.Println("Something went wrong while parsing range selection!")
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