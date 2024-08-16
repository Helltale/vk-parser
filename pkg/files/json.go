package files

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"parser/config"
	"path/filepath"
)

func ToJSON(body string, index int) {
	resPath, err := config.GetResultDirectory()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(2)
	}

	filename := filepath.Join(resPath, fmt.Sprintf("response%d.json", index))

	var jsonData map[string]interface{}
	err = json.Unmarshal([]byte(body), &jsonData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	formattedJSON := formatJSON(jsonData)

	err = os.WriteFile(filename, formattedJSON, 0644)
	if err != nil {
		fmt.Printf("error: can not write in file %s", err)
		return
	}

	fmt.Printf("done: created response%d.json\n", index)
}

func formatJSON(data interface{}) []byte {
	switch v := data.(type) {
	case map[string]interface{}:
		var buffer bytes.Buffer
		buffer.WriteString("{\n")
		keys := make([]string, 0, len(v))
		for key := range v {
			keys = append(keys, key)
		}
		for i, key := range keys {
			buffer.WriteString(fmt.Sprintf("  \"%s\": %s", key, formatJSON(v[key])))
			if i < len(keys)-1 {
				buffer.WriteString(",\n")
			} else {
				buffer.WriteString("\n")
			}
		}
		buffer.WriteString("}")
		return buffer.Bytes()
	case []interface{}:
		var buffer bytes.Buffer
		buffer.WriteString("[\n")
		for i, value := range v {
			buffer.WriteString(fmt.Sprintf("  %s", formatJSON(value)))
			if i < len(v)-1 {
				buffer.WriteString(",\n")
			} else {
				buffer.WriteString("\n")
			}
		}
		buffer.WriteString("]")
		return buffer.Bytes()
	case string:
		return []byte(fmt.Sprintf("\"%s\"", v))
	case float64:
		return []byte(fmt.Sprintf("%.f", v))
	case bool:
		return []byte(fmt.Sprintf("%t", v))
	case nil:
		return []byte("null")
	default:
		return []byte(fmt.Sprintf("%v", data))
	}
}

func ToJSON2(jsonString string, index int) error {
	resPath, err := config.GetResultDirectory()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(2)
	}

	filename := filepath.Join(resPath, fmt.Sprintf("response%d.json", index))

	var data interface{}
	err = json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(jsonString))
	if err != nil {
		return err
	}
	fmt.Printf("done: created response%d.json\n", index)
	return nil
}
