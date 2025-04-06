package database

import (
    "fmt"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
    "github.com/vanthang24803/mini/internal/config"
)

var DB *sqlx.DB

func InitPostgres(cfg *config.Config) error {
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        cfg.Database.Host,
        cfg.Database.Port,
        cfg.Database.Username,
        cfg.Database.Password,
        cfg.Database.Name,
    )

    db, err := sqlx.Connect("postgres", dsn)
    if err != nil {
        return fmt.Errorf("error connecting to database: %v", err)
    }

    if err := db.Ping(); err != nil {
        return fmt.Errorf("error pinging database: %v", err)
    }

    DB = db
    return nil
}

func GetDB() *sqlx.DB {
    return DB
}