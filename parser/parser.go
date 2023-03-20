package parser

import (
	"michaelhenry/envject/value_encoders"
	"os"
	"regexp"
	"strings"
)

func ReplaceEnvVariables(input string, ignorePattern string, valueEncoder value_encoders.ValueEncoder) string {
	// Handling the following formats:
	// - $ENV_NAME
	// - ${ENV_NAME}
	// - $(ENV_NAME)
	envPattern := regexp.MustCompile(`\$(\w+)|\$(\{?(\w+)\})\.?|\$(\(?(\w+)\))\.?`)

	// Replace all matches with the corresponding environment variable value
	modifiedContent := envPattern.ReplaceAllStringFunc(input, func(match string) string {
		envName := strings.TrimPrefix(match, "$")
		envName = strings.TrimPrefix(envName, "(")
		envName = strings.TrimPrefix(envName, "{")
		envName = strings.TrimSuffix(envName, "}")
		envName = strings.TrimSuffix(envName, ")")

		if ignorePattern != "" {
			if regexp.MustCompile(ignorePattern).MatchString(envName) {
				return match
			}
		}

		envValue, found := os.LookupEnv(envName)
		if found {
			return valueEncoder.Encode(envValue)
		} else {
			return match // If the variable is not found, leave the original string as-is
		}
	})

	if  valueEncoder.AdditionalCode() != "" && modifiedContent != input && !strings.Contains(modifiedContent, valueEncoder.AdditionalCode()) {
		modifiedContent += valueEncoder.AdditionalCode()
	}
	return modifiedContent
}