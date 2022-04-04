package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iman_task/api-gateway/api/handlers/handler"
	"github.com/iman_task/api-gateway/api/handlers/post"
	configPkg "github.com/iman_task/api-gateway/config"
	"github.com/iman_task/api-gateway/pkg/logger"
	"github.com/iman_task/api-gateway/services"
	"net/http"
)

//Option ...
type Option struct {
	Conf           configPkg.Config
	Logger         logger.Logger
	ServiceManager services.ServiceManager
}

//New ...
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Static("/swagger/", "swagger")

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.UseRawPath = true

	var handlerV1 = handler.New(&handler.HandlerV1{
		PostHandler: *post.New(&post.PostHandlerConfig{
			Conf:           option.Conf,
			Logger:         option.Logger,
			ServiceManager: option.ServiceManager,
		}),
	})

	// endpoints without version
	{
		router.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "Welcome to Posts CRUD API GATEWAY!")
		})
	}

	api := router.Group("/v1")
	{
		api.GET("/posts/list/", handlerV1.PostHandler.ListPosts)
		api.GET("/posts/detail/:id/", handlerV1.PostHandler.DetailPost)
		api.PUT("/posts/update/:id/", handlerV1.PostHandler.UpdatePost)
		api.DELETE("/posts/delete/:id/", handlerV1.PostHandler.DeletePost)
		api.POST("/posts/collect/", handlerV1.PostHandler.CollectPosts)
		api.GET("/posts/collect/check/", handlerV1.PostHandler.CheckCollectStatus)
	}

	return router
}
