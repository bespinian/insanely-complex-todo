package store

import (
	"context"
	"time"

	"github.com/bespinian/ict-todo/backend/tasks/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	database   *mongo.Database
	collection mongo.Collection
}

func NewMongoDatabase(uri string, database string) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	db := client.Database(database)
	return db, nil
}

func NewMongoStore(database *mongo.Database) *MongoStore {
	collection := database.Collection("tasks")
	return &MongoStore{database: database, collection: *collection}
}

func (s *MongoStore) List(ctx context.Context) ([]models.Task, error) {
	cursor, err := s.collection.Find(ctx, bson.D{})
	if err != nil {
		return []models.Task{}, err
	}

	tasks := []models.Task{}
	if err = cursor.All(ctx, &tasks); err != nil {
		return []models.Task{}, err
	}
	defer cursor.Close(ctx)

	return tasks, nil
}

func (s *MongoStore) Add(ctx context.Context, task models.Task) (models.Task, error) {
	if task.Id == "" {
		task.Id = primitive.NewObjectID().String()
	}
	_, err := s.collection.InsertOne(ctx, task)
	return task, err
}

func (s *MongoStore) Get(ctx context.Context, id string) (models.Task, error) {
	var task models.Task
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	if err := s.collection.FindOne(ctx, filter).Decode(&task); err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (s *MongoStore) Update(ctx context.Context, task models.Task) (models.Task, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: task.Id}}
	_, err := s.collection.ReplaceOne(ctx, filter, task)
	return task, err
}

func (s *MongoStore) Delete(ctx context.Context, id string) error {
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	_, err := s.collection.DeleteOne(ctx, filter)
	return err
}
