package parser

import (
	"os"
	"regexp"
	"strings"
)


func ReplaceEnvVariables(input string) string {
	// Handling the following formats:
	// - $ENV_NAME
	// - ${ENV_NAME}
	// - $(ENV_NAME)
	envPattern := regexp.MustCompile(`\$(\w+)|\$(\{?(\w+)\})\.?|\$(\(?(\w+)\))\.?`)

	// Replace all matches with the corresponding environment variable value
	return envPattern.ReplaceAllStringFunc(input, func(match string) string {
		envName := strings.TrimPrefix(match, "$")
		envName = strings.TrimPrefix(envName, "(")
		envName = strings.TrimPrefix(envName, "{")
		envName = strings.TrimSuffix(envName, "}")
		envName = strings.TrimSuffix(envName, ")")

		envValue, found := os.LookupEnv(envName)
		if found {
			return envValue
		} else {
			return match // If the variable is not found, leave the original string as-is
		}
	})
}