package controllers

import (
	"Project-Management-System/models"
	"Project-Management-System/services"
	"github.com/gofiber/fiber/v2"
	"time"
)

func CreateComment(c *fiber.Ctx) error {
	var comment models.CommentModel
	if err := c.BodyParser(&comment); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	taskId := c.Params("task_id")
	comment.TaskId = taskId
	comment.CreatedAt = time.Now()

	createdComment, err := services.CreateComment(comment)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(createdComment)
}

func GetComments(c *fiber.Ctx) error {
	taskId := c.Params("task_id")
	comments, err := services.GetCommentsByTaskId(taskId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(comments)
}

func UpdateComment(c *fiber.Ctx) error {
	commentId := c.Params("comment_id")
	var updatedComment models.CommentModel

	if err := c.BodyParser(&updatedComment); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err := services.UpdateComment(commentId, updatedComment)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("comment not found")
	}

	return c.Status(fiber.StatusOK).JSON(updatedComment)
}

func DeleteComment(c *fiber.Ctx) error {
	commentId := c.Params("comment_id")
	err := services.DeleteComment(commentId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("comment not found")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.StatusNoContent)
}
