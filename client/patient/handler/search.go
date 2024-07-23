package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"weikang/client/patient/form"
	"weikang/pkg"
)

func EsSearch(c *gin.Context) {
	var search form.EsSearch
	if err := c.ShouldBindJSON(&search); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	where, err := pkg.EsSearchWhere(search.Index, search.BeginTime, search.EndTime, search.Name, search.Status, search.Page, search.Size)
	if err != nil {
		c.JSON(http.StatusBadRequest, &gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	var res []map[string]interface{}
	json.Unmarshal(where, &res)
	c.JSON(http.StatusOK, &gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    res,
	})

}
