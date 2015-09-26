package authtoken

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
)

// Token build a token from the triplet login password salt
func Token(login, password, salt string) string {

	var buffer bytes.Buffer

	buffer.WriteString(login)
	buffer.WriteString(password)
	buffer.WriteString(salt)

	hash := sha256.Sum256(buffer.Bytes())

	return base64.URLEncoding.EncodeToString(hash[:32])
}
