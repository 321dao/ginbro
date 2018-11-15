package handlers

import (
	"github.com/dejavuzhou/ginbo/boilerplate/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("user", userAll)
	groupApi.GET("user/:id", userOne)
	groupApi.POST("user", userCreate)
	groupApi.PATCH("user", userUpdate)
	groupApi.DELETE("user/:id", userDelete)
}

func userAll(c *gin.Context) {
	mdl := models.User{}
	query := &models.PaginationQuery{}
	err := c.ShouldBindQuery(query)
	if handleError(c, err) {
		return
	}
	list, total, err := mdl.All(query)
	if handleError(c, err) {
		return
	}
	jsonPagination(c, list, total, query)
}
func userOne(c *gin.Context) {
	var mdl models.User
	id, err := parseParamID(c)
	if handleError(c, err) {
		return
	}
	mdl.Id = id
	data, err := mdl.One()
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}
func userCreate(c *gin.Context) {
	var mdl models.User
	err := c.ShouldBind(&mdl)
	if handleError(c, err) {
		return
	}

	err = mdl.Create()
	if handleError(c, err) {
		return
	}
	jsonData(c, mdl)
}

func userUpdate(c *gin.Context) {
	var mdl models.User
	err := c.ShouldBind(&mdl)
	if handleError(c, err) {
		return
	}
	err = mdl.Update()
	if handleError(c, err) {
		return
	}
	jsonSuccess(c)
}

func userDelete(c *gin.Context) {
	var mdl models.User
	id, err := parseParamID(c)
	if handleError(c, err) {
		return
	}
	mdl.Id = id
	err = mdl.Delete()
	if handleError(c, err) {
		return
	}
	jsonSuccess(c)
}
