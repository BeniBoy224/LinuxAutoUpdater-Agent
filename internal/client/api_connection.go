package client

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
)

func ApiConnectionGet(endpoint string) (*http.Response, error) {
	url := "http://192.168.1.48:3333/" + endpoint

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		slog.Error("Failed to create request", "error", err)
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Error connecting to server", "error", err)
		return nil, err
	}

	return resp, nil
}

func ApiConnectionPost(endpoint string, body any) (*http.Response, error) {
	url := "http://192.168.1.48:3333/" + endpoint

	jsonBody, err := json.Marshal(body)
	if err != nil {
		slog.Error("Failed to marshal JSON", "error", err)
		return nil, err
	}

	resp, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		slog.Error("Error connecting to server", "error", err)
		return nil, err
	}

	return resp, nil
}

func ApiConnectionPatch(endpoint string, body any) (*http.Response, error) {
	url := "http://192.168.1.48:3333/" + endpoint

	jsonBody, err := json.Marshal(body)
	if err != nil {
		slog.Error("Failed to marshal JSON", "error", err)
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPatch,
		url,
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		slog.Error("Failed to create request", "error", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Error connecting to server", "error", err)
		return nil, err
	}

	return resp, nil
}
