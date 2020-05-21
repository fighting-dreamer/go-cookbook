package tools

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func GetBoolWithDefault(key string, defaultVal bool) bool {
	value := os.Getenv(key)
	if value == "" {
		value = viper.GetString(key)
	}

	boolVal, err := strconv.ParseBool(value)
	if err != nil {
		return defaultVal
	}
	return boolVal
}

func GetIntOrPanic(key string) int {
	CheckKey(key)
	v, err := strconv.Atoi(FatalGetString(key))
	PanicIfErrorForKey(err, key)
	return v
}

func GetInt64OrPanic(key string) int64 {
	CheckKey(key)
	v, err := strconv.Atoi(FatalGetString(key))
	PanicIfErrorForKey(err, key)
	return int64(v)
}

func FatalGetString(key string) string {
	CheckKey(key)
	value := os.Getenv(key)
	if value == "" {
		value = viper.GetString(key)
	}
	return value
}

func GetString(key string) string {
	value := os.Getenv(key)
	if value == "" {
		value = viper.GetString(key)
	}
	return value
}

func CheckKey(key string) {
	if !viper.IsSet(key) && os.Getenv(key) == "" {
		log.Fatalf("%s key is not set", key)
	}
}

func PanicIfErrorForKey(err error, key string) {
	if err != nil {
		log.Fatalf("Could not parse key: %s, Error: %s", key, err)
	}
}

func IsValidString(value string) bool {
	if len(value) == 0 || len(strings.TrimSpace(value)) == 0 {
		return false
	}
	return true
}

func GetBoolOrPanic(key string) bool {
	if !viper.IsSet(key) {
		return false
	}
	v, err := strconv.ParseBool(viper.GetString(key))
	PanicIfErrorForKey(err, key)
	return v
}

func GetUniqueID() string {
	return uuid.New().String()
}
