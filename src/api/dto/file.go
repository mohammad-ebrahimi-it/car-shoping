package dto

import "mime/multipart"

type FileFormRequest struct {
	File *multipart.FileHeader `form:"file" binding:"required" swaggerignore:"true"`
}

type UploadFileRequest struct {
	FileFormRequest
	Description string `form:"description" json:"description"`
}

type CreateFileRequest struct {
	Name        string `json:"name"`
	Directory   string `json:"directory"`
	Description string `json:"description"`
	MimeType    string `json:"mimeType"`
}

type UpdateFileRequest struct {
	Description string `json:"description"`
}

type FileResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Directory   string `json:"directory"`
	Description string `json:"description"`
	MimeType    string `json:"mimeType"`
}
