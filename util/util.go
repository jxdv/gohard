package util

import (
    "runtime"
    "errors"
    "log"
    "os"
)

func FatalErr (err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func PathExists(assetPath string) (exists bool, err error) {
    _, err = os.Stat(assetPath)
    if err == nil {
        return true, nil
    }

    if os.IsNotExist(err) {
        return false, nil
    }
    return false, err
}

func IsLinux() (platform string, err error) {
    if runtime.GOOS == "windows" {
        return "win", nil
    } else if runtime.GOOS == "linux" {
        return "linux", nil
    } else {
        return "", errors.New("Unsupported platform detected!")
    }
}