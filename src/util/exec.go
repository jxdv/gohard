package util

import (
    "fmt"
    "os"
    "strings"
    "runtime"
    "os/exec"
)

func ExecCmd(cmd string) {
    // Replace $USER with the actual user env variable
    if runtime.GOOS == "linux" {
        userEnv := os.Getenv("USER")
        cmd = strings.ReplaceAll(cmd, "$USER", userEnv)
    }

    var cmdInstance *exec.Cmd

    if runtime.GOOS == "linux" {
        cmdInstance = exec.Command("/bin/sh", "-c", cmd)
    } else {
        cmdInstance = exec.Command("cmd", "/C", cmd)
    }

    cmdInstance.Stdout = os.Stdout
    cmdInstance.Stderr = os.Stderr

    err := cmdInstance.Run()
    if err != nil {
        fmt.Printf("Error: %v while executing %s\n", err, cmd)
    }
}