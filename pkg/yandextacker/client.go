package yandextacker

import (
	"net/http"

	"codeberg.org/kalsin/UtelBot/pkg/configs"
)

type Connection struct {
	Client  *http.Client
	Headers map[string]string
	Host    string
	Timeout uint8
	Retries uint8
	Verify  bool
}

func NewConnection(config *configs.YandexTrackerConfig) *Connection {
	return &Connection{
		Client: &http.Client{},
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "OAuth " + config.Token,
			"X-Org-Id":      config.OrgID,
		},
		Host:    "https://api.tracker.yandex.net/v2",
		Timeout: 10,
		Retries: 10,
		Verify:  true,
	}
}
