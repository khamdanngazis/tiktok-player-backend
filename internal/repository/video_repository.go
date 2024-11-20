package repository

import "tiktok-player-backend/internal/domain"

type VideoRepository interface {
	FetchVideos(keyword string) ([]domain.Video, error)
}
