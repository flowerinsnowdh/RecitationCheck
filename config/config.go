/*
 * Copyright (c) 2024 flowerinsnow
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package config

type Config struct {
    Dict string      `toml:"dict"`
    Rows []RowObject `toml:"rows"`
}

type RowObject struct {
    Desc  string            `toml:"desc"`
    Words map[string]string `toml:"words"`
}

type StringPair struct {
    Key string
    Val string
}
