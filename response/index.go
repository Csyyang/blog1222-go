package response

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func BadRes(c *gin.Context, data interface{}) {
	fmt.Println(data)
	c.JSON(200, map[string]interface{}{
		"code":    "01",
		"message": data,
	})
}

func SuccessRes(c *gin.Context, data interface{}) {
	c.JSON(200, map[string]interface{}{
		"code":    "00",
		"message": data,
	})
}

func SuccessData(c *gin.Context, data interface{}) {
	c.JSON(200, map[string]interface{}{
		"code": "00",
		"data": data,
	})
}
