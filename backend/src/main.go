package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error al abrir la base de datos: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	schema, err := os.ReadFile("db/schema.sql")
	if err != nil {
		log.Fatalf("Error al leer el archivo schema.sql: %v", err)
	}
	_, err = db.Exec(string(schema))
	if err != nil {
		log.Fatalf("Error al ejecutar el script SQL: %v", err)
	}

	r.Run(":8080")
}
