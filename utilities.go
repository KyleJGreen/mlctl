package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/template"
)

// Configuration represents the configuration loaded from the JSON file.
type Configuration struct {
	WorkflowNamePrefix      string `json:"WorkflowNamePrefix"`
	DataLoaderImage         string `json:"DataLoaderImage"`
	DataPreprocessingImage  string `json:"DataPreprocessingImage"`
	ModelTrainingImage      string `json:"ModelTrainingImage"`
	ModelEvaluationImage    string `json:"ModelEvaluationImage"`
	ModelDeploymentImage    string `json:"ModelDeploymentImage"`
	DataPostprocessingImage string `json:"DataPostprocessingImage"`
}

func LoadConfiguration(filename string) (*Configuration, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Configuration
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func TemplateFiles(templateDir, targetDir string, config *Configuration) error {
	files, err := os.ReadDir(templateDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".tmpl") {
			continue
		}

		templatePath := fmt.Sprintf("%s/%s", templateDir, file.Name())
		targetPath := fmt.Sprintf("%s/%s", targetDir, strings.TrimSuffix(file.Name(), ".tmpl"))

		if err := TemplateFile(templatePath, targetPath, config); err != nil {
			return err
		}
	}

	return nil
}

func TemplateFile(templatePath, targetPath string, config *Configuration) error {
	data, err := os.ReadFile(templatePath)
	if err != nil {
		return err
	}

	tmpl, err := template.New("").Parse(string(data))
	if err != nil {
		return err
	}

	targetFile, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer targetFile.Close()

	return tmpl.Execute(targetFile, config)
}
