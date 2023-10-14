// +build windows

package util

import (
    "golang.org/x/sys/windows"
)

func IsAdmin() (admin bool, err error) {
    var sid *windows.SID

    initErr := windows.AllocateAndInitializeSid(
            &windows.SECURITY_NT_AUTHORITY,
            2,
            windows.SECURITY_BUILTIN_DOMAIN_RID,
            windows.DOMAIN_ALIAS_RID_ADMINS,
            0, 0, 0, 0, 0, 0,
            &sid)
    if initErr != nil {
        return false, initErr
    }
    defer windows.FreeSid(sid)

    token := windows.Token(0)

    member, isMemberErr := token.IsMember(sid)
    if isMemberErr != nil {
        return false, isMemberErr
    }
    return member, nil
}