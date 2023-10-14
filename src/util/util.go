package util

import (
    "runtime"
    "errors"
    "log"
    "os"
    "os/user"

    "golang.org/x/sys/windows"
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

func IsAdmin() (admin bool, err error) {
    currentPlatform, err := IsLinux()
    if err != nil {
        return false, err
    }

    if currentPlatform == "linux" {

        currentUser, err := user.Current()
        if err != nil {
            return false, err
        }

        return currentUser.Username == "root", nil
    } else if currentPlatform == "win" {
        var sid *windows.SID

        err := windows.AllocateAndInitializeSid(
                &windows.SECURITY_NT_AUTHORITY,
                2,
                windows.SECURITY_BUILTIN_DOMAIN_RID,
                windows.DOMAIN_ALIAS_RID_ADMINS,
                0, 0, 0, 0, 0, 0,
                &sid)
        if err != nil {
            FatalErr(err)
        }
        defer windows.FreeSid(sid)

        token := windows.Token(0)

        member, err := token.IsMember(sid)
        if err != nil {
            FatalErr(err)
        }
        return member, nil
    }
    return false, nil
}