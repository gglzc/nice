package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/dgrijalva/jwt-go"
  routers "/Users/laishengqi/code/nice/routers"
)


func main() {
  router := gin.Default()
  router.Use(gin.Logger())
  router.LoadHTMLGlob("/template/login")
  router.GET("/item", func(c *gin.Context) {
    c.JSON(http.StatusOK,gin.H{
      "success":"Accept the data",
    })
  })
  router.POST("/account/id" ,)  //查詢會員資料
  router.GET("/item")
  router.POST("/item")
  routers.AuthRouters(router)
  
  router.Run(":8000") 
}