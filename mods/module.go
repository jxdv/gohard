package mods

type Module struct {
    name String
    description String
    command String
    superuser bool
    restart bool
    os String
}