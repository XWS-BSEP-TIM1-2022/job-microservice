package persistance

import (
	"go.mongodb.org/mongo-driver/mongo"
	"job-microservice/model"
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

/*
func (store *JobMongoDBStore) Get(ctx context.Context, id primitive.ObjectID) (user *model.User, err error) {
	span := tracer.StartSpanFromContext(ctx, "Get")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	filter := bson.M{"_id": id}
	return store.filterOne(ctx, filter)
}

func (store *JobMongoDBStore) GetByEmail(ctx context.Context, email string) (user *model.User, err error) {
	span := tracer.StartSpanFromContext(ctx, "GetByEmail")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	filter := bson.M{"email": email}
	return store.filterOne(ctx, filter)
}

func (store *JobMongoDBStore) GetByUsername(ctx context.Context, username string) (user *model.User, err error) {
	span := tracer.StartSpanFromContext(ctx, "GetByUsername")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	filter := bson.M{"username": username}
	return store.filterOne(ctx, filter)
}

func (store *JobMongoDBStore) GetAll(ctx context.Context) ([]*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "GetAll")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	filter := bson.D{{}}
	return store.filter(ctx, filter)
}

func (store *JobMongoDBStore) Create(ctx context.Context, user *model.User) (*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "Create")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	result, err := store.users.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.Id = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (store *JobMongoDBStore) Update(ctx context.Context, userId primitive.ObjectID, user *model.User) (*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "Update")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	updatedUser := bson.M{
		"$set": user,
	}
	filter := bson.M{"_id": userId}
	_, err := store.users.UpdateOne(ctx, filter, updatedUser)

	if err != nil {
		return nil, err
	}
	user.Id = userId
	return user, nil
}

func (store *JobMongoDBStore) Delete(ctx context.Context, id primitive.ObjectID) error {
	span := tracer.StartSpanFromContext(ctx, "Delete")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	filter := bson.M{"_id": id}
	_, err := store.users.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (store *JobMongoDBStore) DeleteAll(ctx context.Context) {
	span := tracer.StartSpanFromContext(ctx, "DeleteAll")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	store.users.DeleteMany(ctx, bson.D{{}})
}

func (store *JobMongoDBStore) filter(ctx context.Context, filter interface{}) ([]*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "filter")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	cursor, err := store.users.Find(ctx, filter)
	defer cursor.Close(ctx)

	if err != nil {
		return nil, err
	}
	return decode(ctx, cursor)
}

func (store *JobMongoDBStore) filterOne(ctx context.Context, filter interface{}) (product *model.User, err error) {
	span := tracer.StartSpanFromContext(ctx, "filterOne")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	result := store.users.FindOne(ctx, filter)
	err = result.Decode(&product)
	return
}

func decode(ctx context.Context, cursor *mongo.Cursor) (users []*model.User, err error) {
	span := tracer.StartSpanFromContext(ctx, "decode")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	for cursor.Next(ctx) {
		var user model.User
		err = cursor.Decode(&user)
		if err != nil {
			return
		}
		users = append(users, &user)
	}
	err = cursor.Err()
	return
}

func decodeExperience(ctx context.Context, cursor *mongo.Cursor) (experiences []*model.Experience, err error) {
	span := tracer.StartSpanFromContext(ctx, "decode")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	for cursor.Next(ctx) {
		var experience model.Experience
		err = cursor.Decode(&experience)
		if err != nil {
			return
		}
		experiences = append(experiences, &experience)
	}
	err = cursor.Err()
	return
}

func (store *JobMongoDBStore) GetAllWithoutAdmins(ctx context.Context) ([]*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "GetAll")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	filter := bson.M{"role": model.USER}
	return store.filter(ctx, filter)
}

func (store *JobMongoDBStore) GetExperiencesByUserId(ctx context.Context, id string) ([]*model.Experience, error) {
	span := tracer.StartSpanFromContext(ctx, "GetExperiencesByUserId")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	filter := bson.M{"userid": id}
	cursor, err := store.experiences.Find(ctx, filter)
	defer cursor.Close(ctx)

	if err != nil {
		return nil, err
	}
	return decodeExperience(ctx, cursor)
}

func (store *JobMongoDBStore) CreateExperience(ctx context.Context, experience *model.Experience) (*model.Experience, error) {
	span := tracer.StartSpanFromContext(ctx, "CreateNewExperience")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	result, err := store.experiences.InsertOne(ctx, experience)
	if err != nil {
		return nil, err
	}
	experience.Id = result.InsertedID.(primitive.ObjectID)
	return experience, nil
}

func (store *JobMongoDBStore) UpdateExperience(ctx context.Context, experienceId primitive.ObjectID, experience *model.Experience) (*model.Experience, error) {
	//TODO implement me
	panic("implement me")
}

func (store *JobMongoDBStore) DeleteExperience(ctx context.Context, id primitive.ObjectID) error {
	span := tracer.StartSpanFromContext(ctx, "DeleteExperience")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	filter := bson.M{"_id": id}
	_, err := store.experiences.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
*/
