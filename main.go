// endpoint
// 		GET /api/v1/&from={from}&to={to}

// sample
// 		GET /api/v1/from/:text
//		GET /api/v1/eng/test
//		GET /api/v1/shn/လွင်ႈၸႅတ်ႈ

package main

import (
	"os"

	"github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
	"github.com/noernova/tai-eng-dictionaryAPI/controllers"
	"github.com/noernova/tai-eng-dictionaryAPI/models"
)

func setupRouter() *gin.Engine {
	ginMode := os.Getenv("GIN_MODE")

	if ginMode == "" { 
		ginMode = "debug"
	}

	gin.SetMode(ginMode)

	router := gin.Default()

  router.SetTrustedProxies(nil)

  router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  false,
		AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge:           86400,
	}))

	router.GET("/", controllers.Index)

	router.GET("/api/v1/api_key=NO8p3FC4qMrTzx1RUjRXNXWrqlLa8DkDjmRgt7s9rDE=/eng/:text", controllers.Eng_to_Shn)
	router.GET("/api/v1/api_key=NO8p3FC4qMrTzx1RUjRXNXWrqlLa8DkDjmRgt7s9rDE=/shn/:text", controllers.Shn_to_Eng)

	return router
}

func main() {
	router := setupRouter()

	models.ConnectDatabase()
	defer models.CloseDatabase()

  port := os.Getenv("PORT")

  if port == "" {
    port = "8080"
  }

	router.Run(":"+port)
}
