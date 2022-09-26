package tests

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/rayato159/manga-store/configs"

	_authRepository "github.com/rayato159/manga-store/internals/auth/repositories"
	_authUsecase "github.com/rayato159/manga-store/internals/auth/usecases"
	_usersRepository "github.com/rayato159/manga-store/internals/users/repositories"

	"github.com/rayato159/manga-store/internals/entities"
	"github.com/rayato159/manga-store/pkg/databases"
	"github.com/rayato159/manga-store/pkg/utils"
	"github.com/rayato159/manga-store/tests/entities_test"
)

// Test users class
type testAuth struct {
	Cfg *configs.Configs
	Db  *sqlx.DB
}

func NewTestAuth() *testAuth {
	utils.LoadDotenv("../.env.test")
	cfg := new(configs.Configs)

	// Fiber configs
	cfg.Fiber.Host = os.Getenv("FIBER_HOST")
	cfg.Fiber.Port = os.Getenv("FIBER_PORT")
	cfg.Fiber.ServerRequestTimeout = os.Getenv("FIBER_REQUEST_TIMEOUT")

	// Database Configs
	cfg.PostgreSQL.Host = os.Getenv("DB_HOST")
	cfg.PostgreSQL.Port = os.Getenv("DB_PORT")
	cfg.PostgreSQL.Protocol = os.Getenv("DB_PROTOCOL")
	cfg.PostgreSQL.Username = os.Getenv("DB_USERNAME")
	cfg.PostgreSQL.Password = os.Getenv("DB_PASSWORD")
	cfg.PostgreSQL.Database = os.Getenv("DB_DATABASE")

	// Redis
	cfg.Redis.Host = os.Getenv("REDIS_HOST")
	cfg.Redis.Port = os.Getenv("REDIS_PORT")
	cfg.Redis.Password = os.Getenv("REDIS_PASSWORD")
	cfg.Redis.Database = os.Getenv("REDIS_DATABASE")

	// App
	cfg.App.Stage = os.Getenv("STAGE")
	cfg.App.Version = os.Getenv("APP_VERSION")
	cfg.App.AdminKey = os.Getenv("ADMIN_KEY")
	cfg.App.JwtSecretKey = os.Getenv("JWT_SECRET_KEY")
	cfg.App.JwtAccessTokenExpires = os.Getenv("JWT_ACCESS_TOKEN_EXPIRES")
	cfg.App.JwtRefreshTokenExpires = os.Getenv("JWT_REFRESH_TOKEN_EXPIRES")
	cfg.App.JwtSessionTokenExpires = os.Getenv("JWT_SESSION_TOKEN_EXPIRES")

	// New Database
	db, err := databases.NewPostgreSQLDBConnection(cfg)
	if err != nil {
		log.Fatalf(err.Error())
		defer db.Close()
	}

	return &testAuth{
		Cfg: cfg,
		Db:  db,
	}
}

// Fake user controller
type testAuthCon struct {
	TestAuthUse entities.AuthUsecase
}

func NewTestAuthController(testAuthUse entities.AuthUsecase) *testAuthCon {
	return &testAuthCon{
		TestAuthUse: testAuthUse,
	}
}

// Function to tests
func TestStartAuth(t *testing.T) {
	// Setup and load configs
	test := NewTestAuth()
	defer test.Db.Close()

	usersRepository := _usersRepository.NewUsersRepository(test.Db)

	testAuthRepository := _authRepository.NewAuthRepository(test.Db)
	testAuthUsecase := _authUsecase.NewAuthUsecase(testAuthRepository, usersRepository)
	testAuthController := NewTestAuthController(testAuthUsecase)
	_ = testAuthController

	// *TestLogin
	// Case -> 1 username is invalid
	// Case -> 2 Password is invalid
	// Case -> 3 Claims type is invalid
	// Case -> 4 Can't claims the access token
	// Case -> 5 Can't claims the refresh token
	// Case -> 6 Can't claims the session token
}

func (tuc *testUsersCon) Login(cfg *configs.Configs, req *entities.UsersRegisterReq) error {
	ctx := context.WithValue(context.TODO(), entities_test.TestUsersCon, "TestCon.TestRegister")
	defer log.Println(ctx.Value(entities_test.TestUsersCon))

	return nil
}
