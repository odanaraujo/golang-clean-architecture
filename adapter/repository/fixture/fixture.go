package fixture

import (
	"context"
	"database/sql"
	"io/fs"
	"log"

	"github.com/maragudk/migrate"
)

func Up(migrationsDir fs.FS) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err := migrate.Up(context.Background(), db, migrationsDir); err != nil {
		panic(err)
	}

	return db
}

func Down(db *sql.DB, migrationsDir fs.FS) {
	if err := migrate.Down(context.Background(), db, migrationsDir); err != nil {
		panic(err)
	}
	defer db.Close()
}
