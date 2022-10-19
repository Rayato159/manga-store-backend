package tests

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rayato159/manga-store/configs"
	"github.com/rayato159/manga-store/internals/entities"
	_usersRepository "github.com/rayato159/manga-store/internals/users/repositories"
	_usersUsecase "github.com/rayato159/manga-store/internals/users/usecases"
	"github.com/rayato159/manga-store/pkg/databases"
	"github.com/rayato159/manga-store/pkg/utils"
)

type UsersRegisterTest struct {
	Input  *entities.UsersRegisterReq
	Expect string
}

type testUsers struct {
	Cfg *configs.Configs
	Db  *sqlx.DB
}

func NewTestUsers() *testUsers {
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

	return &testUsers{
		Cfg: cfg,
		Db:  db,
	}
}

// Fake user controller
type testUsersCon struct {
	TestUsersUse entities.UsersUsecase
}

func NewTestUsersController(testUsersUse entities.UsersUsecase) *testUsersCon {
	return &testUsersCon{
		TestUsersUse: testUsersUse,
	}
}

// Test struct
type testRegister struct {
	Label  string
	Input  *entities.UsersRegisterReq
	Expect any
	Type   string
}

// Function to tests
func TestRegister(t *testing.T) {
	// Setup and load configs
	test := NewTestUsers()
	defer test.Db.Close()

	testUsersRepository := _usersRepository.NewUsersRepository(test.Db)
	testUsersUsecase := _usersUsecase.NewUsersUsecase(testUsersRepository)
	testUsersController := NewTestUsersController(testUsersUsecase)

	// Test setup
	tests := []testRegister{
		{
			Label: "error, confirm password is not match",
			Input: &entities.UsersRegisterReq{
				Username:        "johndoe",
				Password:        "123456",
				ConfirmPassword: "111111",
				Role:            "user",
				Key:             "",
			},
			Expect: "error, confirm password is not match",
			Type:   "error",
		},
		{
			Label: "error, role is invalid",
			Input: &entities.UsersRegisterReq{
				Username:        "johndoe",
				Password:        "123456",
				ConfirmPassword: "123456",
				Role:            "god",
				Key:             "",
			},
			Expect: "error, role is invalid",
			Type:   "error",
		},
		{
			Label: "error, admin key is invalid",
			Input: &entities.UsersRegisterReq{
				Username:        "johndoe",
				Password:        "123456",
				ConfirmPassword: "123456",
				Role:            "admin",
				Key:             "imadmin",
			},
			Expect: "error, admin key is invalid",
			Type:   "error",
		},
		{
			Label: "error, username has been already taken",
			Input: &entities.UsersRegisterReq{
				Username:        "usertest",
				Password:        "123456",
				ConfirmPassword: "123456",
				Role:            "user",
				Key:             "",
			},
			Expect: "error, username has been already taken",
			Type:   "error",
		},
		{
			Label: "no error, user",
			Input: &entities.UsersRegisterReq{
				Username:        "user",
				Password:        "123456",
				ConfirmPassword: "123456",
				Role:            "user",
				Key:             "",
			},
			Expect: "no error, user",
			Type:   "result",
		},
		{
			Label: "no error, admin",
			Input: &entities.UsersRegisterReq{
				Username:        "admin",
				Password:        "123456",
				ConfirmPassword: "123456",
				Role:            "admin",
				Key:             "UMHNTiXpstOZk3IB",
			},
			Expect: "no error, admin",
			Type:   "result",
		},
	}

	for i := range tests {
		switch tests[i].Type {
		case "error":
			fmt.Printf("case: %v -> %v\n", i+1, tests[i].Label)
			if _, err := testUsersController.Register(test.Cfg, tests[i].Input); err.Error() != tests[i].Expect.(string) {
				t.Errorf("expect: %v but got -> %v", tests[i].Expect.(string), err.Error())
			}
		case "result":
			fmt.Printf("case: %v -> %v\n", i+1, tests[i].Label)
			result, err := testUsersController.Register(test.Cfg, tests[i].Input)
			if err != nil {
				t.Errorf("expect: %v but got -> %v", "<nil>", err.Error())
			} else if result == nil {
				t.Errorf("expect: %v but got -> %v", "result", "<nil>")
			}
		}
	}
}

func (tuc *testUsersCon) Register(cfg *configs.Configs, req *entities.UsersRegisterReq) (*entities.UsersRegisterRes, error) {
	ctx := context.WithValue(context.Background(), entities.UsersCon, time.Now().UnixMilli())
	log.Printf("called:\t%v", utils.Trace())
	defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.UsersCon).(int64)))

	if req.Password != req.ConfirmPassword {
		return nil, errors.New("error, confirm password is not match")
	}

	switch req.Role {
	case entities.Admin:
		if req.Key != cfg.App.AdminKey {
			return nil, errors.New("error, admin key is invalid")
		}
	case entities.User:
	default:
		return nil, errors.New("error, role is invalid")
	}

	res, err := tuc.TestUsersUse.Register(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
