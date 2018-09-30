package goconf

import (
	"testing"

	"github.com/zsounder/golib/assert"
)

type TestOptions struct {
	AutoOptions
	HTTPAddress string   `default:"0.0.0.0:0000"`
	Hosts       []string `flag:"hosts" cfg:"hosts" default:"127.0.0.0,127.0.0.1"`
	LogLevel    int      `default:"3"`
	BoolVar     bool     `default:"false"`
}

func Test_Pass_Struct(t *testing.T) {
	ops := TestOptions{}
	err := Resolve(ops)
	assert.NotNil(t, err)
}

func Test_Pass_Ptr(t *testing.T) {
	ops := &TestOptions{}
	err := Resolve(ops)
	assert.Nil(t, err)
}

func Test_Default(t *testing.T) {
	ops := &TestOptions{}
	err := Resolve(ops)

	assert.Nil(t, err)
	assert.Equal(t, ops.HTTPAddress, "0.0.0.0:0000")
	assert.Equal(t, len(ops.Hosts), 2)
	assert.Equal(t, ops.Hosts[0], "127.0.0.0")
	assert.Equal(t, ops.Hosts[1], "127.0.0.1")
	assert.Equal(t, ops.LogLevel, 3)
}

func Test_Default_conf_1_json(t *testing.T) {
	ops := &TestOptions{}
	err := Resolve(ops, "./examples/conf_1.json")

	assert.Nil(t, err)
	assert.Equal(t, ops.HTTPAddress, "127.0.0.1:1")
	assert.Equal(t, len(ops.Hosts), 3)
	assert.Equal(t, ops.Hosts[0], "10.0.61.29")
	assert.Equal(t, ops.Hosts[1], "10.0.61.30")
	assert.Equal(t, ops.Hosts[2], "10.0.61.31")
	assert.Equal(t, ops.LogLevel, 3)
}

func Test_Default_conf_1_toml(t *testing.T) {
	ops := &TestOptions{}
	err := Resolve(ops, "./examples/conf_1.toml")

	assert.Nil(t, err)
	assert.Equal(t, ops.HTTPAddress, "127.0.0.1:1")
	assert.Equal(t, len(ops.Hosts), 3)
	assert.Equal(t, ops.Hosts[0], "10.0.61.29")
	assert.Equal(t, ops.Hosts[1], "10.0.61.30")
	assert.Equal(t, ops.Hosts[2], "10.0.61.31")
	assert.Equal(t, ops.LogLevel, 3)
}

func Test_Default_conf_1_toml_conf_2_toml(t *testing.T) {
	ops := &TestOptions{}
	err := Resolve(ops, "./examples/conf_1.toml", "./examples/conf_2.toml")

	assert.Nil(t, err)
	assert.Equal(t, ops.HTTPAddress, "127.0.0.1:2")
	assert.Equal(t, len(ops.Hosts), 4)
	assert.Equal(t, ops.Hosts[0], "10.0.61.29")
	assert.Equal(t, ops.Hosts[1], "10.0.61.30")
	assert.Equal(t, ops.Hosts[2], "10.0.61.31")
	assert.Equal(t, ops.Hosts[3], "10.0.61.32")
	assert.Equal(t, ops.LogLevel, 6)
}

func Test_Default_conf_3_toml_inherit(t *testing.T) {
	ops := &TestOptions{}
	err := Resolve(ops, "./examples/conf_3.toml")

	assert.Nil(t, err)
	assert.Equal(t, ops.HTTPAddress, "127.0.0.1:2")
	assert.Equal(t, len(ops.Hosts), 5)
	assert.Equal(t, ops.Hosts[0], "10.0.61.29")
	assert.Equal(t, ops.Hosts[1], "10.0.61.30")
	assert.Equal(t, ops.Hosts[2], "10.0.61.31")
	assert.Equal(t, ops.Hosts[3], "10.0.61.32")
	assert.Equal(t, ops.Hosts[4], "10.0.61.33")
	assert.Equal(t, ops.LogLevel, 2)
	assert.Equal(t, ops.BoolVar, true)
}
