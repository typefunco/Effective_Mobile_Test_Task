package handlers

import (
	"effectiveMobile/internal/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
	ctx.JSON(200, "It's working")
	return
}

func GetAllSongs(ctx *gin.Context) {
	var song entity.Song

	songs, err := song.GetAllSongs()
	if err != nil {
		ctx.JSON(500, "Can't get songs")
		return
	}
	ctx.JSON(200, songs)
}

func GetSongs(ctx *gin.Context) {
	pageParam := ctx.Param("number")
	limit, err := strconv.Atoi(pageParam)
	if err != nil || limit < 1 {
		ctx.JSON(400, gin.H{"error": "Invalid page number"})
		return
	}

	var song entity.Song
	songs, err := song.GetSongs(limit)
	if err != nil {
		ctx.JSON(500, "Can't get songs")
		return
	}

	ctx.JSON(200, songs)
}

func GetSong(ctx *gin.Context) {
	songID := ctx.Param("song_id")
	verseCount := ctx.Param("verse")

	id, err := strconv.Atoi(songID)
	if err != nil || id < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number of verses."})
		return
	}

	count, err := strconv.Atoi(verseCount)
	if err != nil || count < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number of verses."})
		return
	}

	var song entity.Song
	text, err := song.GetSongWithVyse(id, count)
	if err != nil {
		ctx.JSON(404, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"verses": text})
}
