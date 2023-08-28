package users_repository

import (
	"RestJwtAuth/internal/app/models/user"
	"RestJwtAuth/internal/pkg/app"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersCRUDRepository struct {
	collection mongo.Collection
	ctx        context.Context
}

func NewCRUDRepository(client *mongo.Client) (*UsersCRUDRepository, error) {
	c := *client.Database(app.DbName).Collection(app.UsersCollection)
	return &UsersCRUDRepository{collection: c}, nil
}

func (r *UsersCRUDRepository) Create(u user.User) error {
	_, err := r.collection.InsertOne(r.ctx, u)
	return err
}

// TODO: fix Read func
func (r *UsersCRUDRepository) Read(username string) (*user.User, error) {
	var u *user.User
	err := r.collection.FindOne(r.ctx, bson.M{"username": username}).Decode(&u)
	//fmt.Println(u.username, u.hashPassword)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// TODO: implement Update func
func (r *UsersCRUDRepository) Update(username string, usr user.User) (*user.User, error) {
	return nil, nil
}

// TODO: implement Delete func
func (r *UsersCRUDRepository) Delete(id string) error {
	return nil
}
