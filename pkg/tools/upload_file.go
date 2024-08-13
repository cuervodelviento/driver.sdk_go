package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type uploadFileResponse struct {
	Filename string `json:"filename"`
}

func UploadFileAndGetURL(driverHubHost string, driverKey string, file *os.File) (string, error) {
	// Do something

	body := &bytes.Buffer{}
	host := driverHubHost
	if !strings.HasPrefix(host, "http://") {
		host = fmt.Sprintf("http://%s", host)
	}
	url := fmt.Sprintf("%s/api/v1/upload", host)

	// Create a form file
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		return "", err
	}
	io.Copy(part, file)

	contentType := writer.FormDataContentType()
	writer.Close()
	res, err := http.Post(url, contentType, body)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	content, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var response uploadFileResponse
	if err := json.Unmarshal(content, &response); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/public/%s", driverHubHost, response.Filename), nil
}
