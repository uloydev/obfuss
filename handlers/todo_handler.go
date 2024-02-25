package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/entities"
	"skripsi.id/obfuss/models"
	"skripsi.id/obfuss/services"
)

type TodoHandler struct {
	db      *gorm.DB
	logger  *zap.Logger
	service *services.TodoService
}

func NewTodoHandler(db *gorm.DB, logger *zap.Logger) *TodoHandler {
	return &TodoHandler{
		db:      db,
		logger:  logger.With(zap.String("type", "TodoHandler")),
		service: services.NewTodoService(db, logger),
	}
}

//	@Summary		Get Todos
//	@Description	Get Todos
//	@Tags			Todo
//	@Accept			json
//	@Produce		json
//	@Param			params	query		models.PaginationParams	true	"Pagination parameters"
//	@Success		200		{object}	models.BaseResponse[[]entities.Todo]
//	@Router			/todo [get]
func (h *TodoHandler) GetTodos(c *gin.Context) {
	var params models.PaginationParams
	if err := c.BindQuery(&params); err != nil {
		c.JSON(400, models.BaseResponse[entities.Todo]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	todos, meta, err := h.service.GetTodos(params)
	if err != nil {
		c.JSON(500, models.BaseResponse[entities.Todo]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	c.JSON(200, models.BaseResponse[[]entities.Todo]{
		Message: "success",
		Data:    todos,
		Meta:    meta,
	})
}

//	@Summary		Get Todo By ID
//	@Description	Get Todo By ID
//	@Tags			Todo
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Todo ID"
//	@Success		200	{object}	models.BaseResponse[entities.Todo]
//	@Router			/todo/{id} [get]
func (h *TodoHandler) GetTodoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, models.BaseResponse[entities.Todo]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	todo, err := h.service.GetTodoByID(uint(id))
	if err != nil {
		c.JSON(500, models.BaseResponse[entities.Todo]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	c.JSON(200, models.BaseResponse[entities.Todo]{
		Message: "success",
		Data:    *todo,
	})
}

//	@Summary		Create Todo
//	@Description	Create Todo
//	@Tags			Todo
//	@Accept			json
//	@Produce		json
//	@Param			todo	body		models.TodoRequest	true	"Todo"
//	@Success		200		{object}	models.BaseResponse[entities.Todo]
//	@Router			/todo [post]
func (h *TodoHandler) Create(c *gin.Context) {
	var req models.TodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(500, models.BaseResponse[entities.Todo]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	todo := entities.Todo{
		Action: req.Action,
	}

	if err := h.service.Create(&todo); err != nil {
		c.JSON(500, models.BaseResponse[entities.Todo]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	c.JSON(200, models.BaseResponse[entities.Todo]{
		Message: "success",
		Data:    todo,
	})
}

//	@Summary		Update Todo
//	@Description	Update Todo
//	@Tags			Todo
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Todo ID"
//	@Param			todo	body		models.TodoRequest	true	"Todo"
//	@Success		200		{object}	models.BaseResponse[entities.Todo]
//	@Router			/todo/{id} [put]
func (h *TodoHandler) Update(c *gin.Context) {
	var req models.TodoRequest

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, models.BaseResponse[entities.Todo]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(500, models.BaseResponse[entities.Todo]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	todo := entities.Todo{
		ID:     uint(id),
		Action: req.Action,
	}

	if err := h.service.Update(&todo); err != nil {
		c.JSON(500, models.BaseResponse[entities.Todo]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	c.JSON(200, models.BaseResponse[entities.Todo]{
		Message: "success",
		Data:    todo,
	})
}

//	@Summary		Delete Todo
//	@Description	Delete Todo
//	@Tags			Todo
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Todo ID"
//	@Success		200	{object}	models.BaseResponse[entities.Todo]
//	@Router			/todo/{id} [delete]
func (h *TodoHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, models.BaseResponse[entities.Todo]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(500, models.BaseResponse[entities.Todo]{
			Message: "error",
			Errors:  []any{err.Error()},
		})
		return
	}

	c.JSON(200, models.BaseResponse[entities.Todo]{
		Message: "success",
	})
}
