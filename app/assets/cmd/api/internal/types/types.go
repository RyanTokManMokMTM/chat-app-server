// Code generated by goctl. DO NOT EDIT.
package types

type HealthCheckResp struct {
	Resp string `json:"resp"`
}

type UploadFileReq struct {
}

type UploadFileResp struct {
	Code uint   `json:"code"`
	Path string `json:"path"`
}

type UploadImageReq struct {
	ImageType string `json:"image_type"`
	Data      string `json:"data"`
}

type UploadImageResp struct {
	Code uint   `json:"code"`
	Path string `json:"path"`
}

type UploadImagesReq struct {
}

type UploadImagesResp struct {
}