package user

import (
	"errors"
	"fmt"
	"github.com/eXtern-OS/core9-common/utils"
	"strings"
)

type User struct {
	Name           string `json:"name"`
	Login          string `json:"login"`
	PasswordHashed string `json:"password_hashed"`

	PublicKey           string `json:"public_key"`            //base64 Encoded
	PrivateKeyEncrypted string `json:"private_key_encrypted"` //base64 Encoded, Encrypted with AES
}

func (u *User) Verify() error {
	allowed := utils.ArrayToMap(strings.Split("abcdefghijklmnopqrstuvwxyz01234567890_", ""), true)
	for _, x := range strings.Split(strings.ToLower(u.Login), "") {
		if _, exists := allowed[x]; !exists {
			return errors.New(fmt.Sprintf("invalid character in the login: '%s'", x))
		}
	}
	return nil
}
