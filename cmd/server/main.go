package main

import (
	"tiktok-player-backend/internal/infrastructure"
	"tiktok-player-backend/internal/interface/handler"
	"tiktok-player-backend/internal/middleware"
	"tiktok-player-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Gin router
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	// Setup dependencies
	videoRepo := infrastructure.NewTikTokScraper()
	videoUsecase := usecase.NewVideoUsecase(videoRepo)
	handler.NewVideoHandler(router, videoUsecase)

	// Start server
	router.Run(":8080")
}
