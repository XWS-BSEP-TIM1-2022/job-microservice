package api

import (
	"job-microservice/application"
)

type JobHandler struct {
	//userService.UnimplementedUserServiceServer
	service *application.JobService
}

func NewJobHandler(
	service *application.JobService) *JobHandler {
	return &JobHandler{
		service: service,
	}
}
