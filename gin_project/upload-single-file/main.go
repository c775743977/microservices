package main

import (
    "github.com/gin-gonic/gin"
	"net/http"
)

func main() {
    r := gin.Default()
    //限制上传最大尺寸
    r.MaxMultipartMemory = 8 << 20
	r.LoadHTMLFiles("./upload-single-file/index.html")
	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
    r.POST("/upload", func(c *gin.Context) {
        file, err := c.FormFile("file")
        if err != nil {
            c.String(500, "上传图片出错")
        }
        // c.JSON(200, gin.H{"message": file.Header.Context})
        c.SaveUploadedFile(file, "321.jpg")
        c.String(http.StatusOK, file.Filename)
    })
    r.Run()
}