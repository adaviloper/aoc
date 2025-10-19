package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Theme struct {
	Info string `json:"info"`
	Warn string `json:"warn"`
	Danger string `json:"danger"`
	Success string `json:"success"`
}

// Config holds user-configurable settings for the CLI.
// It is loaded from $HOME/.config/aoc/config.json on startup.
type Config struct {
	BaseDirectory string `json:"base_directory"`
	TemplateLang  string `json:"template_language"`
	Theme Theme `json:"theme"`
}

// Cfg is the singleton configuration instance populated by Init().
var Cfg Config

// Init loads configuration from disk, creating a default file if missing.
// It also expands a leading ~ in BaseDirectory to the user's HOME.
func Init() {
	homeDir := os.Getenv("HOME")
	configDir := filepath.Join(homeDir, ".config", "aoc")
	configPath := filepath.Join(configDir, "config.json")

	if data, err := os.ReadFile(configPath); err == nil {
		if unmarshalErr := json.Unmarshal(data, &Cfg); unmarshalErr != nil {
			log.Fatal("Error during Unmarshal(): ", unmarshalErr)
		}
	} else {
		cwd, cwdErr := os.Getwd()
		if cwdErr != nil {
			log.Fatal("failed to get current working directory: ", cwdErr)
		}
		Cfg = Config{
			BaseDirectory: filepath.Join(cwd, "advent-of-code"),
			TemplateLang:  "ts",
			Theme: Theme{
				Info: "#f9e2af",
				Warn: "#fab387",
				Danger: "#f38ba8",
				Success: "#a6e3a1",
			},
		}
		if mkErr := os.MkdirAll(configDir, 0o755); mkErr == nil {
			if jsonBytes, mErr := json.MarshalIndent(Cfg, "", "  "); mErr == nil {
				_ = os.WriteFile(configPath, jsonBytes, 0o644)
			}
		}
	}

	if strings.HasPrefix(Cfg.BaseDirectory, "~") {
		Cfg.BaseDirectory = strings.Replace(Cfg.BaseDirectory, "~", homeDir, 1)
	}
}
