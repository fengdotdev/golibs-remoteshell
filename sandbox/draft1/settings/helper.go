package settings

import (
	"fmt"
	"os"
	"path/filepath"
)

func SaveSettings(settings *Settings) error {
	if settings == nil {
		return nil // No settings to save, nothing to do
	}

	fullpath := "settings.json"
	return SaveSettingsToJson(fullpath, settings)
}

func SaveSettingsToJson(filePath string, settings *Settings) error {
	jsonString, err := settings.ToJSON()
	if err != nil {
		return err
	}

	jsondata := []byte(jsonString)
	if len(jsondata) == 0 {
		return nil // No data to write, nothing to do
	}

	err = WriteToFile(filePath, jsondata)
	if err != nil {
		return err
	}

	return nil
}

func GetSettingsFrom() (*Settings, error) {
	fullpath := "settings.json"
	return GetSettingsFromJson(fullpath)
}

func GetSettingsFromJson(filePath string) (*Settings, error) {
	data, err := ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil // No data to read, return nil
	}
	jsonString := string(data)

	settings := NewZeroSettings()
	err = settings.FromJSON(jsonString)
	if err != nil {
		return nil, err
	}

	return settings, nil
}

func WriteToFile(filePath string, data []byte) error {
	dirpath := filepath.Dir(filePath)
	// create folders if not exists
	err := os.MkdirAll(dirpath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directories for %s: %w", filePath, err)
	}

	// create file if not exists
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filePath, err)
	}
	defer file.Close()

	// write data to file
	if _, err := file.Write(data); err != nil {
		return fmt.Errorf("failed to write data to file %s: %w", filePath, err)
	}

	return nil
}

func ReadFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filePath, err)
	}
	defer file.Close()

	data := make([]byte, 0)
	if _, err := file.Read(data); err != nil {
		return nil, fmt.Errorf("failed to read data from file %s: %w", filePath, err)
	}
	return data, nil
}
