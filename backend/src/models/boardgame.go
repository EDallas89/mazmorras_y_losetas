package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func RegisterBoardgameRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/boardgames", func(ctx *gin.Context) {
		var boardgames []Boardgame
		if err := db.Find(&boardgames).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, boardgames)
	})
}
