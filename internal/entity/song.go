package entity

import (
	"context"
	"effectiveMobile/internal/repo"
	"errors"
	"fmt"
	"log/slog"
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

func (s *Song) DeleteSong(id int) error {
	conn, err := repo.ConnectToDB()
	if err != nil {
		slog.Warn("Can't connect to DB")
		return err
	}
	defer conn.Close(context.Background())

	query := "DELETE FROM songs WHERE song_id = $1"
	_, err = conn.Exec(context.Background(), query, id)
	if err != nil {
		slog.Info(err.Error())
		return err
	}

	slog.Info("SONG Succefully deleted")
	return nil
}

func (s *Song) UpdateSong(id int, newSong Song) error {
	query := "UPDATE songs SET"
	updates := []string{}
	args := []interface{}{}
	argIndex := 1

	if newSong.SongName != "" {
		updates = append(updates, fmt.Sprintf(" song_name = $%d", argIndex))
		args = append(args, newSong.SongName)
		argIndex++
	}
	if newSong.SongText != "" {
		updates = append(updates, fmt.Sprintf(" song_text = $%d", argIndex))
		args = append(args, newSong.SongText)
		argIndex++
	}
	if newSong.ReleaseDate != "" {
		updates = append(updates, fmt.Sprintf(" release_date = $%d", argIndex))
		args = append(args, newSong.ReleaseDate)
		argIndex++
	}
	if newSong.SongLink != "" {
		updates = append(updates, fmt.Sprintf(" song_link = $%d", argIndex))
		args = append(args, newSong.SongLink)
		argIndex++
	}
	if newSong.SongAuthor != "" {
		updates = append(updates, fmt.Sprintf(" song_author = $%d", argIndex))
		args = append(args, newSong.SongAuthor)
		argIndex++
	}

	if len(updates) == 0 {
		return fmt.Errorf("no fields to update")
	}

	query += fmt.Sprintf("%s WHERE song_id = $%d", joinUpdates(updates), argIndex)
	args = append(args, id)

	conn, err := repo.ConnectToDB()
	if err != nil {
		slog.Warn("Can't connect to DB: %v", err)
		return err
	}

	// debug
	fmt.Println("Executing query:", query, "with args:", args)

	_, err = conn.Exec(context.Background(), query, args...)
	if err != nil {
		slog.Warn("Can't update dataaaa: %v", err)
		return err
	}
	return nil
}

func joinUpdates(updates []string) string {
	return " " + join(updates, ", ")
}

func join(a []string, sep string) string {
	if len(a) == 0 {
		return ""
	}
	result := a[0]
	for _, s := range a[1:] {
		result += sep + s
	}
	return result
}

func (s *Song) NewSong() (*int, error) {
	if s.SongName == "" {
		return nil, fmt.Errorf("song_name must not be empty")
	}
	if s.SongText == "" {
		return nil, fmt.Errorf("song_text must not be empty")
	}
	if s.ReleaseDate == "" {
		return nil, fmt.Errorf("release_date must not be empty")
	}
	if s.SongLink == "" {
		return nil, fmt.Errorf("song_link must not be empty")
	}
	if s.SongAuthor == "" {
		return nil, fmt.Errorf("song_author must not be empty")
	}

	query := "INSERT INTO songs (song_name, song_text, release_date, song_link, song_author) VALUES ($1, $2, $3, $4, $5)"

	conn, err := repo.ConnectToDB()
	if err != nil {
		slog.Warn("Can't connect to DB: %v", err)
		return nil, err
	}

	_, err = conn.Exec(context.Background(), query, s.SongName, s.SongText, s.ReleaseDate, s.SongLink, s.SongAuthor)
	if err != nil {
		slog.Warn("Can't insert new song: %v", err)
		return nil, err
	}

	var songId int
	query1 := "SELECT song_id FROM songs WHERE song_name = $1 "
	err = conn.QueryRow(context.Background(), query1, s.SongName).Scan(&songId)
	if err != nil {
		slog.Warn("Can't get ID of new song: %v", err)
		return nil, err
	}

	return &songId, nil
}
