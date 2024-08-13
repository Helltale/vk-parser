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

	// Преобразуем JSON-строку в объект
	var jsonData map[string]interface{}
	err = json.Unmarshal([]byte(body), &jsonData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Форматируем JSON с отступами, но сохраняем ссылки в оригинальном формате
	formattedJSON := formatJSON(jsonData)

	// Записываем форматированный JSON в файл
	err = os.WriteFile(filename, formattedJSON, 0644)
	if err != nil {
		fmt.Println("error: can not write in file", err)
		return
	}

	fmt.Printf("done: created response%d.json\n", index)
}

func formatJSON(data interface{}) []byte {
	switch v := data.(type) {
	case map[string]interface{}:
		var buffer bytes.Buffer
		buffer.WriteString("{\n")
		for key, value := range v {
			buffer.WriteString(fmt.Sprintf("  \"%s\": %s,\n", key, formatJSON(value)))
		}
		buffer.WriteString("}")
		return buffer.Bytes()
	case []interface{}:
		var buffer bytes.Buffer
		buffer.WriteString("[\n")
		for _, value := range v {
			buffer.WriteString(fmt.Sprintf("  %s,\n", formatJSON(value)))
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
