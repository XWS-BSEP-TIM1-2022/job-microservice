package config

import (
	"os"
)

type Config struct {
	Port           string
	JobDBHost      string
	JobDBPort      string
	JobServiceName string
}

func NewConfig() *Config {
	return &Config{
		Port:           getEnv("JOB_SERVICE_PORT", "8088"),
		JobDBHost:      getEnv("JOB_DB_HOST", "dislinkt:WiYf6BvFmSpJS2Ob@xws.cjx50.mongodb.net/jobsDB"),
		JobDBPort:      getEnv("JOB_DB_PORT", ""),
		JobServiceName: getEnv("JOB_SERVICE_NAME", "job_service"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
