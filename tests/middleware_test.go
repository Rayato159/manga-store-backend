package tests

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/rayato159/manga-store/configs"
	"github.com/rayato159/manga-store/pkg/utils"
)

type testMiddlewares struct {
	Cfg *configs.Configs
}

type testAuthentication struct {
	Label      string
	Input      string
	StatusCode int
}

type testAuthorization struct {
	Label      string
	Input      string
	StatusCode int
}

func NewTestMiddlewares() *testMiddlewares {
	utils.LoadDotenv("../.env.test")
	cfg := new(configs.Configs)

	// Fiber configs
	cfg.Fiber.Host = os.Getenv("FIBER_HOST")
	cfg.Fiber.Port = os.Getenv("FIBER_PORT")
	cfg.Fiber.ServerRequestTimeout = os.Getenv("FIBER_REQUEST_TIMEOUT")

	// App
	cfg.App.Stage = os.Getenv("STAGE")
	cfg.App.Version = os.Getenv("APP_VERSION")
	cfg.App.AdminKey = os.Getenv("ADMIN_KEY")
	cfg.App.JwtSecretKey = os.Getenv("JWT_SECRET_KEY")
	cfg.App.JwtAccessTokenExpires = os.Getenv("JWT_ACCESS_TOKEN_EXPIRES")
	cfg.App.JwtRefreshTokenExpires = os.Getenv("JWT_REFRESH_TOKEN_EXPIRES")
	cfg.App.JwtSessionTokenExpires = os.Getenv("JWT_SESSION_TOKEN_EXPIRES")

	return &testMiddlewares{
		Cfg: cfg,
	}
}

func TestJwtAuthentication(t *testing.T) {
	test := NewTestMiddlewares()

	baseUrl, err := utils.ConnectionUrlBuilder("fiber", test.Cfg)
	if err != nil {
		t.Error(err.Error())
	}
	testUrl := fmt.Sprintf("http://%v/v1/tests/%v", baseUrl, "authentication")
	fmt.Printf("test url -> %v\n", testUrl)

	tests := []testAuthentication{
		{
			Label:      "empty token -> not pass",
			StatusCode: 401,
			Input:      "",
		},
		{
			Label:      "wrong token -> not pass",
			StatusCode: 401,
			Input:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOTc3YTRlZGUtZDBlZC00NWYwLTkzNTItNTcyMGUwNjdmMjRiIiwicm9sZSI6InVzZXIiLCJpc3MiOiJhY2Nlc3NfdG9rZW4iLCJzdWIiOiJ1c2Vyc19hY2Nlc3NfdG9rZW4iLCJhdWQiOlsidXNlcnMiXSwiZXhwIjozODEzMDcyNzgzLCJuYmYiOjE2NjU1ODkxMzYsImlhdCI6MTY2NTU4OTEzNiwianRpIjoiNDYyYzk1YjYtMzIwOS00YWMwLTllOTktYjAxZmU0Mzg2ZGNlIn0.rDnz6vXyY-opOHeEnky6eI3EY",
		},
		{
			Label:      "user -> pass",
			StatusCode: 200,
			Input:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOTc3YTRlZGUtZDBlZC00NWYwLTkzNTItNTcyMGUwNjdmMjRiIiwicm9sZSI6InVzZXIiLCJpc3MiOiJhY2Nlc3NfdG9rZW4iLCJzdWIiOiJ1c2Vyc19hY2Nlc3NfdG9rZW4iLCJhdWQiOlsidXNlcnMiXSwiZXhwIjozODEzMDcyNzgzLCJuYmYiOjE2NjU1ODkxMzYsImlhdCI6MTY2NTU4OTEzNiwianRpIjoiNDYyYzk1YjYtMzIwOS00YWMwLTllOTktYjAxZmU0Mzg2ZGNlIn0.rDnz6vXyY-opOHeEnky6eI3EYJ51bP52HWMsLpTpJPo",
		},

		{
			Label:      "admin -> pass",
			StatusCode: 200,
			Input:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNmRhNmY3OTUtYWVhMy00YWQzLTlhMWMtNDY5MTQzMTU2MDYxIiwicm9sZSI6ImFkbWluIiwiaXNzIjoiYWNjZXNzX3Rva2VuIiwic3ViIjoidXNlcnNfYWNjZXNzX3Rva2VuIiwiYXVkIjpbInVzZXJzIl0sImV4cCI6MzgxMzA3Mjc2MywibmJmIjoxNjY1NTg5MTE2LCJpYXQiOjE2NjU1ODkxMTYsImp0aSI6IjRhMzFlNTM3LTYyZWMtNGU1NC04ODQ5LWIxMDkyNmQ1ZmNmNiJ9.buacTlvCypAGGz0rjMX08YwOGAEQr1ljO1_JPOgdN_M",
		},
	}

	for i := range tests {
		fmt.Printf("case: %v -> %v\n", i+1, tests[i].Label)
		req, err := http.NewRequest("GET", testUrl, nil)
		if err != nil {
			t.Errorf("expect: %v but got -> %v", "<nil>", err.Error())
		}
		req.Header = http.Header{
			"Host":          {"0.0.0.0"},
			"Content-Type":  {"application/json"},
			"Authorization": {fmt.Sprintf("Bearer %v", tests[i].Input)},
		}
		client := &http.Client{}

		res, err := client.Do(req)
		if err != nil {
			t.Errorf("expect: %v but got -> %v", "<nil>", err.Error())
		}
		if res.StatusCode != tests[i].StatusCode {
			t.Errorf("expect: %v but got -> %v", tests[i].StatusCode, res.StatusCode)
		}
	}
}

