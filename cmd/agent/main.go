package main

import (
	"agent/internal/system"
	"log/slog"
	"os"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"

	"agent/internal/config"
)

func main() {
	Init()
	startUp()
}

func Init() {
	startTime := time.Now().Format("2026-01-01_13-01-01")
	// set env to PROD in production
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	var handler slog.Handler

	if env == "prod" {
		// ✅ production: rolling log file
		handler = slog.NewJSONHandler(
			&lumberjack.Logger{
				Filename:   "logs/linux-autoupdater-agent-" + startTime + ".log",
				MaxSize:    100, // MB
				MaxBackups: 10,
				MaxAge:     14, // days
				Compress:   true,
			},
			&slog.HandlerOptions{
				Level: slog.LevelInfo,
			},
		)
	} else {
		// 🧑‍💻 dev: stdout, readable
		handler = slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			},
		)
	}

	slog.SetDefault(slog.New(handler))
}

func startUp() {
	// On start
	slog.Info("Starting Linux Auto Updater Agent")

	system.GetOS()

	config.ConfigStartup()
}
