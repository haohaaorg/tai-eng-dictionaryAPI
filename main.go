// endpoint
// 		GET /api/v1/&from={from}&to={to}

// sample
// 		GET /api/v1/from/:text
//		GET /api/v1/eng/test
//		GET /api/v1/shn/လွင်ႈၸႅတ်ႈ

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/noernova/tai-eng-dictionaryAPI/controllers"
	"github.com/noernova/tai-eng-dictionaryAPI/models"
)

port := os.Getenv("PORT")

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", controllers.Index)

	router.GET("/api/v1/api_key=NO8p3FC4qMrTzx1RUjRXNXWrqlLa8DkDjmRgt7s9rDE=/eng/:text", controllers.Eng_to_Shn)
	router.GET("/api/v1/api_key=NO8p3FC4qMrTzx1RUjRXNXWrqlLa8DkDjmRgt7s9rDE=/shn/:text", controllers.Shn_to_Eng)

	return router
}

func main() {
	router := setupRouter()

	models.ConnectDatabase()
	defer models.CloseDatabase()

  if port == "" {
    port = "8080"
  }

	router.Run(port)
}
