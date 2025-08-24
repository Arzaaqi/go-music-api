package controllers

import (
	"go-music-api/models"
	"go-music-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MusicController struct {
	service services.MusicService
}

func NewMusicController(service services.MusicService) *MusicController {
	return &MusicController{service}
}

// GetAll godoc
// @Summary Get all musics
// @Description Ambil semua data musik
// @Tags musics
// @Produce json
// @Success 200 {array} models.Music
// @Router /musics [get]
func (ctl *MusicController) GetMusics(c *gin.Context) {
	musics, err := ctl.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, musics)
}

// GetMusicById godoc
// @Summary Get music by id
// @Description Ambil data musik berdasarkan ID
// @Tags musics
// @Produce json
// @Param id path int true "Music ID"
// @Success 200 {object} models.Music
// @Failure 404 {string} string "Not Found"
// @Router /musics/{id} [get]
func (ctl *MusicController) GetMusicById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	music, err := ctl.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "music not found"})
		return
	}
	c.JSON(http.StatusOK, music)
}

// Create godoc
// @Summary Tambah musik baru
// @Description Tambahkan musik ke dalam koleksi
// @Tags musics
// @Accept json
// @Produce json
// @Param music body models.Music true "Music Data"
// @Success 201 {object} models.Music
// @Router /musics [post]
func (ctl *MusicController) CreateMusic(c *gin.Context) {
	var newMusic models.Music
	if err := c.ShouldBindJSON(&newMusic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := ctl.service.Add(newMusic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

// Update godoc
// @Summary Update musik
// @Description Edit data musik berdasarkan ID
// @Tags musics
// @Accept json
// @Produce json
// @Param id path int true "Music ID"
// @Param music body models.Music true "Music Data"
// @Success 200 {object} models.Music
// @Failure 404 {string} string "Not Found"
// @Router /musics/{id} [put]
func (ctl *MusicController) UpdateMusic(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var newMusic models.Music
	if err := c.ShouldBindJSON(&newMusic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := ctl.service.Edit(uint(id), newMusic)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "music not found"})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// Delete godoc
// @Summary Hapus musik
// @Description Hapus data musik berdasarkan ID
// @Tags musics
// @Param id path int true "Music ID"
// @Success 200 {string} string "Deleted successfully"
// @Failure 404 {string} string "Not Found"
// @Router /musics/{id} [delete]
func (ctl *MusicController) DeleteMusic(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctl.service.Remove(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "music not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "music deleted"})
}
