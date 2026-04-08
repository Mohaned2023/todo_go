package auth

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// - Throws any, The rand function can read.
func HashPassword(password string) string {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		panic(err)
	}
	// See https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html#argon2id
	var (
		time      uint32 = 3
		memory    uint32 = 64 * 1024
		threads   uint8  = 2
		keyLength uint32 = 32
	)
	hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLength)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)
	encoded := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, memory, time, threads, b64Salt, b64Hash,
	)
	return encoded;
}

func ComparePassword(password string, encoded string) bool {
	parts := strings.Split(encoded, "$")
	var (
		time    uint32
		memory  uint32
		threads uint8
	)
	fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &time, &threads)
	salt, _ := base64.RawStdEncoding.DecodeString(parts[4])
	hash, _ := base64.RawStdEncoding.DecodeString(parts[5])
	keyLen := uint32(len(hash))

	computed := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)
	if subtle.ConstantTimeCompare(hash, computed) == 1 {
		return true
	}

	return false
}
