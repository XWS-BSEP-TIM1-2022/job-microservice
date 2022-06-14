package application

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"job-microservice/model"
	"job-microservice/startup/config"
	"strings"
)

type JobService struct {
	store  model.JobStore
	config *config.Config
}

var Log = logrus.New()

func NewJobService(store model.JobStore, config *config.Config) *JobService {
	return &JobService{
		store:  store,
		config: config,
	}
}

func (service *JobService) Get(ctx context.Context, id primitive.ObjectID) (*model.Job, error) {
	Log.Info("Get job by id: " + id.Hex())
	return service.store.Get(ctx, id)
}

func (service *JobService) GetAll(ctx context.Context) ([]*model.Job, error) {
	Log.Info("Get all jobs")
	return service.store.GetAll(ctx)
}

func (service *JobService) Create(ctx context.Context, job *model.Job) (*model.Job, error) {
	Log.Info("Create job")
	return service.store.Create(ctx, job)
}

func (service *JobService) Delete(ctx context.Context, id primitive.ObjectID) error {
	Log.Info("Delete job by id: " + id.Hex())
	return service.store.Delete(ctx, id)
}

func (service *JobService) Search(ctx context.Context, searchParam string) ([]*model.Job, error) {
	Log.Info("Search job by searchParam: " + searchParam)

	jobs, err := service.GetAll(ctx)
	if err != nil {
		Log.Error("Getting all jobs in Search method by searchParam: " + searchParam)
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
	Log.Info("List contains, param: " + param)
	
	for _, item := range list {
		if strings.Contains(strings.ToLower(item), param) {
			return true
		}
	}
	return false
}
