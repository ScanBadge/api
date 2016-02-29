package utility

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
)

// ReadData reads input from the specified file path.
func ReadData(filename string) ([]byte, error) {
	if filename == "" {
		return nil, fmt.Errorf("No filename specified")
	}

	return ioutil.ReadFile(filename)
}

// DecodeBase64 decodes the provided encoded base64 bytes.
func DecodeBase64(encoded []byte) ([]byte, error) {
	// Perhaps it's better to not cast it, but I could not get base64.StdEncoding.Decode to work... :(
	decoded, err := base64.StdEncoding.DecodeString(string(encoded))

	if err != nil {
		return nil, err
	}

	return decoded, nil
}

// HashPassword hashes a password with the bcrypt algorithm using a cost of 10.
func HashPassword(password string) string {
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}
