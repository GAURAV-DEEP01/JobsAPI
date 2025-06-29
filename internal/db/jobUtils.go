package db

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Job struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title    string             `bson:"title" json:"title"`
	Company  string             `bson:"company" json:"company"`
	Location string             `bson:"location" json:"location"`
}

var (
	GetAllJobs = func() ([]Job, error) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		collection := GetCollection(JobCollection)

		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)

		var jobs []Job
		if err := cursor.All(ctx, &jobs); err != nil {
			return nil, err
		}

		return jobs, nil
	}

	AddJob = func(job Job) error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		collection := GetCollection(JobCollection)

		_, err := collection.InsertOne(ctx, job)
		return err
	}

	GetJobByID = func(id string) (Job, error) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return Job{}, errors.New("invalid job ID format")
		}

		collection := GetCollection(JobCollection)

		var job Job
		if err := collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&job); err != nil {
			return Job{}, err
		}

		return job, nil
	}

	DeleteJob = func(id string) error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return errors.New("invalid job ID format")
		}

		collection := GetCollection(JobCollection)

		_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
		return err
	}
)
