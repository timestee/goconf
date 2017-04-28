package goconf

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

type AutoOptions struct {
	AutoConfFiles  string `flag:"_auto_conf_files_"`
	AutoDirRunning string `flag:"_auto_dir_running_"`
}

type Config struct {
	FS  *flag.FlagSet
	FL  *FileLoader
	ops interface{}
}

func (c *Config) GenTemplate(opts interface{}, fname string) error {
	var tomap map[string]interface{} = make(map[string]interface{})
	innserResolve(opts, nil, nil, tomap, false)
	return genTemplate(tomap, fname)
}

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

func (c *Config) ValidateAndPanic(opts interface{}) *Config {
	fmt.Println("[Config]")
	b, _ := json.MarshalIndent(opts, "", "   ")
	fmt.Println(string(b))
	return c
}

// -- GlobalConfig
func NewConfig(name string, errorHandling flag.ErrorHandling) *Config {
	return &Config{
		FS: flag.NewFlagSet(name, errorHandling),
		FL: NewFileLoader(),
	}
}

var GlobalConfig = NewConfig(os.Args[0], flag.ExitOnError)

func GenTemplate(opts interface{}, fname string) error {
	return GlobalConfig.GenTemplate(opts, fname)
}

func Resolve(opts interface{}, files ...string) *Config {
	return GlobalConfig.Resolve(opts, files, false)
}

func ResolveAutoFlag(opts interface{}, files ...string) *Config {
	return GlobalConfig.Resolve(opts, files, true)
}

func ValidateAndPanic(ops interface{}) *Config {
	return GlobalConfig.ValidateAndPanic(ops)
}
