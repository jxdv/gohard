package util

import (
    "runtime"
    "errors"
    "fmt"
    "os/user"
)

func IsLinux() (platform string, err error) {
    if runtime.GOOS == "windows" {
        return "win", nil
    } else if runtime.GOOS == "linux" {
        return "linux", nil
    } else {
        return "", errors.New("Unsupported platform detected!")
    }
}

func IsAdmin() (admin bool, err error) {
    currentPlatform, err := IsLinux()
    if err != nil {
        return false, err
    }

    if currentPlatform == "linux" {
        fmt.Println("Linux detected.")

        currentUser, err := user.Current()
        if err != nil {
            return false, err
        }
        return currentUser.Username == "root", nil
    } else if currentPlatform == "win" {
        // Can't reliably check for admin on windows unfortunately
        fmt.Println("Windows detected.")
        return false, nil
    }
    return false, nil
}

