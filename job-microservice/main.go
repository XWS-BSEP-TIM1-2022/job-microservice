package main

import (
	"job-microservice/startup"
	cfg "job-microservice/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
	defer server.Stop()
}
