package salt

import (
	"encoding/hex"

	"golang.org/x/crypto/scrypt"
)

func Convert(pass, salt string) (string, error) {
	converted, err := scrypt.Key([]byte(pass), []byte(salt), 16384, 8, 1, 32)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(converted[:]), nil
}
