package ui

import (
    "fmt"

    "github.com/jxd1337/gohard/mods"
)

func Run(modules []mods.Module) {
    for idx, module := range modules {
        fmt.Println(idx + 1)
        fmt.Println(module.Name)
        fmt.Println(module.Description)
        fmt.Println(module.Command)
    }
}