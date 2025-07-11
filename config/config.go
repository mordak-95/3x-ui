package config

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed version
var version string

//go:embed name
var name string

type LogLevel string

const (
	Debug  LogLevel = "debug"
	Info   LogLevel = "info"
	Notice LogLevel = "notice"
	Warn   LogLevel = "warn"
	Error  LogLevel = "error"
)

func GetVersion() string {
	return strings.TrimSpace(version)
}

func GetName() string {
	return strings.TrimSpace(name)
}

func GetLogLevel() LogLevel {
	if IsDebug() {
		return Debug
	}
	logLevel := os.Getenv("XUI_LOG_LEVEL")
	if logLevel == "" {
		return Info
	}
	return LogLevel(logLevel)
}

func IsDebug() bool {
	return os.Getenv("XUI_DEBUG") == "true"
}

func GetBinFolderPath() string {
	binFolderPath := os.Getenv("XUI_BIN_FOLDER")
	if binFolderPath == "" {
		binFolderPath = "bin"
	}
	return binFolderPath
}

func GetDBFolderPath() string {
	dbFolderPath := os.Getenv("XUI_DB_FOLDER")
	if dbFolderPath == "" {
		dbFolderPath = "/etc/x-ui"
	}
	return dbFolderPath
}

func GetDBPath() string {
	return fmt.Sprintf("%s/%s.db", GetDBFolderPath(), GetName())
}

func GetLogFolder() string {
	logFolderPath := os.Getenv("XUI_LOG_FOLDER")
	if logFolderPath == "" {
		logFolderPath = "/var/log"
	}
	return logFolderPath
}

func GetPostgresHost() string {
	// return os.Getenv("XUI_PG_HOST")
	return "dokploy.roshanbin.com"
}

func GetPostgresUser() string {
	// return os.Getenv("XUI_PG_USER")
	return "postgres"
}

func GetPostgresPassword() string {
	// return os.Getenv("XUI_PG_PASSWORD")
	return "dijpveapofaezb12"
}

func GetPostgresDBName() string {
	// return os.Getenv("XUI_PG_DBNAME")
	return "xui"
}

func GetPostgresPort() string {
	// port := os.Getenv("XUI_PG_PORT")
	// if port == "" {
	// 	return "5432"
	// }
	return "5432"
}
