package viperEx

import (
	"strings"

	"github.com/spf13/viper"
)

// Get a string value or default if not set
func GetStringOrDefault(key string, defaultValue string) string {
	value := viper.GetString(key)
	if value != "" {
		return value
	}
	return defaultValue
}

func GetIntOrDefault(key string, defaultValue int) int {
	value := viper.GetInt(key)
	return value
	// if value != nil {
	// 	return value
	// }
	//return defaultValue
}

// Get a string slice from viper
// Fixes the pflag binding problem
func GetStringSlice(name string) []string {
	result := viper.GetStringSlice(name)
	if len(result) == 1 && strings.HasPrefix(result[0], "[") && strings.HasSuffix(result[0], "]") {
		result2 := strings.Trim(result[0], "[]")
		if result2 == "" {
			return []string{}
		}
		return strings.Split(result2, ",")
	}
	return result
}

func GetStringSliceOrDefault(key string, defaultValue []string) []string {
	value := GetStringSlice(key)
	if value != nil {
		return value
	}
	return defaultValue
}
