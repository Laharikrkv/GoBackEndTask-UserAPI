package service

import (
	"context"
	"time"

	"go-api-task/internal/logger"
	"go.uber.org/zap" 

	"go-api-task/db/sqlc"
	"go-api-task/internal/repository"
	"go-api-task/internal/dto"
)

type UserService interface {
	CreateUser(ctx context.Context, name string, dob time.Time) (sqlc.User, error)
	UpdateUser(ctx context.Context, id int32, name string, dob time.Time) (sqlc.User, error)
	DeleteUser(ctx context.Context, id int32) error
	GetUserById(ctx context.Context, id int32) (dto.UserAgeResponse, error)
	GetUser(ctx context.Context) ([]dto.UserAgeResponse, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context,name string,dob time.Time,) (sqlc.User, error) {
	//logger
	logger.Log.Info("Creating user in service",zap.String("name", name))
	return s.repo.CreateUser(ctx, name, dob)
}

func (s *userService) UpdateUser(ctx context.Context,id int32,name string,dob time.Time) (sqlc.User, error) {
	//logger
	logger.Log.Info("Updating user in service",zap.Int32("id", id))
	return s.repo.UpdateUser(ctx, id, name, dob)
}

func (s *userService) DeleteUser(ctx context.Context,id int32) (error) {
	//logger
	logger.Log.Info("Deleting user in service",zap.Int32("id", id))
	return s.repo.DeleteUser(ctx, id)
}

func (s *userService) GetUserById(ctx context.Context,id int32) (dto.UserAgeResponse, error) {
	//logger
	logger.Log.Info("Fetching user in service",zap.Int32("id", id))
	user, err := s.repo.GetUserById(ctx, id)
	if err != nil {
		return dto.UserAgeResponse{}, err
	}
	//Calculating Age from DOB
	age := calculateAge(user.Dob)
	//Saving age into Response DTO
	resp := dto.UserAgeResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Format("2006-01-02"),
		Age:  age,
	}
	logger.Log.Info("Fetched user with Age",zap.Int32("id", id))
	return resp, nil
}

func (s *userService) GetUser(ctx context.Context,)([]dto.UserAgeResponse, error){
	//logger
	logger.Log.Info("Fetching all users in service")
	users, err := s.repo.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	//Saving slice of Response DTO
	responses := make([]dto.UserAgeResponse, 0, len(users))
	for _, user := range users {
		age := calculateAge(user.Dob)
		responses = append(responses, dto.UserAgeResponse{ID:   user.ID,Name: user.Name,Dob:  user.Dob.Format("2006-01-02"),Age:  age})
	}
	return responses, nil
}

//Function to calculate Age of user to send while fetching user details
func calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()
	//Checking the day of DOB is passed by or not
	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}