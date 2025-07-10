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

// GetPostgreSQLDSN returns the PostgreSQL connection string
func GetPostgreSQLDSN() string {
	dsn := os.Getenv("XUI_POSTGRES_DSN")
	if dsn != "" {
		return dsn
	}

	// Build DSN from individual environment variables
	host := os.Getenv("XUI_POSTGRES_HOST")
	if host == "" {
		host = "93.126.18.78"
	}

	port := os.Getenv("XUI_POSTGRES_PORT")
	if port == "" {
		port = "5432"
	}

	user := os.Getenv("XUI_POSTGRES_USER")
	if user == "" {
		user = "postgres"
	}

	password := os.Getenv("XUI_POSTGRES_PASSWORD")
	if password == "" {
		password = "ebitfeweb6xcij3o"
	}

	dbname := os.Getenv("XUI_POSTGRES_DB")
	if dbname == "" {
		dbname = "xui"
	}

	sslmode := os.Getenv("XUI_POSTGRES_SSLMODE")
	if sslmode == "" {
		sslmode = "disable"
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)
}

func GetLogFolder() string {
	logFolderPath := os.Getenv("XUI_LOG_FOLDER")
	if logFolderPath == "" {
		logFolderPath = "/var/log"
	}
	return logFolderPath
}
