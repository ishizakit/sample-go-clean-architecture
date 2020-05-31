package infra

import (
	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"

	// mysql driver
	_ "github.com/mattn/go-sqlite3"
)

// NewInMemorySQLiteDB DBクライアントの取得
func NewInMemorySQLiteDB() (*sqlx.DB, error) {
	// DB生成
	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}
	// マイグレーション
	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}
	_, err = migrate.Exec(db.DB, "sqlite3", migrations, migrate.Up)
	if err != nil {
		return nil, err
	}

	return db, nil
}
