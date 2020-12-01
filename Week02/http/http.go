package http

import (
	"github.com/gin-gonic/gin"
	"handleError/internal/service"
	"log"
	"net/http"
)


type userParameter struct {
	UserName string `json:"user_name"`
}

func UserInfo(c *gin.Context)  {
	parameter := new(userParameter)
	err := c.ShouldBindJSON(&parameter)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "error",
			"detail": err.Error(),
		})
		return
	}
	info, err := service.GetUserInfo(parameter.UserName)
	if err != nil {
		log.Printf("get_user_info: %+v", err)
		c.JSON(200, gin.H{
			"reason": err.Error(),
			"status": "fail",
			"userInfo": "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "success",
		"userInfo": info,
		"reason": "",
	})
}
