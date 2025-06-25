package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		log.Printf("Error parsing multipart form: %v", err)
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		log.Printf("Error retrieving file from form: %v", err)
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Error reading file content: %v", err)
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}
	fileContent := string(fileBytes)

	convertedData, err := service.AutoDetectAndConvert(fileContent)
	if err != nil {
		log.Printf("Error during conversion: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	timestamp := time.Now().UTC().Format("20060102_150405")
	originalExt := filepath.Ext(handler.Filename)
	newFileName := fmt.Sprintf("result_%s%s", timestamp, originalExt)

	err = os.WriteFile(newFileName, []byte(convertedData), 0644)
	if err != nil {
		log.Printf("Error writing result file: %v", err)
		http.Error(w, "Could not save result file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Файл успешно обработан и сохранен как %s.\n\nРезультат конвертации:\n\n%s", newFileName, convertedData)
}
