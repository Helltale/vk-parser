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

	filename := filepath.Join(resPath, fnparts[0], "jsons", fmt.Sprintf("%s_response_%s.json", fnparts[0], fnparts[1]))

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

func AllInOneJSON(ids []string) error {
	resPath, err := config.GetResultDirectory()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(2)
	}

	for _, id := range ids {
		allData := make(map[string]interface{})

		for _, filename := range []string{
			filepath.Join(resPath, id, "jsons", fmt.Sprintf("%s_response_user.json", id)),
			filepath.Join(resPath, id, "jsons", fmt.Sprintf("%s_response_video.json", id)),
			filepath.Join(resPath, id, "jsons", fmt.Sprintf("%s_response_wall.json", id)),
			filepath.Join(resPath, id, "jsons", fmt.Sprintf("%s_response_photo-profile.json", id)),
			filepath.Join(resPath, id, "jsons", fmt.Sprintf("%s_response_photo-saved.json", id)),
			filepath.Join(resPath, id, "jsons", fmt.Sprintf("%s_response_photo-wall.json", id))} {
			data, err := readJSONFile(filename)
			if err != nil {
				return err
			}
			allData[filepath.Base(filename)] = data
		}

		err = writeJSONFile(filepath.Join(resPath, id, "jsons", fmt.Sprintf("%s_full.json", id)), allData)
		if err != nil {
			return err
		}

		fmt.Printf("info: created %s\n", fmt.Sprintf("%s_full.json", id))
	}
	return nil
}

func readJSONFile(filename string) (map[string]interface{}, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func writeJSONFile(filename string, data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
