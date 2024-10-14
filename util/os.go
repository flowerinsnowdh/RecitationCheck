package util

import (
    "fmt"
    "os"
    "os/exec"
    "runtime"
)

func OpenBrowser(url string) error {
    var cmd string
    var args []string

    switch runtime.GOOS {
    case "windows":
        cmd = "rundll32"
        args = []string{"url.dll,FileProtocolHandler", url}
    case "linux":
        cmd = "xdg-open"
        args = []string{url}
    default:
        var err error
        _, err = fmt.Fprintln(os.Stderr, "unsupported platform")
        return err
    }

    return exec.Command(cmd, args...).Start()
}
