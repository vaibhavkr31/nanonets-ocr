package utils

import (
	"fmt"
	"github.com/dripcapital/nanonets-ocr/app/config"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"io"
	"net/http"
	"os"
)

var logger *zap.Logger

func init() {
	logger = config.InitLogger("utils")
}

func SaveImageLocally(url string) string {
	// don't worry about errors
	logger.Debug("Starting to Save Image in Local")
	response, err := http.Get(url)
	if err != nil {
		logger.Error(fmt.Sprintf("Error While Fetching the Image From Url : %s", err))
	}
	defer response.Body.Close()

	//open a file for writing
	fileId := uuid.New()
	filePath := fmt.Sprintf("./demo/%s.pdf", fileId.String())

	file, err := os.Create(filePath)
	if err != nil {
		logger.Error(fmt.Sprintf("Error While Saving the Image Locally : %s", err))
	}
	defer file.Close()

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		logger.Error(fmt.Sprintf("Error Returned While Cpying Response in File : %s", err))
	}
	logger.Debug(fmt.Sprintf("Saved the File Locally with filePath: %s", filePath))
	return filePath
}

func DeleteFileLocally(filePath string) {
	e := os.Remove(filePath)
	if e != nil {
		fmt.Println("Error While Deleting a File")
	}
}
