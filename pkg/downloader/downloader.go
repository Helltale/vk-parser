package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func download(url, filename string) error {
	// Создаем HTTP-клиент
	client := &http.Client{}

	// Отправляем HTTP-запрос
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Определяем тип файла
	contentType := resp.Header.Get("Content-Type")
	fileExtension := getFileExtension(contentType)

	// Создаем файл для сохранения
	file, err := os.Create(fmt.Sprintf("%s.%s", filename, fileExtension))
	if err != nil {
		return err
	}
	defer file.Close()

	// Копируем содержимое ответа в файл
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Downloaded %s.%s\n", filename, fileExtension)
	return nil
}

func getFileExtension(contentType string) string {
	switch contentType {
	case "image/jpeg":
		return "jpg"
	case "image/png":
		return "png"
	case "video/mp4":
		return "mp4"
	case "video/webm":
		return "webm"
	default:
		return filepath.Ext(contentType)[1:]
	}
}
