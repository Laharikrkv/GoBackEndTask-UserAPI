package handler

import (
	"strconv"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-api-task/internal/dto"
	"go-api-task/internal/service"

	"go.uber.org/zap" 
	"go-api-task/internal/logger"
)

var validate = validator.New() 


type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}
//-------------------------------------POST API-----------------------------------------
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req dto.CreateUserRequest
//Parsing Request Body into Dto	
if err := c.BodyParser(&req); err != nil {
	//logger
	logger.Log.Warn("Invalid JSON",zap.Error(err),)
	return c.Status(400).JSON(fiber.Map{"error": "invalid JSON"})
}
//Validating the Request Body
if err := validate.Struct(&req); err != nil {
	//logger
	logger.Log.Warn("Validation failed",zap.Any("errors", err),)
	var validationErrors validator.ValidationErrors
	formattedErrors := make([]dto.ApiError, 0)
	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			formattedErrors = append(formattedErrors, dto.ApiError{Field: fieldError.Field(), Msg:   dto.MsgForTag(fieldError.Tag())})
		}
	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": formattedErrors})
}
//Parsing string Dob into Time Dob
timeDob, _ := time.Parse("2006-01-02", req.Dob)
//Calling UserService' method to create a user
user, err := h.service.CreateUser(c.Context(), req.Name, timeDob)
if err != nil {
	//logger
	logger.Log.Error("Failed to create user from service",zap.Error(err))
	return c.Status(400).JSON(fiber.Map{"error": err.Error()})
}
//logger
logger.Log.Info("User created successfully",
zap.Int32("user_id", user.ID),
zap.String("name", user.Name),
)
return c.Status(201).JSON(dto.UserResponse{ID:   user.ID,Name: user.Name,Dob:  user.Dob.Format("2006-01-02")})
}
//-----------------------------------PUT API-------------------------------------------------
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	//Parsing path variable into int32
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		//logger
		logger.Log.Warn("Invalid JSON",zap.Error(err))
		return c.Status(400).JSON(fiber.Map{"error": "invalid user id"})
	}
	//Parsing request body
	var req dto.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		//logger
		logger.Log.Warn("Invalid JSON",zap.Error(err))
		return c.Status(400).JSON(fiber.Map{"error": "invalid JSON"})
	}
	//Validating the request body
	if err := validate.Struct(&req); err != nil {
		//logger
		logger.Log.Warn("Validation failed",zap.Any("errors", err))
		var validationErrors validator.ValidationErrors
		formattedErrors := make([]dto.ApiError, 0)
		if errors.As(err, &validationErrors) {
			for _, fieldError := range validationErrors {
				formattedErrors = append(formattedErrors, dto.ApiError{Field: fieldError.Field(),Msg:   dto.MsgForTag(fieldError.Tag())})
			}
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": formattedErrors})
	}
	//Parsing string Dob into Time Dob
	timeDob, _ := time.Parse("2006-01-02", req.Dob)
	//Calling UserService' method to update a user
	user, err := h.service.UpdateUser(c.Context(),int32(id),req.Name,timeDob)
	if err != nil {
			logger.Log.Error("Failed to update user from service",zap.Error(err))
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
		logger.Log.Info("User updated successfully",
		zap.Int32("user_id", user.ID),
		zap.String("name", user.Name),
	)
	return c.Status(201).JSON(dto.UserResponse{ID:   user.ID,Name: user.Name,Dob:  user.Dob.Format("2006-01-02"),})
}
//-------------------------DELETE API--------------------------------------------------
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error{
	//Parsing id into int32
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		//logger
		logger.Log.Warn("Invalid Id given to delete",zap.Error(err))
		return c.Status(400).JSON(fiber.Map{"error": "invalid user id"})
	}
	//Calling UserService' method to delete user by Id
	err1 := h.service.DeleteUser(c.Context(), int32(id))
	if err1 != nil {
		//logger
		logger.Log.Error("Failed to delete user from service",zap.Error(err))
		return c.Status(500).JSON(fiber.Map{"error": err1.Error()})
	}
//logger
logger.Log.Info("User Deleted successfully",
zap.Int("user_id", id),)
return c.Status(200).JSON(fiber.Map{"message": "Deleted user","id": id})
}
//---------------------------------GET API (ID)---------------------------------------------
func (h *UserHandler) GetUserById(c *fiber.Ctx) error{
	//Parsing id into int32
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		//logger
		logger.Log.Warn("Invalid Id given to get",zap.Error(err))
		return c.Status(400).JSON(fiber.Map{"error": "invalid user id"})
	}
	//Calling UserService' method to get a user by id
	user, err1 := h.service.GetUserById(c.Context(), int32(id))
	if err1 != nil{
		//logger
		logger.Log.Error("Failed to get user from service",zap.Error(err))
		return c.Status(500).JSON(fiber.Map{"error" : err1.Error()})
	}
//logger
logger.Log.Info("User Fetched successfully",
zap.Int("user_id", id))
return c.Status(200).JSON(user)
}
//-------------------------------------------GET API(LIST)-----------------------------------
func (h *UserHandler) GetUser(c *fiber.Ctx) error{
	//Calling UserService' method to get all users
	users, err := h.service.GetUser(c.Context())
	if err != nil{
		//logger
		logger.Log.Error("Failed to get all users from service",zap.Error(err))
		return c.Status(500).JSON(fiber.Map{"error" : err.Error()})
	}
	//logger
	logger.Log.Info("All Users Fetched successfully")
	return c.Status(200).JSON(users)
}