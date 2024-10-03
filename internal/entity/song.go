package entity

import (
	"context"
	"effectiveMobile/internal/repo"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5"
)

type Song struct {
	SongId      uint   `json:"song_id"`
	SongName    string `json:"song_name"`
	SongText    string `json:"song_text"`
	ReleaseDate string `json:"release_date"`
	SongLink    string `json:"song_link"`
	SongAuthor  string `json:"song_author"`
}

func (s *Song) GetAllSongs() (*[]Song, error) {
	conn, err := repo.ConnectToDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT song_id, song_name, song_text, release_date, song_link, song_author FROM songs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []Song
	for rows.Next() {
		var song Song
		if err := rows.Scan(&song.SongId, &song.SongName, &song.SongText, &song.ReleaseDate, &song.SongLink, &song.SongAuthor); err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &songs, nil
}

func (s *Song) GetSongs(limit int) (*[]Song, error) {
	conn, err := repo.ConnectToDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT song_id, song_name, song_text, release_date, song_link, song_author FROM songs LIMIT $1", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []Song
	for rows.Next() {
		var song Song
		if err := rows.Scan(&song.SongId, &song.SongName, &song.SongText, &song.ReleaseDate, &song.SongLink, &song.SongAuthor); err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &songs, nil
}

func (s *Song) GetSongWithVyse(songId, verse int) ([]string, error) {

	conn, err := repo.ConnectToDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	query := "SELECT song_text FROM songs WHERE song_id = $1"
	var lyrics string
	err = conn.QueryRow(context.Background(), query, songId).Scan(&lyrics)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("Song not found")
		} else {
			return nil, errors.New("Error fetching verse")
		}
	}

	verses := strings.Split(lyrics, "|")

	if verse > len(verses) {
		verse = len(verses)
	}

	return verses[:verse], nil
}
