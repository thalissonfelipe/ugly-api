package services

import (
	"context"
	"fmt"
	"time"

	"github.com/thalissonfelipe/ugly-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// MService struct
type MService struct {
	Client *mongo.Client
}

// GetMovies should find all movies in the mongo database
func (m *MService) GetMovies() (*[]models.Movie, error) {
	collection := m.Client.Database("ugly_db").Collection("movies")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	var movies []models.Movie
	err = cursor.All(ctx, &movies)

	return &movies, err
}

// GetMovie should find a movie by name in the mongo database
func (m *MService) GetMovie(name string) (*models.Movie, error) {
	collection := m.Client.Database("ugly_db").Collection("movies")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var movie models.Movie
	err := collection.FindOne(ctx, bson.M{"name": name}).Decode(&movie)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, nil
		default:
			return nil, err
		}
	}
	return &movie, err
}

// CreateMovie should insert a new movie in the mongo database
func (m *MService) CreateMovie(movie *models.Movie) error {
	collection := m.Client.Database("ugly_db").Collection("movies")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, movie)
	return err
}

// UpdateMovie should update a movie by name in the mongo database
func (m *MService) UpdateMovie(movie *models.Movie) error {
	collection := m.Client.Database("ugly_db").Collection("movies")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	update := bson.D{{"$set", movie}}
	err := collection.FindOneAndUpdate(ctx, bson.M{"name": movie.Name}, update).Err()
	fmt.Println(err)
	return err
}

// DeleteMovie should delete a movie by name in the mongo database
func (m *MService) DeleteMovie(name string) error {
	collection := m.Client.Database("ugly_db").Collection("movies")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := collection.FindOneAndDelete(ctx, bson.M{"name": name}).Err()
	return err
}
