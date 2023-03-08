package names

import (
	"fmt"
	"math/rand"
	"strings"
)

func GetPhoneNumber() string {
	return fmt.Sprintf("%d%d%d", rand.Intn(999), rand.Intn(999), rand.Intn(999))
}

func GetCatchall(catchall, firstName, lastName string) (string, error) {
	if strings.HasPrefix(catchall, "@") {
		return fmt.Sprintf("%s%d%s%s", firstName, rand.Intn(99999), lastName, catchall), nil
	}

	return "", fmt.Errorf("catchall must start with @")
}
