package importer

import (
	"database/sql"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type SelectedPublication struct {
	Title           string
	Authors         []string
	PublicationDate time.Time
	Description     string
	URL             string
}

func (s SelectedPublication) Insert(db *sql.DB) error {
	stmt := `
	INSERT INTO selected_publications
	(title, authors, publication_date, description, url)
	VALUES (?, ?, ?, ?, ?)
	`

	_, err := db.Exec(stmt,
		s.Title,
		strings.Join(s.Authors, ", "),
		s.PublicationDate.Format("2006-01-02"),
		s.Description,
		s.URL,
	)

	if err != nil {
		log.Println("Encountered and error when trying to insert into the DB")
		return err
	}

	return nil

}

func InitializeTable() (*sql.DB, error) {
	os.Remove("./app.db")
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Println("Encountered an error when creating the DB")
		return nil, err
	}

	_, err = db.Exec("DROP TABLE IF EXISTS selected_publications")
	if err != nil {
		log.Println("Encountered an error when trying to drop selected publications table")
		return nil, err
	}

	createTableStmt := `
	CREATE TABLE selected_publications (
		id INTEGER NOT NULL PRIMARY KEY,
		title TEXT,
		authors TEXT,
		publication_date TEXT,
		description TEXT,
		url TEXT
	)
	`
	_, err = db.Exec(createTableStmt)
	if err != nil {
		log.Println("Encountered an error when creating the DB Table")
		return nil, err
	}
	return db, nil
}
