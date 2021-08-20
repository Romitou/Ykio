package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {
	db := ConnectDatabase(os.Getenv("DB_DSN"))
	token := os.Getenv("SEND_TOKEN")
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to Ykio! :)",
		})
	})
	r.POST("/send", func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(400, gin.H{
				"error": "No token is present in the header of your request.",
			})
			return
		}
		bearerToken := strings.TrimPrefix(authHeader, "Bearer ")
		if token != bearerToken {
			ctx.JSON(401, gin.H{
				"error": "The provided token is invalid.",
			})
			return
		}
		file, err := ctx.FormFile("content")
		if err != nil {
			log.Println(err)
			ctx.JSON(500, gin.H{
				"error": "An internal error has occurred. Try again later!",
			})
			return
		}

		image := Image{
			Name:  file.Filename,
			Views: 0,
		}

		err = ctx.SaveUploadedFile(file, "./images/"+image.Name)
		if err != nil {
			log.Println(err)
			ctx.JSON(500, gin.H{
				"error": "An internal error has occurred. Try again later!",
			})
			return
		}

		result := db.Create(&image)
		if result.Error != nil {
			log.Println(err)
			ctx.JSON(500, gin.H{
				"error": "An internal error has occurred. Try again later!",
			})
			return
		}

		ctx.JSON(200, gin.H{
			"message":  "The file has been saved.",
			"fileName": image.Name,
		})

	})
	r.GET("/:name", func(ctx *gin.Context) {
		start := time.Now()
		name := ctx.Param("name")
		var image Image
		result := db.Where(&Image{
			Name: name,
		}).First(&image)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				ctx.JSON(404, gin.H{
					"error": "This image does not exist or no longer exists.",
				})
				return
			}
			log.Println(result.Error)
			ctx.JSON(500, gin.H{
				"error": "An internal error has occurred. Try again later!",
			})
			return
		}
		data, err := ioutil.ReadFile("./images/" + image.Name)
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": "This image does not exist or no longer exists.",
			})
			db.Delete(&image)
			return
		}
		ctx.Header("X-Image-Id", strconv.Itoa(int(image.ID)))
		ctx.Header("X-Processing-Time", time.Since(start).String())
		ctx.Data(200, "image/"+filepath.Ext(image.Name)[1:], data)
		image.Views++
		db.Save(&image)
	})
	err := r.Run()
	if err != nil {
		log.Fatal("an error occurred while starting the application: ", err)
	}
}
