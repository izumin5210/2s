package main

import (
	"bytes"
	"fmt"
)

// Name is application name.
const Name = "2s"

// Version is application version.
const Version = "v0.1.0"

// Revision describes current commit hash generated by `git describe --always`.
var Revision string

// OutputVersion retruns version string.
func OutputVersion() string {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "%s version %s", Name, Version)
	if len(Revision) > 0 {
		fmt.Fprintf(&buf, " (%s)", Revision)
	}
	fmt.Fprint(&buf, "\n")

	return buf.String()
}
