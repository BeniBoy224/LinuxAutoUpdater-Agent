/*
- Function to get system info
- Function to get running software/apps
- Function to get package managers
- Function to get current installed packages && versions of packages if requested
- Fuction to execute commands (preform updates, get general info)
*/
package system

import (
	"log/slog"
	"os"
	"runtime"
)

func GetOS() {
	if conditions := runtime.GOOS; conditions == "linux" {
		return
	} else {
		slog.Error("Non Linux OS detected - Agent does not currently support this OS")
		// Stop the agent from running as unsupported OS
		os.Exit(1)
	}
}
