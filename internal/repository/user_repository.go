package repository

import (
	"context"
	"time"

	"go-api-task/db/sqlc"
)

type UserRepository interface {
	CreateUser(ctx context.Context, name string, dob time.Time) (sqlc.User, error)
	UpdateUser(ctx context.Context, id int32, name string, dob time.Time) (sqlc.User, error)
	DeleteUser(ctx context.Context, id int32) error
	GetUserById(ctx context.Context, id int32) (sqlc.User, error)
	GetUser(ctx context.Context) ([]sqlc.User, error)
}

type userRepository struct {
	q *sqlc.Queries
}

func NewUserRepository(q *sqlc.Queries) UserRepository {
	return &userRepository{q: q}
}

func (r *userRepository) CreateUser(ctx context.Context,name string,dob time.Time) (sqlc.User, error) {
	return r.q.CreateUser(ctx, sqlc.CreateUserParams{
		Name: name,
		Dob:  dob,
	})
}

func (r *userRepository) UpdateUser(ctx context.Context,id int32,name string,dob time.Time) (sqlc.User, error) {
	return r.q.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  dob,
	})
}

func (r *userRepository) DeleteUser(ctx context.Context,id int32) error {
	return r.q.DeleteUser(ctx, id)
}

func (r *userRepository) GetUserById(ctx context.Context,id int32) (sqlc.User, error) {
	return r.q.GetUserById(ctx, id)
}

func (r *userRepository) GetUser(ctx context.Context) ([]sqlc.User, error) {
	return r.q.GetUser(ctx)
}
