package uploads

import (
	"blog1222-go/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Uploadfile_image(c *gin.Context) {
	//获取表单数据 参数为name值
	f, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	//将文件保存至本项目根目录中
	c.SaveUploadedFile(f, "./images/"+f.Filename)
	//保存成功返回正确的Json数据
	c.JSON(http.StatusOK, gin.H{
		"code":    "00",
		"message": "OK",
		"url":     config.Configs.ServeConfig.Ip + "/images/" + f.Filename,
	})

}
