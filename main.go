package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"michaelhenry/envject/parser"
)


func main() {
	// Define command-line flags
	sourcePath := flag.String("file", "", "File to inject the environment variables")
	outputPath := flag.String("output", "", "The output file. (This creates a new file instead of overriding the original file.)")
	flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()

	// Load the file contents
	fileBytes, err := ioutil.ReadFile(*sourcePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileContent := string(fileBytes)
	updatedContent := parser.ReplaceEnvVariables(fileContent)

	// Check if debug flag is true
  if flag.Lookup("debug").Value.String() == "true" {
		fmt.Println(updatedContent)
	}

	if *outputPath == "" {
		outputPath = sourcePath
	}

	// Write the updated content to the same file
	err = ioutil.WriteFile(*outputPath, []byte(updatedContent), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Print a success message
	fmt.Println("Environment variables injected successfully!")
}