package migrator

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/repository/mysql"
	migrate "github.com/rubenv/sql-migrate"
	"log"
)

const migrationAddressDirectory = "../internal/repository/mysql/migrations"

type Migrator struct {
	dialect    string
	dbConfig   mysql.Config
	migrations *migrate.FileMigrationSource
}

func New(dbConfig mysql.Config) Migrator {
	// OR: Read migrations from a folder:
	migrations := &migrate.FileMigrationSource{
		Dir: migrationAddressDirectory,
	}

	return Migrator{dbConfig: dbConfig, dialect: "mysql", migrations: migrations}
}

func (m Migrator) Up() {
	db, err := sql.Open(m.dialect, fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true",
		m.dbConfig.Username, m.dbConfig.Password, m.dbConfig.Host, m.dbConfig.Port, m.dbConfig.DBName))
	if err != nil {
		log.Fatalf("can't open mysql db: %v", err)
	}

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Up)
	if err != nil {
		log.Fatalf("can't apply migrations: %v", err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}

func (m Migrator) Down() {
	db, err := sql.Open(m.dialect, fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true",
		m.dbConfig.Username, m.dbConfig.Password, m.dbConfig.Host, m.dbConfig.Port, m.dbConfig.DBName))
	if err != nil {
		panic(fmt.Errorf("can't open mysql db: %v", err))
	}

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Down)
	if err != nil {
		panic(fmt.Errorf("can't rollback migrations: %v", err))
	}
	fmt.Printf("Rollbacked %d migrations!\n", n)
}
