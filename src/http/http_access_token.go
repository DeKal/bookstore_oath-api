package http

import (
	"net/http"

	accesstoken "github.com/DeKal/bookstore_oath-api/src/domain/access_token"
	svcaccesstoken "github.com/DeKal/bookstore_oath-api/src/service/access_token"
	"github.com/DeKal/bookstore_utils-go/errors"
	"github.com/gin-gonic/gin"
)

// AccessTokenHandler handle all call to Access token
type AccessTokenHandler interface {
	GetByID(c *gin.Context)
	Create(c *gin.Context)
}
type accessTokenHandler struct {
	service svcaccesstoken.Service
}

// NewHandler return new handler for access token
func NewHandler(service svcaccesstoken.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetByID(c *gin.Context) {
	accessTokenID := c.Param("access_token_id")
	accessToken, err := h.service.GetByID(accessTokenID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	request := &accesstoken.Request{}
	if err := c.ShouldBindJSON(request); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	accessToken, err := h.service.Create(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}
