package base

import (
	"errors"
	"golang.org/x/crypto/chacha20poly1305"

)

const (
	SecretKeyLength  = chacha20poly1305.KeySize
	PublicKeyLength  = 32
	PrivateKeyLength = 32
)

var (
	ErrNotEnoughBytes = errors.New("not enough bytes read from crypto random source")
	ErrKeyLength      = errors.New("key has to be 32 bytes long")
)

// Encryption - Interface for encryption and decryption
type Encryption interface {
	Encrypt(dst, msg []byte) ([]byte, error)
	EncryptString(msg string) ([]byte, error)
	Decrypt(dst, msg []byte) ([]byte, error)
	DecryptString(msg []byte) (string, error)
}
