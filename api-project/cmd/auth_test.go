package main

import (
	"api-project/internal/auth"
	"api-project/internal/user"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDb() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func initData(db *gorm.DB) {
	db.Create(&user.User{
		Email:    "a2@a.ru",
		Password: "$2a$10$o0luk/j/5b0RMqqtZIB7getvcVIPBOgHT31/7tnQ9gcPzKzD10G96",
		Name:     "Test V",
	})
}

func clearData(db *gorm.DB) {
	db.Unscoped().
		Where("email = ?", "a2@a.ru").
		Delete(&user.User{})

}

func TestLoginSussess(t *testing.T) {
	db := initDb()
	initData(db)
	defer clearData(db)

	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(auth.LoginRequest{
		Email:    "a2@a.ru",
		Password: "123456",
	})

	response, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != http.StatusAccepted {
		t.Fatalf("Expected %d got %d", http.StatusAccepted, response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	var responseData auth.LoginResponse
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		t.Fatal(err)
	}
	if responseData.Token == "" {
		t.Fatal("Token is empty")
	}
}

func TestLoginFailed(t *testing.T) {
	db := initDb()
	initData(db)
	defer clearData(db)

	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(auth.LoginRequest{
		Email:    "a2@a.ru",
		Password: "1234567",
	})

	response, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != http.StatusUnauthorized {
		t.Fatalf("Expected %d got %d", http.StatusUnauthorized, response.StatusCode)
	}
}
