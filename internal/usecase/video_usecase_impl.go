package usecase

import (
	"tiktok-player-backend/internal/domain"
	"tiktok-player-backend/internal/repository"
)

type videoUsecaseImpl struct {
	videoRepo repository.VideoRepository
}

func NewVideoUsecase(repo repository.VideoRepository) VideoUsecase {
	return &videoUsecaseImpl{videoRepo: repo}
}

func (u *videoUsecaseImpl) SearchVideos(keyword string) ([]domain.Video, error) {
	return u.videoRepo.FetchVideos(keyword)
}
