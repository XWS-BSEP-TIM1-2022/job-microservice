package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Job struct {
	Id              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Position        string             `json:"position"`
	Description     string             `json:"description"`
	DailyActivities []string           `json:"dailyActivities"`
	Prerequisites   []string           `json:"prerequisites"`
	CompanyName     string             `json:"companyName"`
	CompanyLocation string             `json:"companyLocation"`
	OpenDate        time.Time          `json:"openDate"`
	Deleted         bool               `json:"deleted"`
}
