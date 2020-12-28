package services

import (
	"context"
	"time"

	c "github.com/thalissonfelipe/ugly-api/config"
	"github.com/thalissonfelipe/ugly-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// MovieService struct
type MovieService struct {
	Client *mongo.Client
}

// GetMovies should find all movies in the mongo database
func (m *MovieService) GetMovies() (*[]models.Movie, error) {
	collection := m.Client.Database(c.MyConfig.DB.DatabaseName).Collection("movies")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	movies := make([]models.Movie, 0)
	err = cur.All(ctx, &movies)

	return &movies, err
}

// GetMovie should find a movie by name in the mongo database
func (m *MovieService) GetMovie(name string) (*models.Movie, error) {
	collection := m.Client.Database(c.MyConfig.DB.DatabaseName).Collection("movies")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var movie models.Movie
	err := collection.FindOne(ctx, bson.M{"name": name}).Decode(&movie)
	return &movie, err
}

// CreateMovie should insert a new movie in the mongo database
func (m *MovieService) CreateMovie(movie *models.Movie) error {
	collection := m.Client.Database(c.MyConfig.DB.DatabaseName).Collection("movies")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, movie)
	return err
}

// UpdateMovie should update a movie by name in the mongo database
func (m *MovieService) UpdateMovie(name string, movie *models.Movie) error {
	// collection := m.Client.Database(c.MyConfig.DB.DatabaseName).Collection("movies")
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// update := bson.D{{"$set", movie}}
	// result := collection.FindOneAndUpdate(ctx, bson.M{"name": name}, update)
	// err := result.Err()
	// log.Println(result)

	// results := collection.FindOne(ctx, bson.M{"name": name})
	// log.Println(results.Err())
	// return err
	return nil
}

// DeleteMovie should delete a movie by name in the mongo database
func (m *MovieService) DeleteMovie(name string) error {
	collection := m.Client.Database(c.MyConfig.DB.DatabaseName).Collection("movies")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := collection.FindOneAndDelete(ctx, bson.M{"name": name}).Err()
	return err
}
