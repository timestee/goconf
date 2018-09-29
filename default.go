package goconf

import (
	"os"
)

var dc = New(os.Args[0])

func GenTemplate(opts interface{}, fname string) error {
	return dc.GenTemplate(opts, fname)
}

func Resolve(opts interface{}, files ...string) error {
	return dc.Resolve(opts, files)
}

func MustResolve(opts interface{}, files ...string) {
	dc.MustResolve(opts, files)
}
