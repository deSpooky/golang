package main

import (
	"database/sql"
	"time"
)

func InitDB(conn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS cards (
		id int primary key,
		url text,
		duration text
	)`)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InsertCard(db *sql.DB, card PostCard) error {
	_, err := db.Exec(`INSERT INTO cards (id, url, duration) VALUES ($1, $2, $3)
	ON CONFLICT (id) DO UPDATE SET url = EXCLUDED.url, duration = EXCLUDED.duration`,
		card.ID, card.Url, card.Duration.String())
	return err
}

func LoadCards(db *sql.DB) ([]PostCard, error) {
	rows, err := db.Query(`SELECT id, url, duration FROM cards`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cards []PostCard
	for rows.Next() {
		var c PostCard
		var dur string
		if err := rows.Scan(&c.ID, &c.Url, &dur); err != nil {
			return nil, err
		}
		c.Duration, _ = time.ParseDuration(dur)
		cards = append(cards, c)
	}
	return cards, nil
}