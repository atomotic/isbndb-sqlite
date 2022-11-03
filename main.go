package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS isbndb (
	isbn text generated always as (json_extract(data, '$.isbn13')), 
	title text generated always as (json_extract(data, '$.title')), 
	data text
);
CREATE INDEX IF NOT EXISTS title_idx ON isbndb (title);
CREATE INDEX IF NOT EXISTS isbn_idx ON isbndb (isbn);
`

func main() {
	db, err := sql.Open("sqlite3", "isbndb.db?cache=shared&mode=rwc&_journal_mode=WAL")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if _, err := db.Exec(schema); err != nil {
		log.Fatal(err)
	}

	raw, err := os.Open("isbndb_2022_09.jsonl")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(raw)
	const maxCapacity = 80 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	counter := 1
	for scanner.Scan() {
		if _, err := db.Exec("INSERT INTO isbndb(data) VALUES(?);", scanner.Bytes()); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("\rBooks imported: %d", counter)
		counter = counter + 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		fmt.Println(scanner.Text())
	}

}
