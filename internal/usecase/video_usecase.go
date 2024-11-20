package usecase

import "tiktok-player-backend/internal/domain"

type VideoUsecase interface {
	SearchVideos(keyword string) ([]domain.Video, error)
}
