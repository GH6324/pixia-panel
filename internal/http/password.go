package httpapi

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 12

func hashPassword(password string) (string, error) {
	if strings.TrimSpace(password) == "" {
		return "", errors.New("empty password")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// verifyPassword returns (ok, needsUpgrade, err).
func verifyPassword(storedHash, password string) (bool, bool, error) {
	if storedHash == "" {
		return false, false, errors.New("empty hash")
	}
	if strings.HasPrefix(storedHash, "$2a$") || strings.HasPrefix(storedHash, "$2b$") || strings.HasPrefix(storedHash, "$2y$") {
		if err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password)); err != nil {
			return false, false, nil
		}
		return true, false, nil
	}
	if isMD5Hash(storedHash) {
		if md5Hex(password) == storedHash {
			return true, true, nil
		}
		return false, false, nil
	}
	return false, false, errors.New("unknown password hash format")
}

func isMD5Hash(val string) bool {
	if len(val) != 32 {
		return false
	}
	for _, r := range val {
		if (r < '0' || r > '9') && (r < 'a' || r > 'f') && (r < 'A' || r > 'F') {
			return false
		}
	}
	return true
}

func md5Hex(val string) string {
	h := md5.Sum([]byte(val))
	return hex.EncodeToString(h[:])
}
