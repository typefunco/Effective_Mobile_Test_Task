package repo

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	//"github.com/joho/godotenv"
)

// todo: Rewrite func for Docker container

// think about it
type DB struct {
	conn *pgx.Conn
}

func ConnectToDB() (*pgx.Conn, error) {
	var db DB
	// err := godotenv.Load()
	// if err != nil {
	// 	slog.Info("Error loading .env file")
	// 	return nil, err
	// }

	dbUrl := os.Getenv("DBURL")
	fmt.Println(dbUrl)
	conn, err := pgx.Connect(context.Background(), dbUrl)
	fmt.Println(err)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	err = conn.Ping(context.Background())
	fmt.Println(err)

	if err != nil {
		slog.Info("Can't connect to Postgres [ERROR]")
		return nil, err
	}

	slog.Info("Succesfully connected [OK]")
	// think about it
	db.conn = conn
	return conn, nil
}

func SaveUser(username, password string) error {
	conn, err := ConnectToDB()
	if err != nil {
		slog.Warn("Can't connect to Postgres [ERROR]")
		return err
	}
	defer conn.Close(context.Background())

	slog.Info("Connected to DB [OK]")

	query := `INSERT INTO users (username, user_password, is_admin) VALUES ($1, $2, $3)`
	_, err = conn.Exec(context.Background(), query, username, password, false)
	if err != nil {
		slog.Warn("Could not insert user [ERROR]", err)
		return err
	}

	slog.Info("User successfully saved! [OK]")
	return nil
}

func CheckIsAdmin(username string) (bool, error) {
	var isAdmin bool
	conn, err := ConnectToDB()
	if err != nil {
		slog.Warn("Can't connect to Postgres [ERROR]")
		return false, err
	}
	defer conn.Close(context.Background())

	slog.Info("Connected to DB [OK]")

	query := "SELECT is_admin FROM USERS WHERE username = $1"
	err = conn.QueryRow(context.Background(), query, username).Scan(&isAdmin)
	fmt.Println(isAdmin)

	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Info("User %s does not exist", username)
			return false, err
		}
		return false, fmt.Errorf("could not check if user is admin: %v", err)
	}
	return isAdmin, nil
}

func UpdateToAdmin(username string) error {
	conn, err := ConnectToDB()
	if err != nil {
		return fmt.Errorf("can't connect to Postgres: %v", err)
	}
	defer conn.Close(context.Background())

	query := "UPDATE users SET is_admin = true WHERE username = $1"
	_, err = conn.Exec(context.Background(), query, username)
	if err != nil {
		return fmt.Errorf("can't update user permission to ADMIN: %v", err)
	}

	return nil
}

// gorm method

// func SaveUser(db *gorm.DB, username, password string) error {
// 	user := User{
// 		Username: username,
// 		Password: password,
// 	}

// 	result := db.Create(&user)
// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	log.Println("User successfully saved!")
// 	return nil
// }
