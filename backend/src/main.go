package main

import (
	"log"
	"mazmorras_y_losetas/src/models"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Leer el Data Source Name (DSN) de las variables de entorno
	dsn := os.Getenv("DATABASE_URL")

	// Conectar a la base de datos usando GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al abrir la base de datos: %v", err)
	}

	// Migrar el esquema de la base de datos
	err = db.AutoMigrate(&models.User{}, &models.Boardgame{})
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

	// Registrar las rutas de Boardgame
	models.RegisterBoardgameRoutes(router, db)

	// Registrar las rutas de usuario
	models.RegisterUserRoutes(router, db)

	// Iniciar el servidor
	router.Run(":8080")
}
