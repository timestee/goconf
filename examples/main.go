package main

import "github.com/timestee/goconf"

type TestOptions struct {
	goconf.AutoOptions
	HTTPAddress string   `default:"0.0.0.0:0000"`
	Hosts       []string `flag:"hosts" cfg:"hosts" default:"127.0.0.0,127.0.0.1"`
	LogLevel    int      `default:"3"`
	BoolVar     bool     `default:"false"`
}

func main() {
	ops := &TestOptions{}
	// conf_3 inherit from conf_1 and conf_2
	goconf.MustResolve(ops, "conf_1.toml", "conf_1.json", "conf_3.toml")
}
