package viperEx

import "github.com/spf13/viper"

// Get a string value or default if not set
func GetStringOrDefault(key string, defaultValue string) string {
	value := viper.GetString(key)
	if value != "" {
		return value
	}
	return defaultValue
}
