package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func replaceEnvVariables(input string) string {
    // Define a regular expression to match strings starting with $
    envPattern := regexp.MustCompile(`\$(\w+)`)

    // Replace all matches with their corresponding environment variable value
    replaced := envPattern.ReplaceAllStringFunc(input, func(match string) string {
        envName := strings.TrimPrefix(match, "$")
        envValue, found := os.LookupEnv(envName)
        if found {
            return envValue
        } else {
            return match // If the variable is not found, leave the original string as-is
        }
    })
    return replaced
}

func main() {
	// Define command-line flags
	filePath := flag.String("file", "", "File to inject environment variables")
	flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()

	// Load the file contents
	fileBytes, err := ioutil.ReadFile(*filePath)
	if err != nil {
			panic(err)
	}

	fileContent := string(fileBytes)
	updatedContent := replaceEnvVariables(fileContent)

	// Check if debug flag is true
  if flag.Lookup("debug").Value.String() == "true" {
		fmt.Println(updatedContent)
	}

	// Write the updated content to the same file
	err = ioutil.WriteFile(*filePath, []byte(updatedContent), 0644)
	if err != nil {
			panic(err)
	}
	// Print a success message
	fmt.Println("File updated successfully!")
}