package distribution

import (
	"fmt"
	"strings"

	"github.com/satori/go.uuid"
)

func GetUUID() string {
	u4, _ := uuid.NewV4()
	tempID := fmt.Sprintf("%s", u4)
	UUID := strings.Replace(tempID, "-", "", -1)
	return UUID
}
