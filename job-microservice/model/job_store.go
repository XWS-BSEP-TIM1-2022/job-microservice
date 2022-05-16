package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobStore interface {
	Get(ctx context.Context, id primitive.ObjectID) (*Job, error)
	GetAll(ctx context.Context) ([]*Job, error)
	Create(ctx context.Context, job *Job) (*Job, error)
	Update(ctx context.Context, id primitive.ObjectID, job *Job) (*Job, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
}
