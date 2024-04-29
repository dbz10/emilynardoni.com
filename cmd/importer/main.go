package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dbz10/emilynardoni.com/pkg/importer"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := importer.InitializeTable()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open("./static/content/selected_publications.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		t, err := time.Parse("2006-01-02", record[2])
		if err != nil {
			log.Fatal(err)
		}

		authors := strings.Split(record[1], ", ")

		err = importer.SelectedPublication{
			Title:           record[0],
			Authors:         authors,
			PublicationDate: t,
			Description:     record[3],
			URL:             record[4],
		}.Insert(db)

		if err != nil {
			log.Fatal(err)
		}
	}

}
