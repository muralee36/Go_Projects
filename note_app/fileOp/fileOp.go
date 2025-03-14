package fileop

import (
	"encoding/json"
	"io"
	"os"
)

func ReadJsonFile(filePath string) ([]map[string]interface{}, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	jsonByte, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var jsonData []map[string]interface{}
	err = json.Unmarshal(jsonByte, &jsonData)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func WriteJsonFile(filePath string, data []map[string]interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return err
	}
return nil
}
