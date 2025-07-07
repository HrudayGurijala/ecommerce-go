package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/HrudayGurijala/ecommerce-go/database"
	"github.com/HrudayGurijala/ecommerce-go/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection = database.UserData(database.Client, "Users")
var ProductCollection *mongo.Collection = database.ProductData(database.Client, "Products")
var Validate = validator.New()

// hashpassword
func HashPassword(password string) string {

}

// verifypassword
func VerifyPassword(userPassword string, givenPassword string) (bool, string) {

}

// signup
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := Validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			return
		}
		
	}
}

// login
func Login() gin.HandlerFunc {

}

// productviewer admin
func ProductViewerAdmin() gin.HandlerFunc {

}

// search product
func Search() gin.HandlerFunc {

}

// search product by params
func SearchProductByQuery() gin.HandlerFunc {

}
