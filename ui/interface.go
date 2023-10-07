package ui

import (
    "fmt"
    "strings"

    "github.com/jxd1337/gohard/mods"
)

func Run(modules []mods.Module) {
    separator := strings.Repeat("-", 80)
    for idx, module := range modules {
        fmt.Println(separator)
        fmt.Println("Module ID -> ", idx + 1)
        fmt.Println("Hardening module -> ", module.Name)
        fmt.Println("Description -> ", module.Description)
        fmt.Println("Command -> ", module.Command)
    }

    prompt()
}

func prompt() {
    fmt.Println("Enter module ID or an interval of modules:")

    var choice string
    fmt.Scanln(&choice)

    fmt.Println(choice)
}