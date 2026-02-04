/*
THIS is the basics of the application (connecting to API, sending logs to server, auth)

 - Keep up to date with the server
 - Getting the latest packages (new versions and logging that they need updating)
 -
*/

package client

import (
	"log/slog"
	"net/http"
	"net/url"
)

func Client() {
	// Check if config file exists in /etc/linux-auto-updater/agent.json
	resp, err := http.PostForm("http://192.168.1.48:3333/agents/setup",
		url.Values{"desciption": {"my agent"}, "platform": {"linux"}})

	if err != nil {
		// Handle error
		slog.Error("Error connecting to server")
	}

	if resp.StatusCode != http.StatusOK {
		slog.Error("Non-OK HTTP status:", "status", resp.StatusCode, "statusText", resp.Body)
	}

	defer resp.Body.Close()

	return
}
