package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joaomarcuscardoso/formValidator/internal/handler"
	"github.com/joaomarcuscardoso/formValidator/internal/handler/form"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializarHandler()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/form", form.GetFormHandler)
		/* v1.POST("/form", handler.CreateFormHandler)
		v1.PUT("/form", handler.UpdateFormHandler)
		v1.DELETE("/form", handler.DeleteFormHandler)
		v1.POST("/tag", handler.CreateTagsHandler) */
	}
}
