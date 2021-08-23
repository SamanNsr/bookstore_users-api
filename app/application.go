package app

import (
	"github.com/SamanNsr/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	
	logger.Info("about to start the applications...")
	router.Run(":8080")
}
