package application

import (
	"job-microservice/model"
	"job-microservice/startup/config"
)

type JobService struct {
	store  model.JobStore
	config *config.Config
}

func NewJobService(store model.JobStore, config *config.Config) *JobService {
	return &JobService{
		store:  store,
		config: config,
	}
}
