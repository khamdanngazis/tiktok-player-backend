package handler

import (
	"net/http"
	"tiktok-player-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type VideoHandler struct {
	videoUsecase usecase.VideoUsecase
}

func NewVideoHandler(router *gin.Engine, videoUsecase usecase.VideoUsecase) {
	handler := &VideoHandler{videoUsecase: videoUsecase}
	router.GET("/search", handler.SearchVideos)
}

func (h *VideoHandler) SearchVideos(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	keyword := c.Query("keyword")
	if keyword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Keyword is required"})
		return
	}

	videos, err := h.videoUsecase.SearchVideos(keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch videos"})
		return
	}

	c.JSON(http.StatusOK, videos)
}
