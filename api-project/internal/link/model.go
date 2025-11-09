package link

import (
	"math/rand/v2"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url  string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
}

func New(url string) *Link {
	return &Link{
		Url: url,
		Hash: randStringRunes(10),
	}
}

var allowedRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
func randStringRunes(n int) string {
	runeSl := make([]rune, n)
	for i := range runeSl {
		runeSl[i] = allowedRunes[rand.IntN(len(allowedRunes))]
	}

	return string(runeSl)
}
