package entities

type AuthContext string

const (
	AuthCon AuthContext = "AuthController"
	AuthUse AuthContext = "AuthUsecase"
	AuthRep AuthContext = "AuthRepository"
)

type AuthRepository interface {
}

type AuthUsecase interface {
}

type Credentials struct {
	Username string `db:"username" json:"username" form:"username"`
	Password string `db:"password" json:"password" form:"password"`
}
