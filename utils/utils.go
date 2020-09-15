package utils

import (
	"blog/model/blogs"
	"crypto/sha1"
	"encoding/base64"
)

var salt = "iloveyou,andihateyou"

type Encrypt interface {
	Encrypt(str string) string
	Decrypt(str string) string
}

type BASE64 struct {
}

func (base BASE64) Encrypt(str string) string {
	return str
}

func (base BASE64) Decrypt(str string) string {
	return str
}

func GenerateTokenByAccount(a *blogs.User) (string, error) {
	if a.Account == "" || a.Password == "" {
		return "", nil
	} else {
		t := base64.StdEncoding.EncodeToString([]byte(a.Account + salt))
		sha := sha1.New()
		sha.Write([]byte(t))
		return base64.StdEncoding.EncodeToString([]byte(string(sha.Sum(nil)))), nil
	}
}
