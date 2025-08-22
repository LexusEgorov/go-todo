package db

import (
	"context"
	"fmt"

	"github.com/LexusEgorov/todo/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	prefix    = "Storage.db"
	opNew     = prefix + "New"
	opConnect = prefix + "Connect"
)

type DB struct {
	DB *pgxpool.Pool
}

func NewDB(cfg config.DBConfig) (*DB, error) {
	db := &DB{}

	connStr := config.GetConnStr(cfg.User, cfg.Password, cfg.Name)
	err := db.connect(connStr)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", opNew, err)
	}

	return db, nil
}

func (d *DB) Close() {
	d.DB.Close()
}

func (d *DB) connect(connStr string) error {
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return fmt.Errorf("%s: %v", opConnect, err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		return fmt.Errorf("%s: %v", opConnect, err)
	}

	d.DB = pool
	return nil
}
