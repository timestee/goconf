package goconf

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

// you could set _auto_conf_files_ to your app's config files,split by with command line flag
type AutoOptions struct {
	AutoConfFiles  string `flag:"_auto_conf_files_"`
	AutoDirRunning string `flag:"_auto_dir_running_"`
}

// Config represents a configuration loader
type Config struct {
	FS  *flag.FlagSet
	FL  *FileLoader
}

// Gen template conf file base on the given struct and save the conf to file.
func (c *Config) GenTemplate(opts interface{}, fname string) error {
	var tomap map[string]interface{} = make(map[string]interface{})
	innserResolve(opts, nil, nil, tomap, false)
	return genTemplate(tomap, fname)
}

// read configuration automatically based on the given struct's field name,
// load configs from struct field's default value, muitiple files and cmdline flags.
func (c *Config) Resolve(opts interface{}, files []string, autlflag bool) *Config {
	// auto flag with default value
	if autlflag {
		innserResolve(opts, c.FS, nil, nil, true)
	}

	// parse cmd args
	c.FS.Parse(os.Args[1:])

	flagInst := c.FS.Lookup("_auto_conf_files_")
	tmp := strings.Trim(flagInst.Value.String(), " ")
	if tmp != "" {
		filesFlag := strings.Split(tmp, ",")
		if len(filesFlag) != 0 {
			files = filesFlag
		}
	}

	fmt.Printf("[Config] file: %v\n", files)
	if len(files) > 0 {
		if err := c.FL.Load(files); err != nil {
			fmt.Printf("[Config] !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! %v\n", err)
		}
	}

	innserResolve(opts, c.FS, c.FL.Data(), nil, false)
	return c
}

// validate the configs and dump it as json string
func (c *Config) ValidateAndPanic(opts interface{}) *Config {
	fmt.Println("[Config]")
	b, _ := json.MarshalIndent(opts, "", "   ")
	fmt.Println(string(b))
	return c
}

func NewConfig(name string, errorHandling flag.ErrorHandling) *Config {
	return &Config{
		FS: flag.NewFlagSet(name, errorHandling),
		FL: NewFileLoader(),
	}
}

// GlobalConfig
var GlobalConfig = NewConfig(os.Args[0], flag.ExitOnError)

// Gen template conf file base on the given struct and save the conf to file.
func GenTemplate(opts interface{}, fname string) error {
	return GlobalConfig.GenTemplate(opts, fname)
}

// read configuration automatically based on the given struct's field name,
// load configs from struct field's default value, muitiple files and cmdline flags.
func Resolve(opts interface{}, files ...string) *Config {
	return GlobalConfig.Resolve(opts, files, false)
}
// auto flag base on given struct's field name
func ResolveAutoFlag(opts interface{}, files ...string) *Config {
	return GlobalConfig.Resolve(opts, files, true)
}

// validate the configs and dump it as json string
func ValidateAndPanic(ops interface{}) *Config {
	return GlobalConfig.ValidateAndPanic(ops)
}

