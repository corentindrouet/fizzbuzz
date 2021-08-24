package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "strconv"
  "fmt"
)


func main() {

  router := gin.Default()

  router.GET("replace", replace)
  router.Run("localhost:8080")
}


func replace(c *gin.Context) {

  int1, _ := strconv.Atoi(c.DefaultQuery("int1", "3"))
  int2, _ := strconv.Atoi(c.DefaultQuery("int2", "5"))
  limit, _ := strconv.Atoi(c.DefaultQuery("limit", "100"))
  str1 := c.DefaultQuery("str1", "fizz")
  str2 := c.DefaultQuery("str2", "buzz")
  fmt.Println(c.Request.URL.Query())

  rep := make([]string, 0)
  var str string

  for i := 1; i <= limit; i++ {
    str = ""
    if i%int1 == 0 {
      str += str1
    }
    if i%int2 == 0 {
      str += str2
    }
    if str == "" {
      str = strconv.Itoa(i)
    }
    rep = append(rep, str)
  }
  c.JSON(http.StatusOK, rep)
}
