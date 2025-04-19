package common

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/vanthang24803/mini/internal/entity"
)

func GenerateCheckSum(user *entity.User) string {
	str := fmt.Sprintf("%s-%s-%s-%s-%s", user.ID.String(), user.Email, user.Username, user.Roles, user.HashedPassword)

	hash := sha256.Sum256([]byte(str))

	return hex.EncodeToString(hash[:])
}

func CompareCheckSum(str string) bool {
	hash, err := hex.DecodeString(str)
	if err != nil {
		return false
	}

	if len(hash) != sha256.Size {
		return false
	}

	return true
}