func TestAuthorization(t *testing.T) {
	test := NewTestMiddlewares()

	baseUrl, err := utils.ConnectionUrlBuilder("fiber", test.Cfg)
	if err != nil {
		t.Error(err.Error())
	}
	testUrl := fmt.Sprintf("http://%v/v1/tests/%v", baseUrl, "authorization")
	fmt.Printf("test url -> %v\n", testUrl)

	tests := []testAuthorization{
		{
			Label:      "user -> not pass",
			StatusCode: 401,
			Input:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOTc3YTRlZGUtZDBlZC00NWYwLTkzNTItNTcyMGUwNjdmMjRiIiwicm9sZSI6InVzZXIiLCJpc3MiOiJhY2Nlc3NfdG9rZW4iLCJzdWIiOiJ1c2Vyc19hY2Nlc3NfdG9rZW4iLCJhdWQiOlsidXNlcnMiXSwiZXhwIjozODEzMDcyNzgzLCJuYmYiOjE2NjU1ODkxMzYsImlhdCI6MTY2NTU4OTEzNiwianRpIjoiNDYyYzk1YjYtMzIwOS00YWMwLTllOTktYjAxZmU0Mzg2ZGNlIn0.rDnz6vXyY-opOHeEnky6eI3EYJ51bP52HWMsLpTpJPo",
		},
		{
			Label:      "admin -> pass",
			StatusCode: 200,
			Input:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNmRhNmY3OTUtYWVhMy00YWQzLTlhMWMtNDY5MTQzMTU2MDYxIiwicm9sZSI6ImFkbWluIiwiaXNzIjoiYWNjZXNzX3Rva2VuIiwic3ViIjoidXNlcnNfYWNjZXNzX3Rva2VuIiwiYXVkIjpbInVzZXJzIl0sImV4cCI6MzgxMzA3Mjc2MywibmJmIjoxNjY1NTg5MTE2LCJpYXQiOjE2NjU1ODkxMTYsImp0aSI6IjRhMzFlNTM3LTYyZWMtNGU1NC04ODQ5LWIxMDkyNmQ1ZmNmNiJ9.buacTlvCypAGGz0rjMX08YwOGAEQr1ljO1_JPOgdN_M",
		},
	}

	for i := range tests {
		fmt.Printf("case: %v -> %v\n", i+1, tests[i].Label)
		req, err := http.NewRequest("GET", testUrl, nil)
		if err != nil {
			t.Errorf("expect: %v but got -> %v", "<nil>", err.Error())
		}
		req.Header = http.Header{
			"Host":          {"0.0.0.0"},
			"Content-Type":  {"application/json"},
			"Authorization": {fmt.Sprintf("Bearer %v", tests[i].Input)},
		}
		client := &http.Client{}

		res, err := client.Do(req)
		if err != nil {
			t.Errorf("expect: %v but got -> %v", "<nil>", err.Error())
		}
		if res.StatusCode != tests[i].StatusCode {
			t.Errorf("expect: %v but got -> %v", tests[i].StatusCode, res.StatusCode)
		}
	}
}
