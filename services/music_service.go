package services

import (
	"fmt"
	"go-music-api/models"
	"go-music-api/repositories"
)

type MusicService interface {
	GetAll() ([]models.Music, error)
	GetByID(id uint) (models.Music, error)
	Add(music models.Music) (models.Music, error)
	Edit(id uint, music models.Music) (models.Music, error)
	Remove(id uint) error
}

type musicService struct {
	repo repositories.MusicRepository
}

func NewMusicService(repo repositories.MusicRepository) MusicService {
	return &musicService{repo}
}

func (s *musicService) GetAll() ([]models.Music, error) {
	return s.repo.GetAllMusics()
}

func (s *musicService) GetByID(id uint) (models.Music, error) {
	return s.repo.GetMusicById(id)
}

func (s *musicService) Add(music models.Music) (models.Music, error) {
	return s.repo.CreateMusic(music)
}

func (s *musicService) Edit(id uint, music models.Music) (models.Music, error) {
	return s.repo.UpdateMusic(id, music)
}

func (s *musicService) Remove(id uint) error {
	// repo.DeleteMusic return bool → sebaiknya diubah jadi error
	success := s.repo.DeleteMusic(id)
	if !success {
		return fmt.Errorf("gagal hapus music dengan id %d", id)
	}
	return nil
}
