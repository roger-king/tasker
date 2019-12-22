package services

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/roger-king/tasker/models"
	log "github.com/sirupsen/logrus"
)

// UserService -
type UserService struct {
	DB        *sqlx.DB
	TableName string
}

// NewUserService - initializes user service
func NewUserService(db *sqlx.DB) *UserService {
	return &UserService{
		DB:        db,
		TableName: "users",
	}
}

// CreateUser - creates a user
func (u *UserService) CreateUser(input *models.NewUserInput) (*models.UserDTO, error) {
	err := input.BeforeCreate()

	if err != nil {
		log.Error("Error encrypting token")
		return nil, err
	}

	sql, _, err := sq.Insert(u.TableName).Columns("email", "name", "username", "bio", "github_url", "access_token").Values(`:email, :name, :username, :bio, :github_url`).ToSql()

	if err != nil {
		log.Error("Failed to create insert query")
		return nil, err
	}

	_, err = u.DB.NamedExec(sql, map[string]interface{}{
		"email":      input.Email,
		"name":       input.Name,
		"username":   input.UserName,
		"bio":        input.Bio,
		"github_url": input.GitHubURL,
	})

	if err != nil {
		log.Error("Failed to insert:", err)
		return nil, err
	}

	return &models.UserDTO{
		Email:     input.Email,
		Name:      input.Name,
		Bio:       input.Bio,
		GitHubURL: input.GitHubURL,
		UserName:  input.UserName,
	}, nil
}
