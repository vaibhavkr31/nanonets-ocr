package nanonets

import (
	"bytes"
	"fmt"
	"github.com/dripcapital/nanonets-ocr/app/config"
	"github.com/dripcapital/nanonets-ocr/app/utils"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

var logger *zap.Logger

func init() {
	logger = config.InitLogger("nanonets")
}

func PredictViaImageFile(filePath string) (string, error) {
	modelId := viper.GetString("blDocModelId")
	url := fmt.Sprintf("https://app.nanonets.com/api/v2/OCR/Model/%s/LabelFile/", modelId)

	file, err := os.Open(filePath)
	if err != nil {
		logger.Error(fmt.Sprintf("Error On Opening the File : %s", err))
		return "", err
	}

	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		logger.Error(fmt.Sprintf("Error While Creating Form File: %s", err))
		return "", err
	}
	_, err = io.Copy(part, file)

	contentType := writer.FormDataContentType()

	err = writer.Close()
	if err != nil {
		fmt.Println(err)
	}

	req, _ := http.NewRequest("POST", url, body)

	req.Header.Add("Content-Type", contentType)
	userName := viper.GetString("userName")
	password := viper.GetString("password")
	req.SetBasicAuth(userName, password)

	res, _ := http.DefaultClient.Do(req)

	respBody, _ := ioutil.ReadAll(res.Body)

	logger.Debug("Deleting the File from Local")

	utils.DeleteFileLocally(filePath)

	logger.Debug("Deleted the File Successfully")

	return string(respBody), nil
}
