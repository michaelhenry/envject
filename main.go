package main

import (
	"flag"
	"fmt"
	"michaelhenry/envject/parser"
	"michaelhenry/envject/value_encoders"
	"os"
	"strings"
)


func main() {
	// Define command-line flags
	sourcePath := flag.String("file", "", "File to inject the environment variables")
	outputPath := flag.String("output", "", "The output file. (This creates a new file instead of overriding the original file.)")
	ignore := flag.String("ignore", "", "Regex pattern to ignore.")
	obfuscateFor := flag.String("obfuscate-for", "", "Obfuscate for particular programming language. (Example: swift)")

	flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()

	// Load the file contents
	fileBytes, err := os.ReadFile(*sourcePath)
	if err != nil {
		fmt.Println(err)
		return
	}


	var valueEncoder value_encoders.ValueEncoder
	switch strings.ToLower(*obfuscateFor) {
	case "swift":
		valueEncoder = value_encoders.NewSwiftValueEncoder()
			// code to be executed when expression equals value1
	default:
		valueEncoder = &value_encoders.RawValueEncoder{}
			// code to be executed when expression does not match any of the cases
	}


	fileContent := string(fileBytes)
	updatedContent := parser.ReplaceEnvVariables(fileContent, *ignore, valueEncoder)

	// Check if debug flag is true
  if flag.Lookup("debug").Value.String() == "true" {
		fmt.Println(updatedContent)
	}

	if *outputPath == "" {
		outputPath = sourcePath
	}

	// Write the updated content to the output file
	err = os.WriteFile(*outputPath, []byte(updatedContent), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Print a success message
	fmt.Println("Environment variables injected successfully!")
}