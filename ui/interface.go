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
}