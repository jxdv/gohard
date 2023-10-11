package mods

type Module struct {
    Name string `json:"name"`
    Description string `json:"desc"`
    Command string `json:"command"`
    RequireSuperuser bool `json:"require_superuser"`
    RequireRestart bool `json:"require_restart"`
    TargetOS string `json:"target_os"`
}

type Section struct {
    Modules []Module `json:""`
}

type Configuration struct {
	Firewall   Section `json:"firewall"`
	User       Section `json:"user"`
	Network    Section `json:"network"`
	Sys        Section `json:"sys"`
	Kernel     Section `json:"kernel"`
	Filesystem Section `json:"filesystem"`
	SSH        Section `json:"ssh"`
}