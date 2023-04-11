package uploadx

import (
	"encoding/base64"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
)

const AvatarFileField = "avatar"

func UploadFileFromRequest(r *http.Request, maxMemory int64, name, filePath string) (string, error) {
	err := r.ParseMultipartForm(maxMemory) // a total of maxMemory bytes of its file parts are stored in memory
	if err != nil {
		return "", err
	}
	f, header, err := r.FormFile(name)
	if err != nil {
		return "", err
	}
	defer f.Close()

	tempFile, err := os.Create(path.Join(filePath, header.Filename))
	if err != nil {
		return "", err
	}

	defer tempFile.Close()

	io.Copy(tempFile, f)
	return header.Filename, nil
}

func UploadFile(f multipart.File, header *multipart.FileHeader, filePath string) (string, error) {

	tempFile, err := os.Create(path.Join(filePath, header.Filename))
	if err != nil {
		return "", err
	}

	defer tempFile.Close()

	_, _ = io.Copy(tempFile, f)
	return header.Filename, nil
}

func UploadImageByBase64(data string, format string, path string) (string, error) {
	uri := uuid.New().String() + "." + format
	index := strings.Index(data, "base64")
	index += 7 //"data:image/$type};base64,(data stating here)xyz...."
	fileData := data[index:]
	bytes, err := base64.StdEncoding.DecodeString(fileData)
	if err != nil {
		return "", err
	}

	err = os.WriteFile(path+"/"+uri, bytes, 0666)
	if err != nil {
		return "", err
	}
	return "/" + uri, nil
}

func UploadFileByBase64(data string, fileName, path string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	err = os.WriteFile(path+"/"+fileName, bytes, 0666)
	if err != nil {
		return "", err
	}
	return "/" + fileName, nil
}
