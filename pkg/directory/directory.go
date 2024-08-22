package directory

import (
	"fmt"
	"os"
	"parser/config"
	"path/filepath"
)

func CreateFullPath(ids []string) {
	resPath, err := config.GetResultDirectory()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(2)
	}

	for _, id := range ids {
		err = os.MkdirAll(filepath.Join(resPath, id, "jsons"), 0777)
		if err != nil {
			panic(err)
		}

		err = os.MkdirAll(filepath.Join(resPath, id, "photos", "wall"), 0777)
		if err != nil {
			panic(err)
		}

		err = os.MkdirAll(filepath.Join(resPath, id, "photos", "saved"), 0777)
		if err != nil {
			panic(err)
		}

		err = os.MkdirAll(filepath.Join(resPath, id, "photos", "profile"), 0777)
		if err != nil {
			panic(err)
		}

		err = os.MkdirAll(filepath.Join(resPath, id, "video"), 0777)
		if err != nil {
			panic(err)
		}

		fmt.Printf("info: created directories for %s\n", id)
	}

}
