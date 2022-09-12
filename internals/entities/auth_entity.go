package entities

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthContext string

const (
	AuthCon AuthContext = "AuthController"
	AuthUse AuthContext = "AuthUsecase"
	AuthRep AuthContext = "AuthRepository"
)

type ClaimsType string

const (
	AccessToken  ClaimsType = "access_token"
	RefreshToken ClaimsType = "refresh_token"
	SessionToken ClaimsType = "session_token"
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

type UsersAccessToken struct {
	Id   string `db:"id" json:"id"`
	Role string `db:"role" json:"role"`
}

type UsersRefreshToken struct {
	Id        string `db:"id" json:"id"`
	Role      string `db:"role" json:"role"`
	ExpiresAt *time.Time
	IssuedAt  *time.Time
}

type UsersSessionToken struct {
	Id        string `db:"id" json:"id"`
	Username  string `db:"username" json:"username"`
	Role      string `db:"role" json:"role"`
	ExpiresAt *time.Time
	IssuedAt  *time.Time
}

type UsersJwtClaimsReq struct {
	UsersAccessToken  *UsersAccessToken
	UsersRefreshToken *UsersRefreshToken
	UsersSessionToken *UsersSessionToken
}

type UsersJwtTokenMapClaims struct {
	Id   string `db:"id" json:"id"`
	Role string `db:"role" json:"role"`
	jwt.RegisteredClaims
}

type UsersJwtSessionMapClaims struct {
	Id       string `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Role     string `db:"role" json:"role"`
	jwt.RegisteredClaims
}
