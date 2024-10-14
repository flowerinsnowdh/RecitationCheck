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
    "github.com/flowerinsnowdh/recitation-check/util"
    "github.com/pelletier/go-toml/v2"
    "os"
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

    var words []config.StringPair = make([]config.StringPair, 0)

    for k, v := range conf.Rows[choose].Words {
        words = append(words, config.StringPair{
            Key: k,
            Val: v,
        })
    }

    util.FisherYates(words) // 打乱顺序

    var scanner *bufio.Reader = bufio.NewReader(os.Stdin)

    for _, word := range words {
        fmt.Println(word.Key)
        for {
            if data, _, err := scanner.ReadLine(); err != nil {
                panic(err)
            } else {
                var line string = string(data)
                if strings.EqualFold(line, "ok") {
                    break
                }
                if strings.EqualFold(line, "tip") {
                    fmt.Println(word.Val)
                } else if strings.EqualFold(line, "quit") || strings.EqualFold(line, "exit") {
                    fmt.Println("bye")
                    os.Exit(0)
                } else if strings.EqualFold(line, "?") {
                    if err := util.OpenBrowser(strings.ReplaceAll(conf.Dict, "{.word}", word.Key)); err != nil {
                        panic(err)
                    }
                } else {
                    fmt.Println("unknown command")
                }
            }
        }
    }
}


