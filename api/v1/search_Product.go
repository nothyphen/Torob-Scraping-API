package v1

import (
	"net/http"
	torobproduct "webscrap/pkg/torob_product"
	"webscrap/serilizers"

	"github.com/gin-gonic/gin"
)

type SearchApi interface {
	Search(ctx *gin.Context)
}

func Search(ctx *gin.Context) {
	var searchRequest serilizers.SearchReques
	err := ctx.ShouldBind(&searchRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : "bad request",
		})
	}

	result, err :=torobproduct.Torob(searchRequest.Product)
	if err != nil{
		ctx.JSON(http.StatusNoContent, gin.H{
			"error" : "your product not found",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"result" : result,
		})
	}
}