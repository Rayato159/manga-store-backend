package entities

import "context"

type AuthContext string

const (
	AuthCon AuthContext = "AuthController"
	AuthUse AuthContext = "AuthUsecase"
	AuthRep AuthContext = "AuthRepository"
)

type AuthRepository interface {
}

type AuthUsecase interface {
	Login(ctx context.Context, req *UsersCredentialsReq) (*UsersCredentialsRes, error)
}

type UsersCredentialsReq struct {
	Username string `db:"username" json:"username" form:"username"`
	Password string `db:"password" json:"password" form:"password"`
}

type UsersCredentialsRes struct {
	AccessToken  string `db:"access_token" json:"access_token"`
	RefreshToken string `db:"refresh_token" json:"refresh_token"`
	SessionToken string `db:"session_token" json:"session_token"`
}

type UsersPayload struct {
	Id       string `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Role     string `db:"role" json:"role"`
}
