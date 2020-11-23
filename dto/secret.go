package dto

type SecretDto struct {
	ID            interface{}
	Key           string
	ApplicationId uint
	Value         []byte
}
