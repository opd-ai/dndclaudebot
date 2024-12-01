package main

import (
	"flag"
	"fmt"
	"os"
)

var directory = flag.String("dirname", "01-Adventure", "Name of the output directory for the adventure")
var setting = flag.String("setting", "SETTING.md", "a file containing the details of the campaign setting")

// main.go
func main() {

	flag.Parse()
	config := Config{
		APIKey:     os.Getenv("CLAUDE_API_KEY"),
		OutputDir:  *directory,
		MaxRetries: 3,
	}

	if config.APIKey == "" {
		fmt.Println("Please set CLAUDE_API_KEY environment variable")
		os.Exit(1)
	}

	client := NewClaudeClient(config.APIKey)

	prompt := os.Args[1]
	if prompt == "" {
		fmt.Println("Please provide a narrative prompt")
		os.Exit(1)
	}

	// Process the adventure
	adventure, err := generateTableOfContents(client, prompt)
	if err != nil {
		fmt.Printf("Error generating table of contents: %v\n", err)
		os.Exit(1)
	}

	if err := generateCoverPrompts(client, &adventure); err != nil {
		fmt.Printf("Error generating cover pages %v\n", err)
		os.Exit(1)
	}

	if err := generateOnePageDungeons(client, &adventure); err != nil {
		fmt.Printf("Error generating one-page dungeons: %v\n", err)
		os.Exit(1)
	}

	if err := expandAdventures(client, &adventure); err != nil {
		fmt.Printf("Error expanding adventures: %v\n", err)
		os.Exit(1)
	}

	if err := generateIllustrationPrompts(client, &adventure); err != nil {
		fmt.Printf("Error generating illustration prompts: %v\n", err)
		os.Exit(1)
	}

	if err := removeCopyrightedMaterial(client, &adventure); err != nil {
		fmt.Printf("Error removing copyrighted material: %v\n", err)
		os.Exit(1)
	}

	if err := saveToFiles(&adventure, config.OutputDir); err != nil {
		fmt.Printf("Error saving files: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Adventure generation complete!")
}
