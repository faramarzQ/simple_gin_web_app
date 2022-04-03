package helpers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// ENV returns environment value for the given key.
// If value is empty, the default value will be returned
func ENV(params ...string) string {

	if params[0] == "" {
		panic("The ENV method needs at least one argument for the key!")
	}

	err := godotenv.Load()
	if err != nil {
		LogError(err.Error())
	}

	value := os.Getenv(params[0])

	fmt.Println(len(params))
	if value == "" {
		value = params[1]
	}

	return value
}
