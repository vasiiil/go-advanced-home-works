package auth_test

import (
	"api-project/configs"
	"api-project/internal/auth"
	"api-project/internal/user"
	"api-project/pkg/db"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func bootstrap() (*auth.AuthHandler, sqlmock.Sqlmock, error) {
	database, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: database,
	}))
	if err != nil {
		return nil, nil, err
	}

	userRepo := user.NewRepository(&db.Db{
		DB: gormDb,
	})

	handler := auth.AuthHandler{
		Config: &configs.AuthConfig{
			JwtSecret: "secret",
		},
		Service: auth.NewService(userRepo),
	}

	return &handler, mock, nil
}

func TestLoginHandlerSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	rows := sqlmock.NewRows([]string{"email", "password"}).
		AddRow("a2@a.ru", "$2a$10$o0luk/j/5b0RMqqtZIB7getvcVIPBOgHT31/7tnQ9gcPzKzD10G96")
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	if err != nil {
		t.Fatal(err)
	}

	data, _ := json.Marshal(auth.LoginRequest{
		Email:    "a2@a.ru",
		Password: "123456",
	})

	reader := bytes.NewReader(data)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/auth/login", reader)
	handler.Login()(w, r)

	if w.Code != http.StatusAccepted {
		t.Errorf("Got %d, expected, %d", w.Code, http.StatusAccepted)
	}
}

func TestRegisterHandlerSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	rows := sqlmock.NewRows([]string{"email", "password", "name"})
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow("1"),
		)
	mock.ExpectCommit()

	if err != nil {
		t.Fatal(err)
	}

	data, _ := json.Marshal(auth.RegisterRequest{
		Email:    "a2@a.ru",
		Password: "123456",
		Name:     "Test V",
	})

	reader := bytes.NewReader(data)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/auth/register", reader)
	handler.Register()(w, r)

	if w.Code != http.StatusCreated {
		t.Errorf("Got %d, expected, %d", w.Code, http.StatusCreated)
	}
}
