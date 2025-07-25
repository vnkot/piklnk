package link

import (
	"math/rand"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url    string `json:"url"`
	Hash   string `json:"hash" gorm:"uniqueIndex"`
	UserId *uint  `json:"userId" gorm:"index;default:null"`
}

func NewLink(url string, userId *uint) *Link {
	link := &Link{
		Url:    url,
		UserId: userId,
	}
	link.GenerateHash()

	return link
}

func (link *Link) GenerateHash() {
	link.Hash = RandStringRunes(6)
}

var letterRunes = []rune("abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
