package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Boardgame struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	ShortDesc   string `json:"short_desc"`
	LongDesc    string `json:"long_desc"`
	ReleaseYear int    `json:"release_year"`
	MinPlayer   int    `json:"min_player"`
	MaxPlayer   int    `json:"max_player"`
	MinTime     int    `json:"min_time"`
	MaxTime     int    `json:"max_time"`
	Age         int    `json:"age"`
}

func main() {
	// Leer el Data Source Name (DSN) de las variables de entorno
	dsn := os.Getenv("DATABASE_URL")

	// Conectar a la base de datos usando GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al abrir la base de datos: %v", err)
	}

	// Migrar el esquema de la base de datos
	err = db.AutoMigrate(&Boardgame{})
	if err != nil {
		log.Fatalf("Error al migrar el esquema de la base de datos: %v", err)
	}

	// Leer y ejecutar el script SQL del archivo schema.sql
	schema, err := os.ReadFile("db/schema.sql")
	if err != nil {
		log.Fatalf("Error al leer el archivo schema.sql: %v", err)
	}
	db.Exec(string(schema))

	// Configurar el router de Gin
	router := gin.Default()
	router.GET("/db", func(ctx *gin.Context) {
		var boardgames []Boardgame
		if err := db.Find(&boardgames).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, boardgames)
	})

	// Iniciar el servidor
	router.Run(":8080")
}
