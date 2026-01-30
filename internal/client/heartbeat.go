package client

import (
	"log/slog"
	"net/http"
	"net/url"
	"time"
)

const HeartBeatInterval = 60 // in seconds

func HeartBeat() {
	for {
		// Check if config file exists in /etc/linux-auto-updater/agent.json
		resp, err := http.PostForm("http://192.168.1.48:3333/agents/heartbeat",
			url.Values{"agent_id": {"test"}})

		if err != nil {
			// Handle error
			slog.Info("Error connecting to server")
		}

		defer resp.Body.Close()

		time.Sleep(HeartBeatInterval)
	}
}
