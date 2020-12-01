package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"handleError/http"
)

func main()  {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.POST("/get_user_info", http.UserInfo)
	ch := make(chan error)
	_ = router.Run(fmt.Sprintf("127.0.0.1:8909"))
	<- ch
}