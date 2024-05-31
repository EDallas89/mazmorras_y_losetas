package main

import (
	"database/sql"
	"log"
	"net/http"
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

	router := gin.Default()
	router.GET("/db", func(ctx *gin.Context) {
		rows, err := db.Query("SELECT id, title, short_desc, release_year FROM boardgames")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var boardgames []map[string]interface{}
		for rows.Next() {
			var id int
			var title, short_desc string
			var release_year int
			if err := rows.Scan(&id, &title, &short_desc, &release_year); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			boardgames = append(boardgames, gin.H{
				"id":           id,
				"title":        title,
				"short_desc":   short_desc,
				"release_year": release_year,
			})
		}

		if err := rows.Err(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, boardgames)
	})

	schema, err := os.ReadFile("db/schema.sql")
	if err != nil {
		log.Fatalf("Error al leer el archivo schema.sql: %v", err)
	}
	_, err = db.Exec(string(schema))
	if err != nil {
		log.Fatalf("Error al ejecutar el script SQL: %v", err)
	}

	router.Run(":8080")
}
