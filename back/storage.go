package main

import (
	"database/sql"
)

type Storage interface {
	CreateItem(item Item) error
	GetItem(id string) (Item, error)
	GetAllItems() ([]Item, error)
	UpdateItem(id string, item Item) error
	DeleteItem(id string) error
}

type PostgresStore struct {
	db *sql.DB
}

func newStorageConnection() (*PostgresStore, error) {
	connStr := "user=user dbname=dbname password=password sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.CriaTabelaItems()
}

func (s *PostgresStore) CriaTabelaItems() error {
	query := `CREATE TABLE if not exists items (
		session_id UUID PRIMARY KEY,
		nome VARCHAR(100) NOT NULL,
		mensagem TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) CreateItem(item Item) error {
	query := `INSERT INTO items (nome, session_id, mensagem, created_at) VALUES ($1, $2, $3, $4);`
	_, err := s.db.Exec(
		query,
		item.Nome,
		item.SessionID,
		item.Mensagem,
		item.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) GetItem(id string) (Item, error) {
	query := `SELECT nome, session_id, mensagem, created_at FROM items WHERE session_id = $1;`
	var item Item
	err := s.db.QueryRow(query, id).Scan(&item.Nome, &item.SessionID, &item.Mensagem, &item.CreatedAt)
	if err != nil {
		return Item{}, err
	}
	return item, nil
}

func (s *PostgresStore) GetAllItems() ([]Item, error) {
	query := `SELECT nome, session_id, mensagem, created_at FROM items;`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []Item{}
	for rows.Next() {
		var item Item
		err := rows.Scan(
			&item.Nome,
			&item.SessionID,
			&item.Mensagem,
			&item.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *PostgresStore) UpdateItem(id string, item Item) error {
	query := `UPDATE items SET nome = $1, mensagem = $2 WHERE session_id = $3;`
	_, err := s.db.Exec(
		query,
		item.Nome,
		item.Mensagem,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) DeleteItem(id string) error {
	query := `DELETE FROM items WHERE session_id = $1;`
	_, err := s.db.Exec(
		query,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}
