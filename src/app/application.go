package app

import (
	"github.com/DeKal/bookstore_oath-api/src/http"
	"github.com/DeKal/bookstore_oath-api/src/repository/db"
	"github.com/DeKal/bookstore_oath-api/src/repository/rest"
	svcaccesstoken "github.com/DeKal/bookstore_oath-api/src/service/access_token"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication start the whole application
func StartApplication() {
	dbRepository := db.NewDBRepository()
	restRepository := rest.NewRepository()
	accessTokenService := svcaccesstoken.NewService(dbRepository, restRepository)
	accesstokenHandler := http.NewHandler(accessTokenService)

	mappingUrls(accesstokenHandler)

	router.Run(":9002")
}

func mappingUrls(handler http.AccessTokenHandler) {
	router.GET("/oauth/access_token/:access_token_id", handler.GetByID)
	router.POST("/oauth/access_token", handler.Create)
}
