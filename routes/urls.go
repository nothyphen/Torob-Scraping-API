package routes

import (
	v1 "webscrap/api/v1"

	"github.com/gin-gonic/gin"
)

func Urls() *gin.Engine {
	r := gin.Default()

	apiv1 := r.Group("api/v1")
	{
		product := apiv1.Group("product")
		{
			product.POST("/search", v1.Search)
		}
	}

	return r
}