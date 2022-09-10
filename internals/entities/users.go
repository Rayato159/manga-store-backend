package entities

import (
	"context"
	"time"
)

type UsersContext string

const (
	UsersCon UsersContext = "UsersController"
	UsersUse UsersContext = "UsersUsecase"
	UsersRep UsersContext = "UsersRepository"
)

type UsersUsecase interface {
	Register(ctx context.Context, req *UsersRegisterReq) (*UsersRegisterRes, error)
}

type UsersRepository interface {
	Register(ctx context.Context, req *UsersRegisterReq) (*UsersRegisterRes, error)
	FindOneUser(ctx context.Context, req string) (*UserInfo, error)
}

type UsersRole string

const (
	Admin UsersRole = "admin"
	User  UsersRole = "user"
)

type UsersRegisterReq struct {
	Username        string    `db:"username" json:"username" copier:"Username"`
	Password        string    `db:"password" json:"password"`
	ConfirmPassword string    `json:"confirm_passowrd"`
	Role            UsersRole `db:"role" json:"role"`
	AdminKey        string    `json:"admin_key"`
}

type UsersRegisterRes struct {
	Username  string    `db:"username" json:"username" copier:"Username"`
	CreatedAt time.Time `db:"created_at" json:"created_at" copier:"CreatedAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at" copier:"UpdatedAt"`
}

type UserInfo struct {
	Username string `db:"username"`
	Password string `db:"password"`
}
