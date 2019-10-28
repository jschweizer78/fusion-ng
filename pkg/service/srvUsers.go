package service

import (
	"fmt"

	"github.com/asdine/storm/v3"
	"github.com/jschweizer78/fusion-ng/pkg/service/model"
)

// SrvUsers is a service for users
type SrvUsers struct {
	DB       storm.Node
	MetaData *model.UserQuestion
}

// NewSrvUsers for user CRUD
func NewSrvUsers(db storm.Node) *SrvUsers {
	return &SrvUsers{
		DB:       db,
		MetaData: &model.UserQuestion{},
	}
}

// GetAll gets all users from storm DB
func (su *SrvUsers) GetAll() ([]*model.User, error) {
	var users []*model.User
	err := su.DB.All(users)
	if err != nil {
		return nil, fmt.Errorf("coul not find users: %v", err)
	}
	return users, nil
}

// GetOne gets one user from storm DB by email
func (su *SrvUsers) GetOne(id string) (*model.User, error) {
	var user *model.User
	err := su.DB.One("Email", id, user)
	if err != nil {
		return nil, fmt.Errorf("coul not find user %s: %v", id, err)
	}
	return user, nil
}

// GetLimit gets users from storm DB by email
func (su *SrvUsers) GetLimit(field, filter string, skip, limit int) ([]*model.User, error) {
	var users []*model.User
	err := su.DB.Find(field, filter, users, storm.Skip(skip), storm.Limit(limit))
	if err != nil {
		return nil, fmt.Errorf("coul not find users by %s: %v", field, err)
	}
	return users, nil
}

// CreateOne creats one user, must have email
func (su *SrvUsers) CreateOne(user *model.User) (string, error) {
	err := su.DB.Save(user)
	if err != nil {
		return "", fmt.Errorf("coul not create user %s: %v", user.Email, err)
	}
	return user.Email, nil
}

// DeleteOne deletes one user from storm by email
func (su *SrvUsers) DeleteOne(email string) error {
	err := su.DB.Delete(DBUsers, email)
	if err != nil {
		return fmt.Errorf("coul not delete user %s: %v", email, err)
	}
	return nil
}
