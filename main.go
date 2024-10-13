/*
 * Copyright (c) 2024 flowerinsnow
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package main

import (
    "bufio"
    "fmt"
    "github.com/flowerinsnowdh/recitation-check/config"
    "github.com/pelletier/go-toml/v2"
    "os"
    "os/exec"
    "runtime"
    "strconv"
    "strings"
)

func main() {
    var file *os.File
    var err error

    if file, err = os.Open("config.toml"); err != nil {
        panic(err)
    }

    var conf *config.Config = &config.Config{}

    if err = toml.NewDecoder(file).Decode(conf); err != nil {
        panic(err)
    }

    _ = file.Close()

    for i, row := range conf.Rows {
        fmt.Println(strconv.Itoa(i) + ".", row.Desc)
    }
    var choose int
    if _, err = fmt.Scanf("%d", &choose); err != nil {
        panic(err)
    }

    if choose >= len(conf.Rows) {
        _, _ = fmt.Fprintln(os.Stderr, "no such")
        os.Exit(1)
    }

    var words map[string]string = conf.Rows[choose].Words

    var scanner *bufio.Reader = bufio.NewReader(os.Stdin)

    for k, v := range words {
        fmt.Println(k)
        for {
            if data, _, err := scanner.ReadLine(); err != nil {
                panic(err)
            } else {
                var line string = string(data)
                if strings.EqualFold(line, "ok") {
                    break
                }
                if strings.EqualFold(line, "tip") {
                    fmt.Println(v)
                } else if strings.EqualFold(line, "quit") || strings.EqualFold(line, "exit") {
                    fmt.Println("bye")
                    os.Exit(0)
                } else if strings.EqualFold(line, "?") {
                    if err := openBrowser(strings.ReplaceAll(conf.Dict, "{.word}", k)); err != nil {
                        panic(err)
                    }
                } else {
                    fmt.Println("unknown command")
                }
            }
        }
    }
}

func openBrowser(url string) error {
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
