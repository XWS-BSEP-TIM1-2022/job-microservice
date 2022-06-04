package api

import (
	jobService "github.com/XWS-BSEP-TIM1-2022/dislinkt/util/proto/job"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"job-microservice/model"
	"strconv"
	"strings"
	"time"
)

func mapJob(job *model.Job) *jobService.Job {
	jobPb := &jobService.Job{
		Id:              job.Id.Hex(),
		UserId:          job.UserID,
		Position:        job.Position,
		Description:     job.Description,
		DailyActivities: job.DailyActivities,
		Prerequisites:   job.Prerequisites,
		CompanyName:     job.CompanyName,
		CompanyLocation: job.CompanyLocation,
		OpenDate:        job.OpenDate.String(),
	}
	return jobPb
}
func mapJobPb(jobPb *jobService.Job) *model.Job {
	id, _ := primitive.ObjectIDFromHex(jobPb.Id)
	t := time.Now()
	if jobPb.OpenDate != "" {
		dateString := strings.Split(jobPb.OpenDate, "T")
		date := strings.Split(dateString[0], "-")
		year, _ := strconv.Atoi(date[0])
		month, _ := strconv.Atoi(date[1])
		day, _ := strconv.Atoi(date[2])

		timeString := strings.Split(dateString[1], ":")
		hour, _ := strconv.Atoi(timeString[0])
		minutes, _ := strconv.Atoi(timeString[1])
		t = time.Date(year, time.Month(month), day, hour, minutes, 0, 0, time.UTC)
	}
	job := &model.Job{
		Id:              id,
		UserID:          jobPb.UserId,
		Position:        jobPb.Position,
		Description:     jobPb.Description,
		DailyActivities: jobPb.DailyActivities,
		Prerequisites:   jobPb.Prerequisites,
		CompanyName:     jobPb.CompanyName,
		CompanyLocation: jobPb.CompanyLocation,
		OpenDate:        t,
	}
	return job
}
