package handlers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mohammad-ebrahimi-it/car-shoping/api/dto"
	"github.com/mohammad-ebrahimi-it/car-shoping/api/helper"
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
	"github.com/mohammad-ebrahimi-it/car-shoping/pkg/logging"
	"github.com/mohammad-ebrahimi-it/car-shoping/services"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type FileHandler struct {
	service *services.FileService
}

func NewFileHandler(cfg *config.Config) *FileHandler {
	return &FileHandler{
		service: services.NewFileService(cfg),
	}
}

// Create godoc
// @Summary Create a File
// @Description Create a File
// @Tags Files
// @Accept x-www-form-urlencoded
// @Produce json
// @Param file formData dto.UploadFileRequest true "Create a file"
// @Param file formData file true "Create a file"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.FileResponse} "File response"
// @Failure 400 {object} helper.BaseHttpResponse "Failure"
// @Router /v1/files/ [post]
// @Security AuthBearer
func (h *FileHandler) Create(c *gin.Context) {
	upload := dto.UploadFileRequest{}

	err := c.ShouldBind(&upload)

	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.ValidationError, err))
		return
	}

	req := dto.CreateFileRequest{}

	req.Description = upload.Description
	req.MimeType = upload.File.Header.Get("Content-Type")
	req.Directory = "upload"

	req.Name, err = saveUploadFile(upload.File, req.Directory)

	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.ValidationError, err))
		return
	}

	res, err := h.service.Create(c, &req)

	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, true, helper.Success))

}

// UpdateFile godoc
// @Summary Update a File
// @Description Update a File
// @Tags Files
// @Accept json
// @Produce json
// @Param id  path int true "Id"
// @Param Request body dto.UpdateFileRequest true "Update a File"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.FileResponse} "File response"
// @Failure 400 {object} helper.BaseHttpResponse "Failure"
// @Router /v1/files/{id} [put]
// @Security AuthBearer
func (h *FileHandler) UpdateFile(c *gin.Context) {
	Update(c, h.service.Update)
}

// DeleteFile godoc
// @Summary Delete a File
// @Description Delete a File
// @Tags Files
// @Accept json
// @Produce json
// @Param id  path int true "Id"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failure"
// @Router /v1/files/{id} [delete]
// @Security AuthBearer
func (h *FileHandler) DeleteFile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, helper.GenerateBaseResponseWithError(nil, false, helper.NotFoundError, errors.New("id not found")))

		return
	}

	file, err := h.service.GetByID(c, id)

	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err),
		)
		return
	}

	err = os.Remove(fmt.Sprintf("%s/%s", file.Directory, file.Name))
	if err != nil {
		logger.Error(logging.IO, logging.RemoveFile, err.Error(), nil)
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err),
		)
		return
	}

	err = h.service.Delete(c, id)

	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err),
		)
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(nil, true, helper.Success))
	return
}

// GetFileById godoc
// @Summary Get a File
// @Description Get a File
// @Tags Files
// @Accept json
// @Produce json
// @Param id  path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.FileResponse} "File response"
// @Failure 400 {object} helper.BaseHttpResponse "Failure"
// @Router /v1/files/{id} [get]
// @Security AuthBearer
func (h *FileHandler) GetFileById(c *gin.Context) {
	GetById(c, h.service.GetByID)
}

// GetByFilter godoc
// @Summary Get Files
// @Description Get Files
// @Tags Files
// @Accept json
// @Produce json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PageList[dto.FileResponse]} "File response"
// @Failure 400 {object} helper.BaseHttpResponse "Failure"
// @Router /v1/files/get-by-filter [post]
// @Security AuthBearer
func (h *FileHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}

func saveUploadFile(file *multipart.FileHeader, directory string) (string, error) {
	randFileName := uuid.New()

	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return "", err
	}

	fileName := file.Filename
	fileNameArr := strings.Split(fileName, ".")
	fileExt := fileNameArr[len(fileNameArr)-1]

	fileName = fmt.Sprintf("%s.%s", randFileName, fileExt)

	dst := fmt.Sprintf("%s/%s", directory, fileName)

	src, err := file.Open()

	if err != nil {
		return "", err
	}

	defer src.Close()

	out, err := os.Create(dst)

	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
