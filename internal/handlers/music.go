package handlers

import (
	"effectiveMobile/internal/entity"
	"fmt"
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
	pageParam := ctx.Param("song_id")
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

func DeleteSong(ctx *gin.Context) {
	songID := ctx.Param("song_id")

	id, err := strconv.Atoi(songID)
	if err != nil || id < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number of song id"})
		return
	}

	var song entity.Song
	err = song.DeleteSong(id)
	if err != nil {
		ctx.JSON(404, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Response": "Song deleted"})
}

func UpdateSong(ctx *gin.Context) {
	var newSong entity.Song
	id := ctx.Param("song_id")

	// Преобразуем строку id в uint
	var songId int
	if _, err := fmt.Sscanf(id, "%d", &songId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid song id"})
		return
	}

	if err := ctx.ShouldBindJSON(&newSong); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if err := newSong.UpdateSong(songId, newSong); err != nil {
		ctx.JSON(404, "Can't update")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "song updated successfully"})
}

func NewSong(ctx *gin.Context) {
	var song entity.Song
	err := ctx.ShouldBindJSON(&song)
	if err != nil {
		ctx.JSON(404, err.Error())
		return
	}

	songId, err := song.NewSong()
	if err != nil {
		ctx.JSON(404, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Response": "Song created.", "Song ID": *songId})
}
