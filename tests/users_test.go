package tests

import (
	"context"
	"errors"
	"log"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/rayato159/manga-store/configs"
	"github.com/rayato159/manga-store/internals/entities"
	_usersRepository "github.com/rayato159/manga-store/internals/users/repositories"
	_usersUsecase "github.com/rayato159/manga-store/internals/users/usecases"
	"github.com/rayato159/manga-store/pkg/databases"
	"github.com/rayato159/manga-store/pkg/utils"
	"github.com/rayato159/manga-store/tests/entities_test"
)

// Test users class
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
	cfg.App.Version = os.Getenv("APP_VERSION")
	cfg.App.Stage = os.Getenv("STAGE")
	cfg.App.AdminKey = os.Getenv("ADMIN_KEY")

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

// Function to tests
func TestStartUsers(t *testing.T) {
	// Setup and load configs
	test := NewTestUsers()
	defer test.Db.Close()

	testUsersRepository := _usersRepository.NewUsersRepository(test.Db)
	testUsersUsecase := _usersUsecase.NewUsersUsecase(testUsersRepository)
	testUsersController := NewTestUsersController(testUsersUsecase)

	// *TestRegister
	// Case 1 -> Password and password confirm is not match
	// Case 2 -> Role is invalid
	// Case 3 -> Admin key is invalid
	// Case 4 -> Username is already taken
	// Case 5 -> Register Success user
	// Case 6 -> Register Success admin
	testUsersRegister := make([]entities_test.UsersRegisterTest, 0)
	for i := 0; i < 6; i++ {
		testUsersRegisterCase := entities_test.UsersRegisterTest{}
		testUsersRegisterCase.Input = new(entities.UsersRegisterReq)
		switch i {
		case 0:
			testUsersRegisterCase = entities_test.UsersRegisterTest{
				Input: &entities.UsersRegisterReq{
					Username:        "johndoe",
					Password:        "123456",
					ConfirmPassword: "111111",
					Role:            "user",
					AdminKey:        "",
				},
				Expect: "error, confirm password is not match",
			}
		case 1:
			testUsersRegisterCase = entities_test.UsersRegisterTest{
				Input: &entities.UsersRegisterReq{
					Username:        "johndoe",
					Password:        "123456",
					ConfirmPassword: "123456",
					Role:            "god",
					AdminKey:        "",
				},
				Expect: "error, role is invalid",
			}
		case 2:
			testUsersRegisterCase = entities_test.UsersRegisterTest{
				Input: &entities.UsersRegisterReq{
					Username:        "johndoe",
					Password:        "123456",
					ConfirmPassword: "123456",
					Role:            "admin",
					AdminKey:        "imadmin",
				},
				Expect: "error, admin key is invalid",
			}
		case 3:
			testUsersRegisterCase = entities_test.UsersRegisterTest{
				Input: &entities.UsersRegisterReq{
					Username:        "usertest",
					Password:        "123456",
					ConfirmPassword: "123456",
					Role:            "user",
					AdminKey:        "",
				},
				Expect: "error, username has been already taken",
			}
		case 4:
			testUsersRegisterCase = entities_test.UsersRegisterTest{
				Input: &entities.UsersRegisterReq{
					Username:        "user",
					Password:        "123456",
					ConfirmPassword: "123456",
					Role:            "user",
					AdminKey:        "",
				},
				Expect: "user",
			}
		case 5:
			testUsersRegisterCase = entities_test.UsersRegisterTest{
				Input: &entities.UsersRegisterReq{
					Username:        "admin",
					Password:        "123456",
					ConfirmPassword: "123456",
					Role:            "admin",
					AdminKey:        "UMHNTiXpstOZk3IB",
				},
				Expect: "admin",
			}
		}
		testUsersRegister = append(testUsersRegister, testUsersRegisterCase)
	}

	for i := 0; i < 4; i++ {
		_, err := testUsersController.Register(test.Cfg, testUsersRegister[i].Input)
		if err.Error() != testUsersRegister[i].Expect {
			t.Errorf("expect: <%v> but got -> <%v>", testUsersRegister[i].Expect, err.Error())
		}
	}

	result5, err5 := testUsersController.Register(test.Cfg, testUsersRegister[4].Input)
	if err5 != nil {
		t.Errorf("expect: <nil> but got -> <%v>", err5.Error())
	} else if result5.Username != "user" {
		t.Errorf("expect: <%v> but got -> <%v>", result5.Username, err5.Error())
	}

	result6, err6 := testUsersController.Register(test.Cfg, testUsersRegister[5].Input)
	if err6 != nil {
		t.Errorf("expect: <nil> but got -> <%v>", err6.Error())
	} else if result6.Username != "admin" {
		t.Errorf("expect: <%v> but got -> <%v>", result6.Username, err6.Error())
	}
}

func (tuc *testUsersCon) Register(cfg *configs.Configs, req *entities.UsersRegisterReq) (*entities.UsersRegisterRes, error) {
	ctx := context.WithValue(context.TODO(), entities_test.TestUsersCon, "TestCon.TestRegister")
	defer log.Println(ctx.Value(entities_test.TestUsersCon))

	if req.Password != req.ConfirmPassword {
		return nil, errors.New("error, confirm password is not match")
	}

	switch req.Role {
	case entities.Admin:
		if req.AdminKey != cfg.App.AdminKey {
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
