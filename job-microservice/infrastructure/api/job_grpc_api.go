package api

import (
	"context"
	jobService "github.com/XWS-BSEP-TIM1-2022/dislinkt/util/proto/job"
	"github.com/XWS-BSEP-TIM1-2022/dislinkt/util/tracer"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"job-microservice/application"
)

type JobHandler struct {
	jobService.UnimplementedJobServiceServer
	service *application.JobService
}

func NewJobHandler(
	service *application.JobService) *JobHandler {
	return &JobHandler{
		service: service,
	}
}

func (handler *JobHandler) GetRequest(ctx context.Context, in *jobService.JobIdRequest) (*jobService.GetResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	id := in.JobId
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	job, err := handler.service.Get(ctx, objectId)
	if err != nil {
		return nil, err
	}
	jobPb := mapJob(job)
	response := &jobService.GetResponse{
		Job: jobPb,
	}
	return response, nil
}

func (handler *JobHandler) GetAllRequest(ctx context.Context, in *jobService.EmptyRequest) (*jobService.JobsResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	jobs, err := handler.service.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	response := &jobService.JobsResponse{
		Jobs: []*jobService.Job{},
	}
	for _, job := range jobs {
		current := mapJob(job)
		response.Jobs = append(response.Jobs, current)
	}
	return response, nil
}

func (handler *JobHandler) PostRequest(ctx context.Context, in *jobService.UserRequest) (*jobService.GetResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "PostRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	jobFromRequest := mapJobPb(in.Job)
	job, err := handler.service.Create(ctx, jobFromRequest)
	if err != nil {
		return nil, err
	}
	jobPb := mapJob(job)
	response := &jobService.GetResponse{
		Job: jobPb,
	}
	return response, nil
}

func (handler *JobHandler) DeleteRequest(ctx context.Context, in *jobService.JobIdRequest) (*jobService.EmptyRequest, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "DeleteRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	id, err := primitive.ObjectIDFromHex(in.JobId)
	if err != nil {
		return nil, err
	}
	err = handler.service.Delete(ctx, id)
	if err != nil {
		return nil, err
	}
	response := &jobService.EmptyRequest{}
	return response, nil
}

func (handler *JobHandler) SearchJobsRequest(ctx context.Context, in *jobService.SearchRequest) (*jobService.JobsResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "SearchJobsRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	jobs, err := handler.service.Search(ctx, in.SearchParam)
	if err != nil {
		return nil, err
	}
	response := &jobService.JobsResponse{
		Jobs: []*jobService.Job{},
	}
	for _, job := range jobs {
		current := mapJob(job)
		response.Jobs = append(response.Jobs, current)
	}
	return response, nil
}
