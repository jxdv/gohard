package mods

type Module struct {
    name string
    description string
    command string
    superuser bool
    restart bool
    os string
}