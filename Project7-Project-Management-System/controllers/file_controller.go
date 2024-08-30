package controllers

import (
	"Project-Management-System/models"
	"Project-Management-System/services"
	"github.com/gofiber/fiber/v2"
	"os"
	"path/filepath"
	"time"
)

func UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	taskId := c.Params("task_id")
	filePath := filepath.Join("uploads", file.Filename)

	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	fileModel := models.FileModel{
		FileName:   file.Filename,
		FilePath:   filePath,
		TaskId:     taskId,
		UploadedBy: 1,
		UploadedAt: time.Now(),
	}
	_, err = services.SaveFile(fileModel)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fileModel)
}

func GetFiles(c *fiber.Ctx) error {
	id := c.Query("task_id")
	files, err := services.GetFiles(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(files)
}

func DownloadFile(c *fiber.Ctx) error {
	id := c.Params("file_id")
	file, err := services.GetFileByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("File not found")
	}

	filePath := file.FilePath
	return c.SendFile(filePath)
}

func DeleteFile(c *fiber.Ctx) error {
	id := c.Params("file_id")
	file, err := services.GetFileById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("File not found")
	}

	if err := os.Remove(file.FilePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if err := services.DeleteFile(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}
