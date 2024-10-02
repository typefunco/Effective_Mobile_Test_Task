CREATE TABLE users(
	user_id SERIAL PRIMARY KEY,
	username VARCHAR(16) UNIQUE,
	user_password VARCHAR,
	is_admin BOOLEAN
)

CREATE TABLE songs(
	song_name VARCHAR(),
	song_text VARCHAR(),
	release_date VARCHAR(),
	song_link VARCHAR(),
	song_author VARCHAR()
)