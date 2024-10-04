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

// @Summary Get all songs
// @Description Retrieves all songs from the database
// @Tags songs
// @Produce json
// @Security Bearer
// @Success 200 {array} entity.Song
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string
// @Router /music/songs [get]
func GetAllSongs(ctx *gin.Context) {
	var song entity.Song

	songs, err := song.GetAllSongs()
	if err != nil {
		ctx.JSON(500, "Can't get songs")
		return
	}
	ctx.JSON(200, songs)
}

// @Summary Get songs with pagination
// @Description Retrieves a specified number of songs
// @Tags songs
// @Produce json
// @Param song_id path int true "Number of songs to retrieve"
// @Success 200 {array} entity.Song
// @Failure 400 {object} string "Invalid page number"
// @Failure 500 {object} string "Can't get songs"
// @Router /music/songs/{song_id} [get]
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

// @Summary Get specific song verses
// @Description Retrieves a specific number of verses from a song
// @Tags songs
// @Produce json
// @Param song_id path int true "Song ID"
// @Param verse path int true "Number of verses to retrieve"
// @Success 200 {object} map[string][]string
// @Failure 400 {object} string "Invalid input"
// @Failure 404 {object} string
// @Router /music/song/{song_id}/{verse} [get]
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
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"verses": text})
}

// @Summary Delete a song
// @Description Deletes a song from the database
// @Tags songs
// @Produce json
// @Security Bearer
// @Param song_id path int true "Song ID to delete"
// @Success 200 {object} map[string]string
// @Failure 400 {object} string "Invalid song id"
// @Failure 401 {object} string "Unauthorized"
// @Failure 403 {object} string "Forbidden"
// @Failure 404 {object} string
// @Router /music/song/{song_id} [delete]
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

// @Summary Update a song
// @Description Updates an existing song in the database
// @Tags songs
// @Accept json
// @Produce json
// @Security Bearer
// @Param song_id path int true "Song ID to update"
// @Param song body entity.Song true "Updated song information"
// @Success 200 {object} map[string]string
// @Failure 400 {object} string "Invalid input"
// @Failure 401 {object} string "Unauthorized"
// @Failure 403 {object} string "Forbidden"
// @Failure 404 {object} string "Can't update"
// @Router /music/songs/{song_id} [patch]
func UpdateSong(ctx *gin.Context) {
	var newSong entity.Song
	id := ctx.Param("song_id")

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

// @Summary Create a new song
// @Description Adds a new song to the database
// @Tags songs
// @Accept json
// @Produce json
// @Security Bearer
// @Param song body entity.Song true "New song information"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} string
// @Failure 401 {object} string "Unauthorized"
// @Failure 403 {object} string "Forbidden"
// @Router /music/song/new [post]
func NewSong(ctx *gin.Context) {
	var song entity.Song
	err := ctx.ShouldBindJSON(&song)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	songId, err := song.NewSong()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Response": "Song created.", "Song ID": *songId})
}
