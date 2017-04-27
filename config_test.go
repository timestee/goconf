package goconf

import (
	"testing"
)

type TestOptions struct {
    AutoOptions
    HTTPAddress string `default:"0.0.0.0:0000"`
    Hosts []string `flag:"hosts" cfg:"hosts" default:"127.0.0.0,127.0.0.1"`
    LogLevel int `default:"3"`
    BoolVar bool `default:"false"`
}


func Test_Default(t *testing.T) {
	ops := &TestOptions{}
	ResolveAutoFlag(ops).ValidateAndPanic(ops)

	if ops.HTTPAddress != "0.0.0.0:0000" {
		t.Errorf("HTTPAddress shoult be: 0.0.0.0:0000 get:%s",ops.HTTPAddress)
	}
	if len(ops.Hosts) != 2 {
		t.Errorf("Hosts length should be 2, get: %s",len(ops.Hosts))
	}
	if ops.Hosts[0] != "127.0.0.0" {
		t.Errorf("Hosts[0] shoult be:127.0.0.0, get: %s",ops.Hosts[0])
	}
	if ops.Hosts[1] != "127.0.0.1" {
		t.Errorf("Hosts[1] shoult be:127.0.0.1, get: %s",ops.Hosts[1])
	}
	if ops.LogLevel != 3 {
		t.Errorf("LogLevel shoult be: 3 get:%s",ops.LogLevel)
	}
}

func Test_Default_conf_1_json(t *testing.T) {
	ops := &TestOptions{}
	ResolveAutoFlag(ops,"./examples/conf_1.json").ValidateAndPanic(ops)
	if ops.HTTPAddress != "127.0.0.1:1" {
		t.Errorf("HTTPAddress shoult be: 127.0.0.1:1 get:%s",ops.HTTPAddress)
	}
	if len(ops.Hosts) != 3 {
		t.Errorf("Hosts length should be 3, get: %s",len(ops.Hosts))
	}
	if ops.Hosts[0] != "10.0.61.29" {
		t.Errorf("Hosts[0] shoult be:10.0.61.29, get: %s",ops.Hosts[0])
	}
	if ops.Hosts[1] != "10.0.61.30" {
		t.Errorf("Hosts[1] shoult be:10.0.61.30, get: %s",ops.Hosts[1])
	}
	if ops.Hosts[2] != "10.0.61.31" {
		t.Errorf("Hosts[2] shoult be:10.0.61.31, get: %s",ops.Hosts[2])
	}
	if ops.LogLevel != 3 {
		t.Errorf("LogLevel shoult be: 3 get:%d",ops.LogLevel)
	}
}

func Test_Default_conf_1_toml(t *testing.T) {
	ops := &TestOptions{}
	ResolveAutoFlag(ops,"./examples/conf_1.toml").ValidateAndPanic(ops)
	if ops.HTTPAddress != "127.0.0.1:1" {
		t.Errorf("HTTPAddress shoult be: 127.0.0.1:1 get:%s",ops.HTTPAddress)
	}
	if len(ops.Hosts) != 3 {
		t.Errorf("Hosts length should be 3, get: %s",len(ops.Hosts))
	}
	if ops.Hosts[0] != "10.0.61.29" {
		t.Errorf("Hosts[0] shoult be:10.0.61.29, get: %s",ops.Hosts[0])
	}
	if ops.Hosts[1] != "10.0.61.30" {
		t.Errorf("Hosts[1] shoult be:10.0.61.30, get: %s",ops.Hosts[1])
	}
	if ops.Hosts[2] != "10.0.61.31" {
		t.Errorf("Hosts[2] shoult be:10.0.61.31, get: %s",ops.Hosts[2])
	}
	if ops.LogLevel != 3 {
		t.Errorf("LogLevel shoult be: 3 get:%d",ops.LogLevel)
	}
}

func Test_Default_conf_1_toml_conf_2_toml(t *testing.T) {
	ops := &TestOptions{}
	ResolveAutoFlag(ops,"./examples/conf_1.toml","./examples/conf_2.toml").ValidateAndPanic(ops)
	if ops.HTTPAddress != "127.0.0.1:2" {
		t.Errorf("HTTPAddress shoult be: 127.0.0.1:2 get:%s",ops.HTTPAddress)
	}
	if len(ops.Hosts) != 4 {
		t.Errorf("Hosts length should be 4, get: %s",len(ops.Hosts))
	}
	if ops.Hosts[0] != "10.0.61.29" {
		t.Errorf("Hosts[0] shoult be:10.0.61.29, get: %s",ops.Hosts[0])
	}
	if ops.Hosts[1] != "10.0.61.30" {
		t.Errorf("Hosts[1] shoult be:10.0.61.30, get: %s",ops.Hosts[1])
	}
	if ops.Hosts[2] != "10.0.61.31" {
		t.Errorf("Hosts[2] shoult be:10.0.61.31, get: %s",ops.Hosts[2])
	}
	if ops.Hosts[3] != "10.0.61.32" {
		t.Errorf("Hosts[3] shoult be:10.0.61.32, get: %s",ops.Hosts[3])
	}
	if ops.LogLevel != 6 {
		t.Errorf("LogLevel shoult be: 6 get:%d",ops.LogLevel)
	}
}

func Test_Default_conf_3_toml_inherit(t *testing.T) {
	ops := &TestOptions{}
	ResolveAutoFlag(ops,"./examples/conf_3.toml",).ValidateAndPanic(ops)
	if ops.HTTPAddress != "127.0.0.1:2" {
		t.Errorf("HTTPAddress shoult be: 127.0.0.1:2 get:%s",ops.HTTPAddress)
	}
	if len(ops.Hosts) != 5 {
		t.Errorf("Hosts length should be 4, get: %s",len(ops.Hosts))
	}
	if ops.Hosts[0] != "10.0.61.29" {
		t.Errorf("Hosts[0] shoult be:10.0.61.29, get: %s",ops.Hosts[0])
	}
	if ops.Hosts[1] != "10.0.61.30" {
		t.Errorf("Hosts[1] shoult be:10.0.61.30, get: %s",ops.Hosts[1])
	}
	if ops.Hosts[2] != "10.0.61.31" {
		t.Errorf("Hosts[2] shoult be:10.0.61.31, get: %s",ops.Hosts[2])
	}
	if ops.Hosts[3] != "10.0.61.32" {
		t.Errorf("Hosts[3] shoult be:10.0.61.32, get: %s",ops.Hosts[3])
	}
	if ops.Hosts[4] != "10.0.61.33" {
		t.Errorf("Hosts[4] shoult be:10.0.61.33, get: %s",ops.Hosts[3])
	}
	if ops.LogLevel != 2 {
		t.Errorf("LogLevel shoult be: 2 get:%s",ops.LogLevel)
	}
	if ops.BoolVar != true {
		t.Errorf("BoolVar shoult be: true get:%v",ops.BoolVar)
	}
}