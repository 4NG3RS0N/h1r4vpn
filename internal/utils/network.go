package utils

import (
	"io"
	"net/http"
	"strings"
	"time"
)

var client = &http.Client{
	Timeout: 5 * time.Second,
}

func PublicIP() string {
	resp, err := client.Get("https://x.ntvstream.workers.dev/")
	if err != nil {
		return "unknown"
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "unknown"
	}

	return strings.TrimSpace(string(body))
}