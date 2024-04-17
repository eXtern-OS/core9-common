package token

import (
	"github.com/eXtern-OS/core9-common/utils"
	"time"
)

type Token struct {
	T         string `json:"token"`
	Timestamp int64  `json:"timestamp"`
	UserId    string `json:"user_id"`
	IsRefresh bool   `json:"is_refresh"`
}

func (t *Token) IsValid() bool {
	return time.Now().Unix() < t.Timestamp+utils.Ternary[int64](t.IsRefresh, 2*24*60*60, 24*60*60)
}

func NewToken(userId string, isRefresh bool) Token {
	return Token{
		T:         utils.SHA256(time.Now().Format(time.RFC3339) + utils.RandomString() + userId),
		Timestamp: time.Now().Unix(),
		UserId:    userId,
		IsRefresh: isRefresh,
	}
}
