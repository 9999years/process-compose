package config

import (
	"fmt"
	"github.com/adrg/xdg"
	"github.com/rs/zerolog/log"
	"os"
	"os/user"
	"path/filepath"
)

var (
	Version         = "undefined"
	Commit          = "undefined"
	Date            = "undefined"
	CheckForUpdates = "false"
	License         = "Apache-2.0"

	scFiles = []string{
		"shortcuts.yaml",
		"shortcuts.yml",
	}
)

const (
	pcConfigEnv  = "PROC_COMP_CONFIG"
	LogFileFlags = os.O_CREATE | os.O_APPEND | os.O_WRONLY | os.O_TRUNC
	LogFileMode  = os.FileMode(0600)
)

func GetLogFilePath() string {
	return filepath.Join(os.TempDir(), fmt.Sprintf("process-compose-%s%s.log", mustUser(), mode()))
}

func procCompHome() string {
	if env := os.Getenv(pcConfigEnv); env != "" {
		return env
	}
	xdgPcHome, err := xdg.ConfigFile("process-compose")
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to create configuration directory for process compose")
	}
	return xdgPcHome
}

func GetShortCutsPath() string {
	for _, path := range scFiles {
		scPath := filepath.Join(procCompHome(), path)
		if _, err := os.Stat(scPath); err == nil {
			return scPath
		}
	}
	return ""
}

func mustUser() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to retrieve user info")
	}
	return usr.Username
}

func mode() string {
	if isClient() {
		return "-client"
	}
	return ""
}

func isClient() bool {
	for _, proc := range os.Args {
		if proc == "process" {
			return true
		}
		if proc == "attach" {
			return true
		}
	}
	return false
}
