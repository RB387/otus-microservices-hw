package password

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
)

const saltCharacters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ,./!@#$%^&*()_+<>?:{}"
const saltLength = 10

func GenerateSalt() (string, error) {
	ret := make([]byte, saltLength)
	for i := 0; i < saltLength; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(saltCharacters))))
		if err != nil {
			return "", err
		}

		ret[i] = saltCharacters[num.Int64()]
	}

	return string(ret), nil
}

func CalculatePassword(rawPassword string, salt string) string {
	shaHash := sha256.New()

	return hex.EncodeToString(shaHash.Sum([]byte(rawPassword + salt)))
}
