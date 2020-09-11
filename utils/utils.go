package utils

type Encrypt interface {
	Encrypt(str string) string
	Decrypt(str string) string
}

type BASE64 struct {

}

func (base BASE64)Encrypt(str string)string{
	return str
}

func (base BASE64)Decrypt(str string)string{
	return str
}