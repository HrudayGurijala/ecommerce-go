package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/HrudayGurijala/ecommerce-go/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup",controllers.SignUp())
	incomingRoutes.POST("/users/login",controllers.Login())
	incomingRoutes.POST("/admin/addproduct",controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview",controllers.SearchProduct())
	incomingRoutes.GET("/users/search",controllers.SearchProductByQuery())

}