package routes

import (
	"github.com/gin-gonic/gin"
	"go-music-api/controllers"
)

func SetupRoutes(musicController *controllers.MusicController) *gin.Engine {
	r := gin.Default()

	r.GET("/musics", musicController.GetMusics)
	r.GET("/musics/:id", musicController.GetMusicById)
	r.POST("/musics", musicController.CreateMusic)
	r.PUT("/musics/:id", musicController.UpdateMusic)
	r.DELETE("/musics/:id", musicController.DeleteMusic)
	return r
}
