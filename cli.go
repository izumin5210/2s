package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	// EnvSlackChannel is environmental variable to set a target channel of Slack.
	EnvSlackChannel = "SLACK_CHANNEL"
	// EnvSlackToken is environmental variable to set Slack WebAPI token
	EnvSlackToken = "SLACK_TOKEN"
)

// Exit codes are values representing an exit code for a error type.
const (
	ExitCodeOK int = 0

	// Errors start at 10
	ExitCodeError = 10 + iota
	ExitCodeParseFlagsError
	ExitCodeChannelNotFound
	ExitCodeTokenNotFound
)

// CLI is the command line interface object.
type CLI struct {
	inStream             io.Reader
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		body	string
		token	string
		channel	string
		version bool
	)

	flags := flag.NewFlagSet(Name, flag.ExitOnError)
	flags.SetOutput(cli.errStream)

	flags.StringVar(&channel, "channel", os.Getenv(EnvSlackChannel), "")
	flags.StringVar(&channel, "c", os.Getenv(EnvSlackChannel), "")

	flags.StringVar(&token, "token", os.Getenv(EnvSlackToken), "")
	flags.StringVar(&token, "t", os.Getenv(EnvSlackToken), "")

	flags.BoolVar(&version, "version", false, "")
	flags.BoolVar(&version, "v", false, "")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagsError
	}

	if version {
		fmt.Fprintf(cli.outStream, OutputVersion())
		return ExitCodeOK
	}

	if len(channel) == 0 {
		fmt.Fprintln(cli.errStream, "Invalid argument: you must set SLACK_CHANNEL.")
		return ExitCodeChannelNotFound
	}

	if len(token) == 0 {
		fmt.Fprintln(cli.errStream, "Invalid argument: you must set SLACK_TOKEN.")
		return ExitCodeTokenNotFound
	}

	parsedArgs := flags.Args()
	if len(parsedArgs) == 0 {
		fmt.Fscanln(cli.inStream, &body)
	} else {
		body = strings.Join(parsedArgs, " ")
	}

	slack, err := NewSlackClient(channel, token)
	if err != nil {
		fmt.Fprintf(cli.errStream, "Failed to create Slack client: %s\n", err)
		return ExitCodeError
	}

	err = slack.PostMessage(body)
	if err != nil {
		fmt.Fprintf(cli.errStream, "Failed to post message to Slack: %s", err)
		return ExitCodeError
	}

	return ExitCodeOK
}
