package services

import (
	"context"
	"time"

	"github.com/thalissonfelipe/ugly-api/config"
	"github.com/thalissonfelipe/ugly-api/models"
	"github.com/thalissonfelipe/ugly-api/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserService struct
type UserService struct {
	Client *mongo.Client
}

// Authenticate ...
func (u *UserService) Authenticate(login *models.Login) error {
	return utils.CheckUser(login.Username, login.Password, u.Client)
}

// GetUsers returns a lisst of users
func (u *UserService) GetUsers() (*[]models.UserResponse, error) {
	collection := u.Client.Database(config.MyConfig.DB.DatabaseName).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	var users = make([]models.UserResponse, 0)
	if err != nil {
		return &users, err
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, &users)

	return &users, err
}

// CreateUser creates a new user
func (u *UserService) CreateUser(user *models.User) error {
	collection := u.Client.Database(config.MyConfig.DB.DatabaseName).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bson.M{"username": user.Username}).Err()
	if err == nil {
		return utils.ErrAlreadyExists
	}
	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashPassword
	_, err = collection.InsertOne(ctx, user)

	return err
}
