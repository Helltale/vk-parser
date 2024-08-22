package files

import (
	"encoding/json"
	"fmt"
	"os"
	"parser/config"
	"path/filepath"
)

func ToJSON2(jsonString string, fnparts []string) error {
	resPath, err := config.GetResultDirectory()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(2)
	}

	filename := filepath.Join(resPath, fmt.Sprintf("%s_response_%s.json", fnparts[0], fnparts[1]))

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
	fmt.Printf("info: created %s_response_%s.json\n", fnparts[0], fnparts[1])
	return nil
}
