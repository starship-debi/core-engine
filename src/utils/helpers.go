package coreengine

import (
	"math/rand"
	"time"
)

func GenerateRandomID() string {
	rand.Seed(time.Now().UnixNano())
	id := make([]byte, 32)
	_, err := rand.Read(id)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", id)
}

func GetUUID() string {
	uuid := uuid.NewV4()
	return uuid.String()
}

func IsUUID(s string) bool {
	_, err := uuid.Parse(s)
	return err == nil
}

func GetEnvironmentVariable(name string) string {
	value := os.Getenv(name)
	if value == "" {
		return ""
	}
	return value
}

func GetEnvironmentVariableOrDefault(name string, defaultValue string) string {
	value := GetEnvironmentVariable(name)
	if value == "" {
		return defaultValue
	}
	return value
}