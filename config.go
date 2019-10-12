package goconf

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
)

var (
	errPassInPtr = fmt.Errorf("unsupported type, pass in as ptr")
	globalLogger = func(s string) { log.Print(s) }
)

//SetGlobalLogger set the global logger func, if nil, lib will keep silent
func SetGlobalLogger(lf func(string)) {
	globalLogger = lf
}

// Config represents a configuration loader
type Config struct {
	flagSet    *flag.FlagSet
	fileLoader *fileLoader
}

// New Config with name and option struct
func New(name string) *Config {
	c := &Config{}

	c.flagSet = flag.NewFlagSet(name, flag.ContinueOnError)
	c.flagSet.SetOutput(ioutil.Discard)
	c.fileLoader = &fileLoader{log: c.log}
	return c
}

func (c *Config) log(msg string) {
	msg = fmt.Sprintf("goconf : %s", msg)
	if globalLogger != nil {
		globalLogger(msg)
	}
}

// GenTemplate Gen template conf file base on the given struct and save the conf to file.
func (c *Config) GenTemplate(opts interface{}, fn string) error {
	tm := make(map[string]interface{})
	innerResolve(opts, nil, nil, tm, false, c.log)
	return genTemplate(tm, fn)
}

// Resolve given files, return error if fail
func (c *Config) Resolve(opts interface{}, files []string) error {
	return c.resolve(opts, files)
}

// MustResolve given files, panic if fail
func (c *Config) MustResolve(opts interface{}, files []string) {
	if err := c.resolve(opts, files); err != nil {
		c.log(fmt.Sprintf("Failed in must model err: %s", err.Error()))
		panic(err)
	}
}

// read configuration automatically based on the given struct's field name,
// load configs from struct field's default value, muitiple files and cmdline flags.
func (c *Config) resolve(opts interface{}, files []string) error {
	if reflect.ValueOf(opts).Kind() != reflect.Ptr {
		return errPassInPtr
	}
	// auto flag with default value
	innerResolve(opts, c.flagSet, nil, nil, true, c.log)
	if err := c.flagSet.Parse(os.Args[1:]); err != nil {
		if err != flag.ErrHelp {
			_, _ = fmt.Fprintf(os.Stderr, "flag: %v\n", err)
			c.flagSet.Usage()
		}
	}

	flagInst := c.flagSet.Lookup("_auto_conf_files_")
	if flagInst != nil {
		tmp := strings.Trim(flagInst.Value.String(), " ")
		if tmp != "" {
			filesFlag := strings.Split(tmp, ",")
			if len(filesFlag) != 0 {
				files = filesFlag
			}
		}
	}

	c.log(fmt.Sprintf("file:  %v", files))
	errs := errArray{ErrorFormat: errArrayDotFormatFunc}
	if len(files) > 0 {
		if err := c.fileLoader.Load(files); err != nil {
			c.log(fmt.Sprintf("Error with %s", err.Error()))
			errs.Push(err)
		}
	}

	innerResolve(opts, c.flagSet, c.fileLoader.Data(), nil, false, c.log)
	if b, err := json.MarshalIndent(opts, "", "   "); err != nil {
		errs.Push(err)
	} else {
		c.log(fmt.Sprintf("Contents:\n%v", string(b)))
	}

	if errs.Err() != nil {
		return errs
	}

	return nil
}
