package internal

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// CreateProjectStructure generates the project structure based on the given map
func CreateProjectStructure(basePath, projectName string, structure map[string]interface{}) error {
	for name, value := range structure {
		fmt.Println("🚀 ~ file: createProjectStructure.go ~ line 14 ~ funcCreateProjectStructure ~ value : ", value)
		itemPath := filepath.Join(basePath, name)

		switch v := value.(type) {
		case map[string]interface{}:
			// fmt.Println("🚀 ~ file: createProjectStructure.go ~ line 18 ~ funcCreateProjectStructure ~ v : ", v)
			// Create directory
			if err := os.MkdirAll(itemPath, 0755); err != nil {
				return fmt.Errorf("error creating directory %s: %w", itemPath, err)
			}
			// Recursively create subdirectories
			if err := CreateProjectStructure(itemPath, projectName, v); err != nil {
				return err
			}
		case string:
			// Create file with package declaration if in "internal" folder

			if err := createGoFileWithPackage(itemPath); err != nil {
				return err
			}

		case nil:
			// Explicitly create an empty fi
			if err := createEmptyFile(itemPath); err != nil {
				return err
			}
		}
	}
	return nil
}

// createGoFileWithPackage creates a Go file with a package declaration
func createGoFileWithPackage(filePath string) error {
	packageName := filepath.Base(filepath.Dir(filePath)) // Get the folder name as package name

	content := fmt.Sprintf("package %s\n", packageName)
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("error creating file %s: %w", filePath, err)
	}

	log.Printf("Created Go file: %s", filePath)
	return nil
}

// createEmptyFile creates an empty file
func createEmptyFile(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file %s: %w", filePath, err)
	}
	file.Close()
	log.Printf("Created empty file: %s", filePath)
	return nil
}
