package uploadx

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"os"
	"path"
	"strings"
)

const AvatarFileField = "avatar"
const CoverFileField = "cover"
const StoryMediaField = "story_media"

//func UploadFileFromRequest(r *http.Request, maxMemory int64, name, filePath string) (string, error) {
//	err := r.ParseMultipartForm(maxMemory) // a total of maxMemory bytes of its file parts are stored in memory
//	if err != nil {
//		return "", err
//	}
//	f, header, err := r.FormFile(name)
//	if err != nil {
//		return "", err
//	}
//	defer f.Close()
//	fileName := strings.ToLower(header.Filename)
//	tempFile, err := os.Create(path.Join(filePath, fileName))
//	if err != nil {
//		return "", err
//	}
//
//	defer tempFile.Close()
//
//	io.Copy(tempFile, f)
//	return "/" + fileName, nil
//}
//
//func UploadFileWithCustomName(f multipart.File, header *multipart.FileHeader, fileName, filePath string) (string, error) {
//	fileType := strings.Split(header.Filename, ".")[1]
//	name := fmt.Sprintf("%s", fileName, fileType)
//	tempFile, err := os.Create(path.Join(filePath, name))
//	if err != nil {
//		return "", err
//	}
//
//	defer tempFile.Close()
//
//	_, _ = io.Copy(tempFile, f)
//	return "/" + name, nil
//}

func RandomFileName(originalName string) string {
	fileType := strings.Split(originalName, ".")
	if len(fileType) < 2 {
		return ""
	}
	randomUUID := strings.ToLower(uuid.NewString())
	name := fmt.Sprintf("%s.%s", randomUUID, fileType[1])
	return name
}

func SaveFileWithRandomName(data []byte, fileName, filePath string) (string, error) {
	fileType := strings.Split(fileName, ".")[1]
	randomUUID := strings.ToLower(uuid.NewString())
	name := fmt.Sprintf("%s.%s", randomUUID, fileType)
	tempFile, err := os.Create(path.Join(filePath, name))
	if err != nil {
		return "", err
	}

	defer tempFile.Close()
	buffer := bytes.NewBuffer(data)

	_, _ = io.Copy(tempFile, buffer)
	return "/" + name, nil
}

func SaveFileWithName(data []byte, name, filePath string) (string, error) {
	tempFile, err := os.Create(path.Join(filePath, name))
	if err != nil {
		return "", err
	}

	defer tempFile.Close()
	buffer := bytes.NewBuffer(data)

	_, _ = io.Copy(tempFile, buffer)
	return "/" + name, nil
}

//
//func UploadFile(f multipart.File, header *multipart.FileHeader, filePath string) (string, error) {
//	fileName := strings.ToLower(header.Filename)
//	tempFile, err := os.Create(path.Join(filePath, header.Filename))
//	if err != nil {
//		return "", err
//	}
//
//	defer tempFile.Close()
//
//	_, _ = io.Copy(tempFile, f)
//
//	return "/" + fileName, nil
//}

func SaveBytesIntoFile(fileName string, bytes []byte, filePath string) (string, error) {
	fileName = strings.ToLower(fileName)
	//logx.Info("path : ", path.Join(util.GetRootDir(), filePath, fileName))
	tempFile, err := os.Create(path.Join(filePath, fileName))
	if err != nil {
		return "", err
	}

	defer tempFile.Close()

	if _, err = tempFile.Write(bytes); err != nil {
		return "", err
	}

	return "/" + fileName, nil
}

//
//func UploadFileWithCustome(f multipart.File, header *multipart.FileHeader, filePath string) (string, error) {
//
//	tempFile, err := os.Create(path.Join(filePath, header.Filename))
//	if err != nil {
//		return "", err
//	}
//
//	defer tempFile.Close()
//
//	_, _ = io.Copy(tempFile, f)
//	return "/" + header.Filename, nil
//}

func SaveImageByBase64(data string, format string, resourcesPath string) (string, error) {

	uri := strings.ToLower(uuid.New().String()) + "." + format
	index := strings.Index(data, "base64")
	index += 7 //"data:image/$type;base64,(data stating here)xyz...."
	fileData := data[index:]
	b, err := base64.StdEncoding.DecodeString(fileData)
	if err != nil {
		return "", err
	}

	pathDir := path.Join(resourcesPath, uri)
	logx.Info("Path : ", pathDir)
	tempFile, err := os.Create(pathDir)
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	buffer := bytes.NewBuffer(b)

	_, _ = io.Copy(tempFile, buffer)
	return "/" + uri, nil
}

//func UploadImageByBase64(data string, format string, path string) (string, error) {
//	uri := strings.ToLower(uuid.New().String()) + "." + format
//	index := strings.Index(data, "base64")
//	index += 7 //"data:image/$type};base64,(data stating here)xyz...."
//	fileData := data[index:]
//	bytes, err := base64.StdEncoding.DecodeString(fileData)
//	if err != nil {
//		return "", err
//	}
//
//	err = os.WriteFile(path+"/"+uri, bytes, 0666)
//	if err != nil {
//		return "", err
//	}
//	return "/" + uri, nil
//}

//func UploadFileByBase64(data string, fileName, path string) (string, error) {
//	bytes, err := base64.StdEncoding.DecodeString(data)
//	if err != nil {
//		return "", err
//	}
//
//	err = os.WriteFile(path+"/"+fileName, bytes, 0666)
//	if err != nil {
//		return "", err
//	}
//	return "/" + fileName, nil
//}
