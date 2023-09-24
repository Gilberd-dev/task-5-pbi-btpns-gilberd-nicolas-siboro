package router

import (
	"github.com/Gilberd-dev/task-5-pbi-btpns-gilberd-nicolas-siboro/controllers"
	"github.com/Gilberd-dev/task-5-pbi-btpns-gilberd-nicolas-siboro/database"
	"github.com/Gilberd-dev/task-5-pbi-btpns-gilberd-nicolas-siboro/middlewares"
	"github.com/gin-gonic/gin"
)

func RouteInit() *gin.Engine {
	route := gin.Default()
	route.Static("/images", "./static/images")

	db := database.GetDB()

	userController := controllers.NewUserController(db)
	photoController := controllers.NewPhotoController(db)

	api := route.Group("/api/v1")

	userRoute := api.Group("/users")
	{
		userRoute.POST("/register", userController.Register)
		userRoute.POST("/login", userController.Login)
		userRoute.PUT("/:userId", userController.Update)
		userRoute.DELETE("/:userId", userController.Delete)
	}

	photoRoute := api.Group("/photo")
	{
		photoRoute.GET("/", middlewares.AuthMiddleware(db), photoController.Get)
		photoRoute.POST("/", middlewares.AuthMiddleware(db), photoController.Create)
		photoRoute.PUT("/", middlewares.AuthMiddleware(db), photoController.Update)
		photoRoute.DELETE("/", middlewares.AuthMiddleware(db), photoController.Delete)
	}

	return route
}
