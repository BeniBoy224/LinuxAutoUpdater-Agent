/*
 - Get latest versions of packages that are installed on the server
 - Check critical services are working (sshd, nginx, docker etc)
 - Ability to specify applications running on the server anc check they are working (specifc containers, http server (https://example.com/api))
 - Disk space checks (alert if reaches thresholds)
 - CPU Usage/Ram Usage
 - Know if a reboot is required and alert server correctly
 - Check agent's own logs to see if anything is going wrong

*/

package health

import (
	"agent/internal/logger"
	"fmt"
	"os"
)

func config() {
	// Check if config file exists in /etc/linux-auto-updater/agent.json
	readAgentConfig()
}

func readAgentConfig() {
	agentConfigJsonPath := "/etc/linux-auto-updater/agent.json"
	// Read config file and parse JSON

	jsonData, error := os.ReadFile(agentConfigJsonPath)

	if error != nil {
		// Handle error
		logger.Log("Error reading agent config file")
	}

	logger.Log("Successfully read agent config file")

	fmt.Printf("jsonData: %v\n", jsonData)
}
