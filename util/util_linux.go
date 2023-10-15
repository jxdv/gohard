// +build linux

package util

import (
    "os/user"
)

func IsAdmin() (admin bool, err error) {
    currentUser, err := user.Current()
    if err != nil {
        return false, err
    }

    return currentUser.Username == "root", nil
}