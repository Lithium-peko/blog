package service

import (
	"blog/utils/r"
	"blog/utils/upload"
	"mime/multipart"
)

type File struct{}

// 文件上传
func (*File) UploadFile(header *multipart.FileHeader) (url string, code int) {
	oss := upload.NewOSS()
	filePath, _, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		return "", r.ERROR_FILE_UPLOAD
	}
	return filePath, r.OK
}
