package app

import (
	"github.com/SamanNsr/bookstore_users-api/controllers/ping"
	"github.com/SamanNsr/bookstore_users-api/controllers/user"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users", user.Search)
	router.GET("/users/:user_id", user.GetUser)
	router.POST("/users", user.CreateUser)
	router.PATCH("/users/:user_id", user.UpdateUser)
	router.DELETE("/users/:user_id", user.DeleteUser)
}
