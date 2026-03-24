package coreengine

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ReadJSONFile reads a JSON file and unmarshals it into the provided struct.
func ReadJSONFile(filePath string, v interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

// WriteJSONFile marshals the provided struct into JSON and writes it to a file.
func WriteJSONFile(filePath string, v interface{}, indent bool) error {
	var data []byte
	var err error

	if indent {
		data, err = json.MarshalIndent(v, "", "  ")
	} else {
		data, err = json.Marshal(v)
	}

	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, data, 0644)
}

// EnsureDir checks if a directory exists and creates it if it doesn't.
func EnsureDir(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return os.MkdirAll(dirPath, 0755)
	}
	return nil
}

// FileExists checks if a file exists and is not a directory.
func FileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// GetFileExtension returns the file extension from a file path.
func GetFileExtension(filePath string) (string, error) {
	ext := filepath.Ext(filePath)
	if ext == "" {
		return "", errors.New("no file extension found")
	}
	return ext, nil
}