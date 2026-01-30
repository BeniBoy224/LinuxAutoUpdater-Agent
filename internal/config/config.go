/*
	 Steps to SUCCESS

	 - Check if agent json file is located in /etc/linux-auto-updater/agent.json

	 - If json config found, send heartbead to controller server

			- if this is the first time the server has seen this hostname and id. create a new UUID on api and send it to agent (along with token)
			- send basic info about server (hardware, software running, versions)
*/
package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
)

type AgentConfig struct {
	AgentID  string `json:"agent_id"`
	Token    string `json:"token"`
	Hostname string `json:"controller_url"`
}

// Location should never be changed
const agentConfigJsonPath = "/etc/linux-auto-updater/agent-config.json"

func ConfigStartup() {
	config, err := ReadAgentConfig()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Agent Config: %+v\n", config)
}

func ReadAgentConfig() (*AgentConfig, error) {
	jsonData, err := os.ReadFile(agentConfigJsonPath)
	if err != nil {
		slog.Error("Failed to read config file: %w", err)
		return nil, errors.New("Failed to read config file: " + err.Error())
	}

	var config AgentConfig
	if err := json.Unmarshal(jsonData, &config); err != nil {
		slog.Error("Error parsing agent config file: %w", err)
		return nil, errors.New("Error parsing agent config file: " + err.Error())
	}

	if config.Hostname == "" {
		slog.Error("Controller URL is missing in config file")
		return nil, errors.New("Controller URL is missing in config file: " + err.Error())
	}

	if config.AgentID == "" {
		// call api to get a new agent id
		// then update agent id and token in config file
	}

	return &config, nil
}

func UpdateAgentConfig(config *AgentConfig) error {
	jsonData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config to JSON: %w", err)
	}

	if err := os.WriteFile(agentConfigJsonPath, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}
	return nil
}
