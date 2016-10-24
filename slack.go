package main

import (
	"github.com/nlopes/slack"
	"github.com/pkg/errors"
)

type Slack interface {
	PostMessage(body string) error
}

type SlackClient struct {
	Channel, Token string
	*slack.Client
}

func NewSlackClient(channel, token string) (Slack, error) {
	if len(channel) == 0 {
		return nil, errors.New("missing Slack channel")
	}

	if len(token) == 0 {
		return nil, errors.New("missing Slack token")
	}

	client := slack.New(token)

	return &SlackClient{
		Channel: channel,
		Token: token,
		Client: client,
	}, nil
}

func (c *SlackClient) PostMessage(body string) error {
	_, _, err := c.Client.PostMessage(c.Channel, body, slack.PostMessageParameters{ AsUser: true })
	return err
}
