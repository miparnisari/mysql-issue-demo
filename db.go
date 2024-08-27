package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/cenkalti/backoff/v4"
	"github.com/oklog/ulid/v2"
	"github.com/pressly/goose/v3"
)

type Config struct {
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxIdleTime time.Duration
	ConnMaxLifetime time.Duration

	ExportMetrics bool
}

type MySQL struct {
	db      *sql.DB
	stbl    sq.StatementBuilderType
	sqlTime interface{}
}

// New creates a new [MySQL] storage.
func New(uri string, cfg *Config) (*MySQL, error) {
	db, err := sql.Open("mysql", uri)
	if err != nil {
		return nil, fmt.Errorf("initialize mysql connection: %w", err)
	}
	if cfg.MaxOpenConns != 0 {
		db.SetMaxOpenConns(cfg.MaxOpenConns)
	}

	if cfg.MaxIdleConns != 0 {
		db.SetMaxIdleConns(cfg.MaxIdleConns)
	}

	if cfg.ConnMaxIdleTime != 0 {
		db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)
	}

	if cfg.ConnMaxLifetime != 0 {
		db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	}

	policy := backoff.NewExponentialBackOff()
	policy.MaxElapsedTime = 1 * time.Minute
	attempt := 1
	err = backoff.Retry(func() error {
		err := db.PingContext(context.Background())
		if err != nil {
			fmt.Println("waiting for mysql, attempt", attempt)
			attempt++
			return err
		}
		return nil
	}, policy)
	if err != nil {
		return nil, fmt.Errorf("ping db: %w", err)
	}

	stbl := sq.StatementBuilder.RunWith(db)

	return &MySQL{
		stbl:    stbl,
		db:      db,
		sqlTime: sq.Expr("NOW()"),
	}, nil
}

func (m *MySQL) Close() {
	m.db.Close()
}

func (m *MySQL) ReadAll(ctx context.Context) ([]*Store, error) {
	sb := m.stbl.
		Select(
			"id", "name").
		From("store")

	rows, err := sb.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	var stores []*Store
	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}

		stores = append(stores, &Store{
			Id:   id,
			Name: name,
		})
	}

	return stores, nil
}

func (m *MySQL) Write(ctx context.Context, storeID string) (*Store, error) {
	txn, err := m.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = txn.Rollback()
	}()

	_, err = m.stbl.
		Insert("store").
		Columns("id", "name").
		Values(storeID, ulid.Make().String()).
		RunWith(txn).
		ExecContext(ctx)
	if err != nil {
		return nil, err
	}

	var id, name string
	err = m.stbl.
		Select("id", "name").
		From("store").
		Where(sq.Eq{"id": storeID}).
		RunWith(txn).
		QueryRowContext(ctx).
		Scan(&id, &name)
	if err != nil {
		return nil, err
	}

	err = txn.Commit()
	if err != nil {
		return nil, err
	}

	return &Store{
		Id:   id,
		Name: name,
	}, nil
}

func (m *MySQL) IsReady(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	if err := m.db.PingContext(ctx); err != nil {
		return err
	}

	_, err := goose.GetDBVersion(m.db)
	if err != nil {
		return err
	}
	return nil
}
