package infrastructure

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"
	"tiktok-player-backend/internal/domain"
	"tiktok-player-backend/internal/repository"
	"time"

	"github.com/chromedp/chromedp"
)

type TikTokScraper struct{}

func NewTikTokScraper() repository.VideoRepository {
	return &TikTokScraper{}
}

func (s *TikTokScraper) FetchVideos(keyword string) ([]domain.Video, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var videoLinks []string
	url := fmt.Sprintf("https://www.tiktok.com/search?q=%s", keyword)

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(5*time.Second), // Wait for content to load
		chromedp.Evaluate(`Array.from(document.querySelectorAll('a')).map(a => a.href)`, &videoLinks),
	)

	if err != nil {
		log.Fatalf("Failed to scrape: %v", err)
	}

	var videos []domain.Video
	for _, link := range videoLinks {
		tikTokVideoPattern := `^https://www\.tiktok\.com/@[\w.-]+/video/\d+$`
		videoRegex := regexp.MustCompile(tikTokVideoPattern)
		if videoRegex.MatchString(link) {
			stringSlice := strings.Split(link, "/")
			video := domain.Video{
				ID:   stringSlice[5],
				User: stringSlice[3],
				URL:  link,
			}
			// Append the video to the slice
			videos = append(videos, video)

		}
	}

	return videos, nil
}
