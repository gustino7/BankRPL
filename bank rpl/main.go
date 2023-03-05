package main

import (
	"bankrpl/entity"
	"bankrpl/handler"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// entity for database

// response in api

// setting database
func setUpDataBase() *gorm.DB {
	dsn := "host=localhost user=postgres password=123456 dbname=pertemuan3 port=5432 TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if err := db.AutoMigrate(
		entity.Nasabah{},
		entity.Rekening{},
	); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	dbSQL.Close()
}

// middleware
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// NASABAH HANDLER #############################

// REKENING HANDLER #####################################

func main() {
	database := setUpDataBase()
	defer CloseDatabaseConnection(database)

	server := gin.Default()
	server.Use(
		Middleware(),
	)

	nasabahHandler := handler.NasabahHandler{DB: database}
	rekeningHandler := handler.RekeningHandler{DB: database}

	//server request
	server.POST("/register/nasabah", nasabahHandler.HandlerRegisterNasabah)
	server.DELETE("/delete/nasabah", nasabahHandler.HandlerDeleteNasabah)
	server.PUT("/update/nasabah", nasabahHandler.HandlerUpdateNasabah)
	server.GET("/show/rekening/:id", nasabahHandler.HandlerGetNasabah)
	server.POST("/register/rekening", rekeningHandler.HandlerRegisterRekening)
	server.DELETE("/delete/rekening", rekeningHandler.HandlerDeleteRekening)

	server.Run(":8080")
}
