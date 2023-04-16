package service

import "github.com/begenov/TaskFlow/cmd/entity"

type VedioService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VedioService {
	return &videoService{}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}
func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
