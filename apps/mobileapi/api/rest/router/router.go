package router

import (
	"context"

	"pillsreminder/apps/mobileapi/api/rest/controllers"
	"pillsreminder/apps/mobileapi/api/rest/middlewares"
	"pillsreminder/apps/mobileapi/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup(ctx context.Context, config config.AppConfig, controllerFactory controllers.ControllerFactory) *gin.Engine {
	if config.Server.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()

	router.Use(gin.Recovery())
	router.NoRoute(middlewares.NoRouteHandler())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	NoAuthRoutes(router, controllerFactory)
	AuthRoutes(router, config.Server.Secure.JwtSecret, controllerFactory)

	return router
}

func NoAuthRoutes(router *gin.Engine, controllerFactory controllers.ControllerFactory) {
	authController := controllerFactory.GetAuthController()

	router.POST("/signup", authController.SignUp)
	router.POST("/login", authController.Login)
}

func AuthRoutes(router *gin.Engine, jwtSecret string, controllerFactory controllers.ControllerFactory) {
	pillsController := controllerFactory.GetPillsController()

	authorized := router.Group("/")
	authorized.Use(middlewares.AuthBase(jwtSecret))

	authorized.GET("/pills", pillsController.GetPills)
	authorized.GET("/pills/:pillID", pillsController.GetPillByID)

	authorized.GET("/plans", pillsController.GetPlans)
	authorized.GET("/plans/:planID", pillsController.GetPlanByID)

	authorized.POST("/plans", pillsController.AddPlan)
	authorized.DELETE("/plans/:planID", pillsController.AddPlan)
}
