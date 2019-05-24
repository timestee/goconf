package goconf

import (
	"os"
)

var dc = New(os.Args[0])

// GenTemplate Gen template conf file base on the given struct and save the conf to file.
func GenTemplate(opts interface{}, fname string) error {
	return dc.GenTemplate(opts, fname)
}

// Resolve given files, return error if fail
func Resolve(opts interface{}, files ...string) error {
	return dc.Resolve(opts, files)
}

// MustResolve given files, panic if fail
func MustResolve(opts interface{}, files ...string) {
	dc.MustResolve(opts, files)
}
