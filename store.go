package main

import "database/sql"

// Store - our store will have two methods, to add a new bird, and to get all existing birds
// Each method returns and error, in case something goes wrong
type Store interface {
	CreateBird(bird *Bird) error
	GetBirds() ([]*Bird, error)
}

type dbStore struct {
	db *sql.DB
}

// CreateBird -
func (store *dbStore) CreateBird(bird *Bird) error {
	_, err := store.db.Query("INSERT INTO birds(species, description) VALUES ($1,$2)", bird.Species, bird.Description)
	return err
}

// GetBirds -
func (store *dbStore) GetBirds() ([]*Bird, error) {
	rows, err := store.db.Query("SELECT species, description from birds;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	birds := []*Bird{}
	for rows.Next() {
		bird := &Bird{}
		if err := rows.Scan(&bird.Species, &bird.Description); err != nil {
			return nil, err
		}
		birds = append(birds, bird)
	}
	return birds, nil
}

var store Store

// InitStore -
func InitStore(s Store) {
	store = s
}
