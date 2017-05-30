package goconf

import (
	"github.com/zsounder/golib/assert"
	"testing"
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
	err := ResolveAutoFlag(ops)
	assert.NotNil(t, err, "ResolveAutoFlag With Struct Get")
}

func Test_Pass_Ptr(t *testing.T) {
	ops := &TestOptions{}
	err := ResolveAutoFlag(ops)
	assert.Nil(t, err, "ResolveAutoFlag With Ptr Get")
}

func Test_Default(t *testing.T) {
	ops := &TestOptions{}
	err := ResolveAutoFlag(ops)

	assert.Nil(t, err,"ResolveAutoFlag")
	assert.Equal(t, ops.HTTPAddress, "0.0.0.0:0000", "HTTPAddress")
	assert.Equal(t, len(ops.Hosts), 2, "Hosts Length")
	assert.Equal(t, ops.Hosts[0], "127.0.0.0", "Hosts[0]")
	assert.Equal(t, ops.Hosts[1], "127.0.0.1", "Hosts[1]")
	assert.Equal(t, ops.LogLevel, 3, "LogLevel")
}

func Test_Default_conf_1_json(t *testing.T) {
	ops := &TestOptions{}
	err := ResolveAutoFlag(ops, "./examples/conf_1.json")

	assert.Nil(t, err, "ResolveAutoFlag")
	assert.Equal(t, ops.HTTPAddress, "127.0.0.1:1", "HTTPAddress")
	assert.Equal(t, len(ops.Hosts), 3, "Hosts Length")
	assert.Equal(t, ops.Hosts[0], "10.0.61.29", "Hosts[0]")
	assert.Equal(t, ops.Hosts[1], "10.0.61.30", "Hosts[1]")
	assert.Equal(t, ops.Hosts[2], "10.0.61.31", "Hosts[2]")
	assert.Equal(t, ops.LogLevel, 3, "LogLevel")
}

func Test_Default_conf_1_toml(t *testing.T) {
	ops := &TestOptions{}
	err := ResolveAutoFlag(ops, "./examples/conf_1.toml")

	assert.Nil(t, err, "ResolveAutoFlag")
	assert.Equal(t, ops.HTTPAddress, "127.0.0.1:1", "HTTPAddress")
	assert.Equal(t, len(ops.Hosts), 3, "Hosts Length")
	assert.Equal(t, ops.Hosts[0], "10.0.61.29", "Hosts[0]")
	assert.Equal(t, ops.Hosts[1], "10.0.61.30", "Hosts[1]")
	assert.Equal(t, ops.Hosts[2], "10.0.61.31", "Hosts[2]")
	assert.Equal(t, ops.LogLevel, 3, "LogLevel")
}

func Test_Default_conf_1_toml_conf_2_toml(t *testing.T) {
	ops := &TestOptions{}
	err := ResolveAutoFlag(ops, "./examples/conf_1.toml", "./examples/conf_2.toml")

	assert.Nil(t, err, "ResolveAutoFlag")
	assert.Equal(t, ops.HTTPAddress, "127.0.0.1:2", "HTTPAddress")
	assert.Equal(t, len(ops.Hosts), 4, "Hosts Length")
	assert.Equal(t, ops.Hosts[0], "10.0.61.29", "Hosts[0]")
	assert.Equal(t, ops.Hosts[1], "10.0.61.30", "Hosts[1]")
	assert.Equal(t, ops.Hosts[2], "10.0.61.31", "Hosts[2]")
	assert.Equal(t, ops.Hosts[3], "10.0.61.32", "Hosts[3]")
	assert.Equal(t, ops.LogLevel, 6, "LogLevel")
}

func Test_Default_conf_3_toml_inherit(t *testing.T) {
	ops := &TestOptions{}
	err := ResolveAutoFlag(ops, "./examples/conf_3.toml")

	assert.Nil(t, err, "ResolveAutoFlag")
	assert.Equal(t, ops.HTTPAddress, "127.0.0.1:2", "HTTPAddress")
	assert.Equal(t, len(ops.Hosts), 5, "Hosts Length")
	assert.Equal(t, ops.Hosts[0], "10.0.61.29", "Hosts[0]")
	assert.Equal(t, ops.Hosts[1], "10.0.61.30", "Hosts[1]")
	assert.Equal(t, ops.Hosts[2], "10.0.61.31", "Hosts[2]")
	assert.Equal(t, ops.Hosts[3], "10.0.61.32", "Hosts[3]")
	assert.Equal(t, ops.Hosts[4], "10.0.61.33", "Hosts[4]")
	assert.Equal(t, ops.LogLevel, 2, "LogLevel")
	assert.Equal(t, ops.BoolVar, true, "BoolVar")
}
