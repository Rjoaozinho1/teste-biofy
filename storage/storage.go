package storage

import (
	"database/sql"

	"github.com/Rjoaozinho1/teste-biofy/repo"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateItem(item repo.Item) error
	GetItem(id string) (repo.Item, error)
	GetAllItens() ([]repo.Item, error)
	UpdateItem(id string, item repo.Item) error
	DeleteItem(id string) error
}

type PostgresStore struct {
	db *sql.DB
}

func NewStorageConnection() (*PostgresStore, error) {
	connStr := "user=joao-pedro dbname=db password=123 sslmode=disable"
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
	return s.CriaTabelaitens()
}

func (s *PostgresStore) CriaTabelaitens() error {
	query := `CREATE TABLE if not exists itens (
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

func (s *PostgresStore) CreateItem(item repo.Item) error {
	query := `INSERT INTO itens (nome, session_id, mensagem, created_at) VALUES ($1, $2, $3, $4);`
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

func (s *PostgresStore) GetItem(id string) (repo.Item, error) {
	query := `SELECT nome, session_id, mensagem, created_at FROM itens WHERE session_id = $1;`
	var item repo.Item
	err := s.db.QueryRow(query, id).Scan(&item.Nome, &item.SessionID, &item.Mensagem, &item.CreatedAt)
	if err != nil {
		return repo.Item{}, err
	}
	return item, nil
}

func (s *PostgresStore) GetAllItens() ([]repo.Item, error) {
	query := `SELECT nome, session_id, mensagem, created_at FROM itens;`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	itens := []repo.Item{}
	for rows.Next() {
		var item repo.Item
		err := rows.Scan(
			&item.Nome,
			&item.SessionID,
			&item.Mensagem,
			&item.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		itens = append(itens, item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return itens, nil
}

func (s *PostgresStore) UpdateItem(id string, item repo.Item) error {
	query := `UPDATE itens SET nome = $1, mensagem = $2 WHERE session_id = $3;`
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
	query := `DELETE FROM itens WHERE session_id = $1;`
	_, err := s.db.Exec(
		query,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}
