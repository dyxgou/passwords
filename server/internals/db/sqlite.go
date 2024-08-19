package internals

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}

	createUsersQuery := `CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT UNIQUE NOT NULL,
    passwords INTEGER DEFAULT 0,
    points INTEGER DEFAULT 0
  );`

	createPasswordsQuery := `CREATE TABLE IF NOT EXISTS passwords (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    content TEXT NOT NULL,
    rot TEXT NOT NULL,
    sha TEXT NOT NULL,
    security_level TEXT CHECK( security_level IN ("horrible", "bad", "mid", "good", "excellent") ),
    user_id INTEGER,
    FOREIGN KEY(user_id) REFERENCES users(id)
  );`

	_, err = db.Exec(createUsersQuery)
	if err != nil {
		log.Fatalf("Failed to create the table %v\n", err)
	}

	_, err = db.Exec(createPasswordsQuery)
	if err != nil {
		log.Fatalf("Failed to create the table %v\n", err)
	}

	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		panic(err)
	}

	return db
}
