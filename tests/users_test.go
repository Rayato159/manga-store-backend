package tests

import (
	"log"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/rayato159/manga-store/configs"
	"github.com/rayato159/manga-store/pkg/databases"
	"github.com/rayato159/manga-store/pkg/utils"
	"github.com/rayato159/manga-store/tests/entities_test"
	_testUsersRepository "github.com/rayato159/manga-store/tests/users_test/repositories_test"
	_testUsersUsecase "github.com/rayato159/manga-store/tests/users_test/usecases_test"
)

// Test users class
type testUsers struct {
	Cfg *configs.Configs
	Db  *sqlx.DB
}

func NewTestUsers() *testUsers {
	utils.LoadDotenv(".env.test")
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
	TestUsersUse entities_test.TestUsersRepository
}

func NewTestUsersController(testUsersUse entities_test.TestUsersUsecase) *testUsersCon {
	return &testUsersCon{
		TestUsersUse: testUsersUse,
	}
}

// Function to tests
func StartTestUsers(t *testing.T) {
	// Setup and load configs
	test := NewTestUsers()
	defer test.Db.Close()

	testUsersRepository := _testUsersRepository.NewTestUsersRepository(test.Db)
	testUsersUsecase := _testUsersUsecase.NewTestUsersUsecase(testUsersRepository)
	testUsersController := NewTestUsersController(testUsersUsecase)
	_ = testUsersController
}
