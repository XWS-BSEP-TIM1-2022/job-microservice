package persistance

import (
	"context"
	"github.com/XWS-BSEP-TIM1-2022/dislinkt/util/tracer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"job-microservice/model"
	"time"
)

const (
	DATABASE   = "jobsDB"
	COLLECTION = "jobs"
)

type JobMongoDBStore struct {
	jobs *mongo.Collection
}

func NewJobMongoDBStore(client *mongo.Client) model.JobStore {
	jobs := client.Database(DATABASE).Collection(COLLECTION)
	return &JobMongoDBStore{
		jobs: jobs,
	}
}

func (store *JobMongoDBStore) Get(ctx context.Context, id primitive.ObjectID) (*model.Job, error) {
	span := tracer.StartSpanFromContext(ctx, "Get")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	filter := bson.M{"_id": id, "deleted": false}
	return store.filterOne(ctx, filter)
}

func (store *JobMongoDBStore) GetAll(ctx context.Context) ([]*model.Job, error) {
	span := tracer.StartSpanFromContext(ctx, "GetAll")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	filter := bson.M{"deleted": false}
	return store.filter(ctx, filter)
}

func (store *JobMongoDBStore) Create(ctx context.Context, job *model.Job) (*model.Job, error) {
	span := tracer.StartSpanFromContext(ctx, "Create")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	job.Deleted = false
	job.OpenDate = time.Now()

	result, err := store.jobs.InsertOne(ctx, job)
	if err != nil {
		return nil, err
	}
	job.Id = result.InsertedID.(primitive.ObjectID)
	return job, nil
}

func (store *JobMongoDBStore) Update(ctx context.Context, jobId primitive.ObjectID, job *model.Job) (*model.Job, error) {
	span := tracer.StartSpanFromContext(ctx, "Update")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	updatedJob := bson.M{
		"$set": job,
	}
	filter := bson.M{"_id": jobId}
	_, err := store.jobs.UpdateOne(ctx, filter, updatedJob)

	if err != nil {
		return nil, err
	}
	job.Id = jobId
	return job, nil
}

func (store *JobMongoDBStore) Delete(ctx context.Context, id primitive.ObjectID) error {
	span := tracer.StartSpanFromContext(ctx, "Delete")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	job, err := store.Get(ctx, id)
	if err != nil {
		return err
	}

	job.Deleted = true

	job, err = store.Update(ctx, id, job)
	if err != nil {
		return err
	}

	return nil
}

func (store *JobMongoDBStore) filter(ctx context.Context, filter interface{}) ([]*model.Job, error) {
	span := tracer.StartSpanFromContext(ctx, "filter")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	cursor, err := store.jobs.Find(ctx, filter)
	defer cursor.Close(ctx)

	if err != nil {
		return nil, err
	}
	return decode(ctx, cursor)
}

func (store *JobMongoDBStore) filterOne(ctx context.Context, filter interface{}) (job *model.Job, err error) {
	span := tracer.StartSpanFromContext(ctx, "filterOne")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	result := store.jobs.FindOne(ctx, filter)
	err = result.Decode(&job)
	return
}

func decode(ctx context.Context, cursor *mongo.Cursor) (jobs []*model.Job, err error) {
	span := tracer.StartSpanFromContext(ctx, "decode")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	for cursor.Next(ctx) {
		var job model.Job
		err = cursor.Decode(&job)
		if err != nil {
			return
		}
		jobs = append(jobs, &job)
	}
	err = cursor.Err()
	return
}
