package goconf

import (
	"errors"
)

var (
	ErrPassinPtr  = errors.New("unsupported type, pass in as ptr")
)