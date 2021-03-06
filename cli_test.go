package main

import (
	"bytes"
	"fmt"
	"strings"

	"testing"
)

func createTestCli() (*CLI, *bytes.Buffer, *bytes.Buffer, *bytes.Buffer) {
	inStream, outStream, errStream := new(bytes.Buffer), new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{inStream: inStream, outStream: outStream, errStream: errStream}
	return cli, inStream, outStream, errStream
}

func TestRun(t *testing.T) {
	// TODO: Not yet implemented.
}

func TestRun_versionFlag(t *testing.T) {
	cli, _, outStream, _ := createTestCli()
	command := "2s --version"
	args := strings.Split(command, " ")

	if actual, expected := cli.Run(args), ExitCodeOK; actual != expected {
		t.Fatalf("%q exited with %q, expected %q", command, actual, expected)
	}

	expected := fmt.Sprintf("2s version %s", Version)
	actual := outStream.String()
	if !strings.Contains(actual, expected) {
		t.Fatalf("%q outputted %q, expected %q", command, actual, expected)
	}
}
