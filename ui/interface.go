package ui

import (
    "strings"
    "strconv"
    "bufio"
    "fmt"
    "os"

    "github.com/jxd1337/gohard/mods"
)

func Run(modules []mods.Module) {
    fmt.Println("Type 'exit' to stop gohard")

    separator := strings.Repeat("-", 80)
    for idx, module := range modules {
        fmt.Println(separator)
        fmt.Println("Module ID -> ", idx + 1)
        fmt.Println("Hardening module -> ", module.Name)
        fmt.Println("Description -> ", module.Description)
        fmt.Println("Command -> ", module.Command)
    }

    reader := bufio.NewReader(os.Stdin)
    fmt.Printf("Enter module ID or interval: ")

    userInput, _ := reader.ReadString('\n')
    userInput = strings.TrimSpace(userInput)

    if userInput == "" {
        // User didn't enter anything
        fmt.Println("No modules selected..")
        return
    } else if userInput == "exit" {
        fmt.Println("Exiting gohard..")
        os.Exit(0)
    }

    selectedModules := parseModuleSelection(userInput, modules)
    if len(selectedModules) == 0 {
        fmt.Println("Invalid module selection..")
        return
    }

    fmt.Println("Executing selected modules:")
    for _, module := range selectedModules {
        fmt.Println(module.Name, "->", module.Description)
        // exec function from another pkg will get called here
    }
}

func parseModuleSelection(input string, modules []Module) []Module {
    var selectedModules []Module

    if input == "-" {
        return modules
    }
}