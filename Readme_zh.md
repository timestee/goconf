# goconf  [![Build Status](https://travis-ci.org/timestee/goconf.svg?branch=master)](https://travis-ci.org/zsounder/goconf) [![Go Walker](https://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/zsounder/goconf)  [![GoDoc](https://godoc.org/github.com/zsounder/goconf?status.svg)](https://godoc.org/github.com/zsounder/goconf)


## 简介

* 自动根据配置类的参数名称读取配置文件
* 支持多个配置文件
* 支持文件继承

参数值按照如下顺序解析(优先级由低到高)
1. 配置类参数的默认值
2. 使用flag指定的值
3. 配置文件中的参数值(json、toml等)
4. 命令行传入的参数值

## struct中的tag
```go
type TestOptions struct {
    Hosts []string `flag:"hosts" cfg:"hosts" default:"127.0.0.0,127.0.0.1"`
}
```
* `flag` ：为命令行传入的参数名
* `cfg` 为配置文件中使用的参数名.
* `default` 为参数的默认值

如果没有定义flag,falg等于参数名的snake case化.

如果没有定义cfg,cfg等于flag.

例如： 下例中, 参数HTTPAddress的flag为http_address, cfg也是http_address.
```go
  HTTPAddress string
```

## 使用

### 多文件加载

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

输出为:

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
   "LogLevel": 1,
   "BoolVar": true
}
```

### 文件继承

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

输出为:

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
