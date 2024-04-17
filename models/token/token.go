package token

import (
	"github.com/eXtern-OS/core9-common/utils"
	"time"
)

type Token struct {
	T         string `json:"token"`
	Timestamp int64  `json:"timestamp"`
	UserId    string `json:"user_id"`
}

func (t *Token) IsValid() bool {
	return time.Now().Unix() < t.Timestamp+24*60*60
}

func NewToken() Token {
	return Token{
		T:         utils.SHA256(time.Now().Format(time.RFC3339) + utils.RandomString()),
		Timestamp: time.Now().Unix(),
	}
}
