package routes

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/handlers"
	"skripsi.id/obfuss/middlewares"
)

type TodoRoute struct {
	handler *handlers.TodoHandler
	prefix  string
}

func NewTodoRoute(db *gorm.DB, logger *zap.Logger) Route {
	return &TodoRoute{
		handler: handlers.NewTodoHandler(db, logger),
		prefix:  "/todo",
	}
}

func (r *TodoRoute) Register(app *gin.RouterGroup) {
	api := app.Group(r.prefix, middlewares.RequestId())
	api.GET("/", r.handler.GetTodos)
	api.GET("/:id", r.handler.GetTodoByID)
	api.POST("/", r.handler.Create)
	api.PUT("/:id", r.handler.Update)
	api.DELETE("/:id", r.handler.Delete)
}
