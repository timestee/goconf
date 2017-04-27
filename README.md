goconf
====
A configuration loader in Go.

## Overview

* Read configuration automatically based on the given struct's field name.
* Load configuration from multiple sources
* multiple file inherit

Values are resolved with the following priorities (lowest to highest):
1. Options struct default value
2. Flags default value
3. Config file value, TOML or JSON file
4. Command line flag





Here is an example of loading configuration in priority:

```go
package main

import "github.com/zsounder/goconf"

type TestOptions struct {
    goconf.AutoOptions
    HTTPAddress string `default:"0.0.0.0:0000"`
    Hosts []string `flag:"hosts" cfg:"hosts" default:"127.0.0.0,127.0.0.1"`
    LogLevel int `default:"3"`
    BoolVar bool `default:"false"`
}

func main() {
   ops := &TestOptions{}
   goconf.ResolveAutoFlag(ops,"conf_1.toml","conf_2.toml").ValidateAndPanic(ops)
}
```

The output will be:

```plain
[Config] auto flag succ, name: _auto_conf_files_ val:
[Config] auto flag succ, name: _auto_dir_running_ val:
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
   "AutoDirRunning": "",
   "HTTPAddress": "127.0.0.1:2",
   "Hosts": [
      "10.0.61.29",
      "10.0.61.30",
      "10.0.61.31",
      "10.0.61.32"
   ],
   "LogLevel": 6,
   "BoolVar": true
}
```

```go
package main

import "github.com/zsounder/goconf"

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
   goconf.ResolveAutoFlag(ops,"conf_3.toml").ValidateAndPanic(ops)
}
```

The output will be:

```plain
[Config] auto flag succ, name: _auto_conf_files_ val:
[Config] auto flag succ, name: _auto_dir_running_ val:
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
   "AutoDirRunning": "",
   "HTTPAddress": "127.0.0.1:2",
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