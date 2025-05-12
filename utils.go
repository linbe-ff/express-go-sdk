package express

import (
	"github.com/google/uuid"
	"strings"
)

func GenerateToken() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
