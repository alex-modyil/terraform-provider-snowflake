package snowflake

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
)

func Exec(db *sql.DB, query string) error {
	log.Print("[DEBUG] stmt ", query)

	_, err := db.Exec(query)
	return err
}

func QueryRow(db *sql.DB, stmt string) *sqlx.Row {
	log.Print("[DEBUG] stmt ", stmt)
	sdb := sqlx.NewDb(db, "snowflake").Unsafe()
	return sdb.QueryRowx(stmt)
}

// TODO https://godoc.org/github.com/jmoiron/sqlx#Stmt.Unsafe
func Query(db *sql.DB, stmt string) (*sqlx.Rows, error) {
	sdb := sqlx.NewDb(db, "snowflake").Unsafe()
	return sdb.Queryx(stmt)
}