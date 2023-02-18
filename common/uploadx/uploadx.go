package uploadx

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
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
