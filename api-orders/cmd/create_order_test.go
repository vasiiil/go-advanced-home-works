package main

import (
	"api-orders/internal/models"
	"api-orders/internal/order"
	"bytes"
	"encoding/json"
	"fmt"
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
	db.Create(&models.User{
		Phone:     "79161880922",
		SessionId: "b84b0215-1a0c-48f6-b409-23cff5bc4980",
	})
	for i := range 5 {
		db.Create(&models.Product{
			Name:        fmt.Sprintf("Name of %d product", i+1),
			Description: fmt.Sprintf("Some very long text of description for %d product", i+1),
			Price:       200,
			Quantity:    20,
		})
	}
}

func clearData(db *gorm.DB) {
	query := `
		TRUNCATE TABLE 
				users,
				products
		RESTART IDENTITY CASCADE;
	`
	db.Exec(query)
}

func TestCreateOrderSuccess(t *testing.T) {
	db := initDb()
	initData(db)
	defer clearData(db)

	ts := httptest.NewServer(App())
	defer ts.Close()

	jsonBody, err := json.Marshal(order.OrderCreateRequest{
		ProductIds: []uint{1, 2, 3},
	})
	if err != nil {
		t.Fatal(err)
	}

	targetURL := ts.URL + "/orders"
	req, err := http.NewRequest(http.MethodPost, targetURL, bytes.NewReader(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uSWQiOiJiODRiMDIxNS0xYTBjLTQ4ZjYtYjQwOS0yM2NmZjViYzQ5ODAifQ.BCpw_B09kt0RIC0a68UE03t_8rWkxcQma4XorjXsrrs")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := io.ReadAll(resp.Body)
		t.Fatalf("Ожидался статус %d, получен %d. Тело ответа: %s", http.StatusCreated, resp.StatusCode, string(bodyBytes))
	}

	var responseOrder models.Order
	err = json.NewDecoder(resp.Body).Decode(&responseOrder)
	if err != nil {
		t.Fatal(err)
	}

	if responseOrder.ID == 0 {
		t.Fatal("Заказ не получил ID в ответе сервера")
	}
}
