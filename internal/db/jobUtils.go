package db

import (
	"context"
	"errors"
	"time"

	"github.com/gaurav-deep01/jobboard-api/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllJobs(cont context.Context) ([]model.Job, error) {
	ctx, cancel := context.WithTimeout(cont, 5*time.Second)
	defer cancel()

	collection := GetCollection(JobCollection)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var jobs []model.Job
	if err := cursor.All(ctx, &jobs); err != nil {
		return nil, err
	}

	return jobs, nil
}

func AddJob(cont context.Context, job model.Job) error {
	ctx, cancel := context.WithTimeout(cont, 5*time.Second)
	defer cancel()

	collection := GetCollection(JobCollection)

	_, err := collection.InsertOne(ctx, job)
	return err
}

func GetJobByID(cont context.Context, id string) (model.Job, error) {
	ctx, cancel := context.WithTimeout(cont, 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Job{}, errors.New("invalid job ID format")
	}

	collection := GetCollection(JobCollection)

	var job model.Job
	if err := collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&job); err != nil {
		return model.Job{}, err
	}

	return job, nil
}

func DeleteJob(cont context.Context, id string) error {
	ctx, cancel := context.WithTimeout(cont, 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid job ID format")
	}

	collection := GetCollection(JobCollection)

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
