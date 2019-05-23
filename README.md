# goconf 
[![Build Status](https://travis-ci.org/timestee/goconf.svg?branch=master)](https://travis-ci.org/timestee/goconf)
[![Go Walker](https://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/timestee/goconf)
[![GoDoc](https://godoc.org/github.com/timestee/goconf?status.svg)](https://godoc.org/github.com/timestee/goconf)
[![Go Report Card](https://goreportcard.com/badge/github.com/timestee/goconf)](https://goreportcard.com/report/github.com/timestee/goconf)

## Overview

* Read configuration automatically based on the given struct's field name.
* Load configuration from multiple sources
* file inherit

Values are resolved with the following priorities (lowest to highest):
1. Options struct default value
2. Flags default value
3. Config file value, TOML or JSON file
4. Command line flag

## About field tags in structs
```go
type TestOptions struct {
    Hosts []string `flag:"hosts" cfg:"hosts" default:"127.0.0.0,127.0.0.1"`
}
```
* `flag` is the name passed from the command line.
* `cfg` is the name used in config files.
* `default` is the default value

If do not define `flag` tag, `flag` will be snake case of the fild name.

If do not define `cfg` tag, `cfg` value will be `flag` value.

For example, flag and cfg will be http_address.
```go
  HTTPAddress string
```

## Usage

### load multiple config files

```go
package main

import "github.com/timestee/goconf"

type TestOptions struct {
    goconf.AutoOptions
    HTTPAddress string `default:"0.0.0.0:0000"`
    Hosts []string `flag:"hosts" cfg:"hosts" default:"127.0.0.0,127.0.0.1"`
    LogLevel int `default:"3"`
    BoolVar bool `default:"false"`
}

func main() {
   ops := &TestOptions{}
   goconf.MustResolve(ops,"conf_1.toml","conf_2.toml")
}
```

`go run main.go --log_level=1`

The output will be:

```plain
[Config] auto flag succ, name: _auto_conf_files_ val:
[Config] auto flag succ, name: http_address val: 0.0.0.0:0000
[Config] auto flag fail, name: hosts val: 127.0.0.0,127.0.0.1 err: type not support []string
[Config] auto flag succ, name: log_level val: 3
[Config] auto flag succ, name: bool_var val: false
[Config] file: [conf_1.toml conf_2.toml]
[Config] load: conf_1.toml
[Config] load: conf_2.toml
[Config]
{
   "AutoConfFiles": "",
   "HTTPAddress": "127.0.0.1:2",
   "Hosts": [
      "10.0.61.29",
      "10.0.61.30",
      "10.0.61.31",
      "10.0.61.32"
   ],
   "LogLevel": 1,
   "BoolVar": true
}
```

### load config file with file inherited

```go
package main

import "github.com/timestee/goconf"

type TestOptions struct {
    goconf.AutoOptions
    HTTPAddress string `default:"0.0.0.0:0000"`
    Hosts []string `flag:"hosts" cfg:"hosts" default:"127.0.0.0,127.0.0.1"`
    LogLevel int `default:"3"`
    BoolVar bool `default:"false"`
}

func main() {
   ops := &TestOptions{}
   // conf_3 inherit from conf_1 and conf_2
   goconf.MustResolve(ops,"conf_3.toml")
}
```
`go run main.go --http_address=0.0.0.0:1111111`

The output will be:

```plain
[Config] auto flag succ, name: _auto_conf_files_ val:
[Config] auto flag succ, name: http_address val: 0.0.0.0:0000
[Config] auto flag fail, name: hosts val: 127.0.0.0,127.0.0.1 err: type not support []string
[Config] auto flag succ, name: log_level val: 3
[Config] auto flag succ, name: bool_var val: false
[Config] file: [conf_3.toml]
[Config] load: ./conf_1.toml
[Config] load: ./conf_2.toml
[Config] load: conf_3.toml
[Config]
{
   "AutoConfFiles": "",
   "HTTPAddress": "0.0.0.0:1111111",
   "Hosts": [
      "10.0.61.29",
      "10.0.61.30",
      "10.0.61.31",
      "10.0.61.32",
      "10.0.61.33"
   ],
   "LogLevel": 2,
   "BoolVar": true
}
```
