package repositories

import (
	"go-music-api/config"
	"go-music-api/models"
)

type MusicRepository interface {
	GetAllMusics() ([]models.Music, error)
	GetMusicById(id uint) (models.Music, error)
	CreateMusic(music models.Music) (models.Music, error)
	UpdateMusic(id uint, music models.Music) (models.Music, error)
	DeleteMusic(id uint) bool
}

type musicRepository struct{}

func NewMusicRepository() MusicRepository {
	return &musicRepository{}
}

func (m *musicRepository) GetAllMusics() ([]models.Music, error) {
	var musics []models.Music
	result := config.DB.Find(&musics)
	return musics, result.Error
}

func (m *musicRepository) GetMusicById(id uint) (models.Music, error) {
	var music models.Music
	result := config.DB.First(&music, id)
	return music, result.Error
}

func (m *musicRepository) CreateMusic(music models.Music) (models.Music, error) {
	result := config.DB.Create(&music)
	return music, result.Error
}

func (m *musicRepository) UpdateMusic(id uint, newMusic models.Music) (models.Music, error) {
	var music models.Music
	if err := config.DB.First(&music, id).Error; err != nil {
		return music, err
	}

	// update semua field sesuai model barumu
	if err := config.DB.Model(&music).Updates(newMusic).Error; err != nil {
		return music, err
	}

	err := config.DB.Save(&music).Error
	return music, err
}

func (m *musicRepository) DeleteMusic(id uint) bool {
	var music models.Music
	if err := config.DB.Delete(&music, id).Error; err != nil {
		return false
	}
	return true
}
