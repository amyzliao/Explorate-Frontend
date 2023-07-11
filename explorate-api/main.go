package main

import (
	"context"
	"explorate/controllers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func main() {
	dsn := "postgresql://postgres:MIEBa_yCmoJCc-phZQPJXA@explorate-11582.7tt.cockroachlabs.cloud:26257/explorate-ngo-opps?sslmode=verify-full"

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil || conn == nil {
		log.Fatal("database issue")
	}

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE, PATCH")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	router.GET("/ngo_opps", controllers.GetAll(conn, "ngo_opps"))
	router.GET("/ngo_opps/:id", controllers.Find(conn, "ngo_opps"))
	router.POST("/ngo_opps", controllers.Insert(conn, "ngo_opps"))
	router.DELETE("/ngo_opps/:id", controllers.Delete(conn, "ngo_opps"))
	router.Run(":3000")
}
