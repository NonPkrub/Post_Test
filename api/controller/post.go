package controller

import (
	"Test/domain"
	"encoding/json"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PostController struct {
	postUseCase domain.PostUseCase
}

func NewPostController(postUseCase domain.PostUseCase) *PostController {
	return &PostController{postUseCase: postUseCase}
}

// @Summary Create a new post
// @Description Create a new post with the provided data
// @Tags Posts
// @Accept json
// @Produce json
// @Param body body domain.PostReq true "Post data"
// @Success 200 {object} domain.PostRes
// @Failure 400 {string} Bad request
// @Router /posts [post]
func (pc *PostController) Create(c *fiber.Ctx) error {
	var form *domain.PostReq
	jsonData, err := json.Marshal(form)
	if err != nil {
		return err
	}

	var newForm domain.PostReq
	err = json.Unmarshal(jsonData, &newForm)
	if err != nil {
		return err
	}

	if err := c.BodyParser(&newForm); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	res, err := pc.postUseCase.Create(&newForm)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	if res.Title == "" {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrBadRequest.Message,
			"status_code": fiber.ErrBadRequest.Code,
			"message":     "title is required",
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "",
		"result":      res,
	})
}

// @Summary Get all posts
// @Description Get all posts with optional pagination
// @Tags Posts
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Success 200 {object} domain.PostResponse
// @Failure 500 {string} Internal server error
// @Router /posts [get]
func (pc *PostController) GetAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))

	pagination := domain.Pagination{
		Page:     page,
		PageSize: pageSize,
	}

	publishedStr := c.Query("published", "")
	var published bool
	if publishedStr == "" {
		published = false
	} else {
		var err error
		published, err = strconv.ParseBool(publishedStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":      fiber.ErrBadRequest.Message,
				"status_code": fiber.ErrBadRequest.Code,
				"message":     "Invalid value for 'published'",
				"result":      nil,
			})
		}

	}

	location, _ := time.LoadLocation("Asia/Bangkok")
	title := c.Query("title", "")
	createdAtStr := c.Query("created_at", "")
	var createdAt time.Time
	if createdAtStr != "" {
		var err error
		createdAt, err = time.ParseInLocation(time.RFC3339, createdAtStr, location)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":      fiber.ErrBadRequest.Message,
				"status_code": fiber.ErrBadRequest.Code,
				"message":     "Invalid value for 'created_at'.",
				"result":      nil,
			})
		}

	}

	reqParams := domain.PostAllReq{
		Published: published,
		Title:     title,
		CreatedAt: createdAt,
	}

	res, err := pc.postUseCase.GetAll(&reqParams, &pagination)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "",
		"result":      res,
	})
}

// @Summary Get a post by ID
// @Description Get a post by its unique identifier
// @Tags Posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} domain.PostRes
// @Failure 404 {string} Resource not found
// @Router /posts/{id} [get]
func (pc *PostController) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	res, err := pc.postUseCase.GetByID(id)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "",
		"result":      res,
	})
}

// @Summary Update a post by ID
// @Description Update a post with new data
// @Tags Posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Param body body domain.PostUpdateReq true "Updated post data"
// @Success 200 {object} domain.PostRes
// @Failure 400 {string} Bad request
// @Failure 404 {string} Resource not found
// @Router /posts/{id} [put]
func (pc *PostController) UpdateByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var form *domain.PostUpdateReq
	jsonData, err := json.Marshal(form)
	if err != nil {
		return err
	}

	var newForm domain.PostUpdateReq
	err = json.Unmarshal(jsonData, &newForm)
	if err != nil {
		return err
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	newForm.ID = uuid

	if err := c.BodyParser(&newForm); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	res, err := pc.postUseCase.UpdateByID(&newForm)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "",
		"result":      res,
	})
}

// @Summary Delete a post by ID
// @Description Delete a post by its unique identifier
// @Tags Posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 204 "No content"
// @Failure 404 {string} Resource not found
// @Router /posts/{id} [delete]
func (pc *PostController) DeleteByID(c *fiber.Ctx) error {
	id := c.Params("id")

	err := pc.postUseCase.DeleteByID(id)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "",
		"result":      nil,
	})
}
