package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	ProjectName      string                 `yaml:"project_name"`
	ProjectType      string                 `yaml:"project_type"`
	ProjectPath      string                 `yaml:"project_path"`
	ProjectStructure map[string]interface{} `yaml:"project_structure"`
}

func createStructure(basePath string, structure map[string]interface{}) error {
	for folder, content := range structure {
		if strings.HasPrefix(folder, "file:") {
			// This is a file, not a folder
			fileName := strings.TrimPrefix(folder, "file:")
			filePath := filepath.Join(basePath, fileName)
			_, err := os.Create(filePath)
			if err != nil {
				return err
			}
		} else {
			// This is a folder
			newPath := filepath.Join(basePath, folder)
			if err := os.MkdirAll(newPath, 0755); err != nil {
				return err
			}
			if subfolders, ok := content.(map[interface{}]interface{}); ok {
				newStructure := make(map[string]interface{})
				for k, v := range subfolders {
					newStructure[k.(string)] = v
				}
				if err := createStructure(newPath, newStructure); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func main() {
	configPath := flag.String("config_path", "", "Path to the configuration file")
	flag.Parse()

	if *configPath == "" {
		fmt.Println("Please provide a config_path.")
		return
	}

	data, err := ioutil.ReadFile(*configPath)
	if err != nil {
		fmt.Println("Error reading YAML file:", err)
		return
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error parsing YAML:", err)
		return
	}

	fullPath := filepath.Join(config.ProjectPath, config.ProjectName)
	if err := os.MkdirAll(fullPath, 0755); err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	if err := createStructure(fullPath, config.ProjectStructure); err != nil {
		fmt.Println("Error creating project structure:", err)
		return
	}

	fmt.Printf("Successfully created project structure at %s\n", fullPath)
}
