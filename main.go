package main

import (
	"encoding/json"
	"fmt"
	"generate_project_structure/internal"
	"log"
	"os"
	"path/filepath"

	_ "embed"

	"github.com/spf13/cobra"
)

//go:embed project_structure.json
var projectStructureJSON []byte

func main() {
	var targetDir string
	var projectName string

	// Define the root command for cobra
	var rootCmd = &cobra.Command{
		Use:   "go3D",
		Short: "A tool to generate Go files with a package declaration",
		Run: func(cmd *cobra.Command, args []string) {
			// Check if project name is provided
			if len(args) < 1 {
				log.Fatalf("Please provide a project name. Usage: go run main.go <project_name>")
			}

			// Get the project name from args
			projectName = args[0]

			// Set the target directory if not set
			if targetDir == "" {
				targetDir, _ = os.Getwd()
				fmt.Println("🚀 Target Directory:", targetDir)
			}

			// Parse the embedded JSON
			var projectStructure map[string]interface{}
			if err := json.Unmarshal(projectStructureJSON, &projectStructure); err != nil {
				log.Fatalf("Error parsing JSON: %v", err)
			}
			jsonPretty, err := json.MarshalIndent(projectStructure, "", "  ")
			if err != nil {
				log.Fatalf("Error formatting JSON: %v", err)
			}
			fmt.Println("🚀 Parsed JSON Structure:\n", string(jsonPretty))

			// Get the root structure
			rootStructure, ok := projectStructure["root"].(map[string]interface{})
			if !ok {
				log.Fatalf("Invalid JSON structure")
			}

			// Specify the base path for project generation
			basePath := filepath.Join(targetDir, projectName)
			// Create the base directory
			if err := os.MkdirAll(basePath, 0755); err != nil {
				log.Fatalf("Error creating base directory: %v", err)
			}

			// Generate the project structure
			if err := internal.CreateProjectStructure(basePath, projectName, rootStructure); err != nil {
				log.Fatalf("Error generating project structure: %v", err)
			}

			fmt.Println("✅ Project structure generated successfully in", basePath)
		},
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
