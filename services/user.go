package services

import (
	"context"
	"time"

	"github.com/thalissonfelipe/ugly-api/models"
	"github.com/thalissonfelipe/ugly-api/utils"

	c "github.com/thalissonfelipe/ugly-api/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UService struct
type UService struct {
	Client *mongo.Client
}

// Authenticate checks if the user passed valid username and login
func (u *UService) Authenticate(login *models.Login) error {
	collection := u.Client.Database(c.MyConfig.DB.DatabaseName).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result := collection.FindOne(ctx, bson.M{"username": login.Username})
	if result.Err() != nil {
		return result.Err()
	}

	exists := models.User{}
	err := result.Decode(&exists)
	if err != nil {
		return err
	}

	err = utils.ComparePassword(exists.Password, login.Password)

	return err
}

// GetUsers returns a lisst of users
func (u *UService) GetUsers() (*[]models.UserResponse, error) {
	collection := u.Client.Database(c.MyConfig.DB.DatabaseName).Collection("users")
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
func (u *UService) CreateUser(user *models.User) error {
	collection := u.Client.Database(c.MyConfig.DB.DatabaseName).Collection("users")
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
