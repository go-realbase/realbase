package utils

import (
	"github.com/twinj/uuid"
	"strings"
)

func GetUUID() string {
	return uuid.NewV4().String()
}

func GetCleanUUID() string {
	return strings.Replace(GetUUID(), "-", "", -1)
}

func WhitelistFields(fields []string, obj map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	for _, k := range fields {
		result[k] = obj[k]
	}

	return result
}

func PathOfUrl(p string) string {
	split := strings.Split(p, "/")
	return split[len(split)-1]
}
