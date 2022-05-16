package application

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"job-microservice/model"
	"job-microservice/startup/config"
	"strings"
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

func (service *JobService) Get(ctx context.Context, id primitive.ObjectID) (*model.Job, error) {
	return service.store.Get(ctx, id)
}

func (service *JobService) GetAll(ctx context.Context) ([]*model.Job, error) {
	return service.store.GetAll(ctx)
}

func (service *JobService) Create(ctx context.Context, job *model.Job) (*model.Job, error) {
	return service.store.Create(ctx, job)
}

func (service *JobService) Delete(ctx context.Context, id primitive.ObjectID) error {
	return service.store.Delete(ctx, id)
}

func (service *JobService) Search(ctx context.Context, searchParam string) ([]*model.Job, error) {
	jobs, err := service.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	searchParam = strings.ToLower(searchParam)

	var retVal []*model.Job
	for _, job := range jobs {
		if strings.Contains(strings.ToLower(job.Description), searchParam) || strings.Contains(strings.ToLower(job.Position), searchParam) || strings.Contains(strings.ToLower(job.CompanyName), searchParam) || strings.Contains(strings.ToLower(job.CompanyLocation), searchParam) || service.listContains(job.DailyActivities, searchParam) || service.listContains(job.Prerequisites, searchParam) {
			retVal = append(retVal, job)
		}
	}
	return retVal, nil
}

func (service *JobService) listContains(list []string, param string) bool {
	for _, item := range list {
		if strings.Contains(strings.ToLower(item), param) {
			return true
		}
	}
	return false
}
