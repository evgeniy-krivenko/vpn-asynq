package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/evgeniy-krivenko/vpn-asynq/internal/entity"
	"github.com/evgeniy-krivenko/vpn-asynq/pkg/e"
	"io"
	mRand "math/rand"
	"modernc.org/strutil"
)

const Method = "chacha20-ietf-poly1305"

type Crypto interface {
	Encrypt(text, key string) (string, error)
	Decrypt(cipherText, key string) (string, error)
	GeneratePassword(passwordLen int) string
	GenerateSSConfig(conn *entity.Connection, cnf struct{ key, name string }) (string, error)
}

type CryptoService struct{}

// Encrypt service - key must be 16, 24 or 32 length
func (c CryptoService) Encrypt(text, key string) (string, error) {
	cpr, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", e.Warp("error created when encrypt cpr: %w", err)
	}

	gcm, err := cipher.NewGCM(cpr)
	if err != nil {
		return "", e.Warp("error created gcm: %w", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", e.Warp("error read nonce: %w", err)
	}

	return string(gcm.Seal(nonce, nonce, []byte(text), nil)), nil
}

func (c CryptoService) Decrypt(cipherText, key string) (string, error) {
	cpr, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", e.Warp("error created cpr when decrypt: %w", err)
	}

	gcm, err := cipher.NewGCM(cpr)
	if err != nil {
		return "", e.Warp("error created gcm when decrypt: %w", err)
	}

	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, ciphertext := cipherText[:nonceSize], cipherText[nonceSize:]
	dc, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		return "", e.Warp("error gcm open when decrypt: %w", err)
	}
	return string(dc), nil
}

func (c CryptoService) GeneratePassword(passwordLen int) string {
	passwordRunes := make([]rune, passwordLen)

	for i := range passwordRunes {
		passwordRunes[i] = randomRune()
	}

	return string(passwordRunes)
}

// GenerateSSConfig - key must be 16, 24 or 32 length
func (c CryptoService) GenerateSSConfig(conn *entity.Connection, cnf struct{ key, name string }) (string, error) {
	switch len(cnf.key) {
	default:
		return "", fmt.Errorf("wrong key len")
	case 16, 24, 32:
		break
	}
	plainSecret, err := c.Decrypt(conn.EncryptedSecret, cnf.key)
	if err != nil {
		return "", e.Warp("error decrypted when gen conf: %w", err)
	}

	userInfo := fmt.Sprintf("%s:%s", Method, plainSecret)

	encodedUserInfo := strutil.Base64Encode([]byte(userInfo))
	conf := fmt.Sprintf("ss://%s@%s:%d#%s", string(encodedUserInfo), conn.IpAddress, conn.Port, cnf.name)
	return conf, nil
}

func randomRune() rune {
	i := mRand.Intn(26)

	return rune('A' + i)
}
