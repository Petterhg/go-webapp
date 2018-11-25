package db

import (
	"go-webapp/model"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	ConnectString string
}

type pgDb struct {
	dbConn *sqlx.DB

	sqlSelectUser *sqlx.Stmt
}

func InitDb(cfg Config) (*pgDb, error) {
	if dbConn, err := sqlx.Connect("postgres", cfg.ConnectString); err != nil {
		return nil, err
	} else {
		p := &pgDb{dbConn: dbConn}
		if err := p.dbConn.Ping(); err != nil {
			log.Printf("Could not ping database: %v\n", err)
			return nil, err
		}
		if err := p.createTablesIfNotExist(); err != nil {
			log.Printf("Could not create table: %v\n", err)
			return nil, err
		}
		if err := p.prepareSqlStatements(); err != nil {
			log.Printf("Could not run prepared sql statement: %v\n", err)
			return nil, err
		}
		return p, nil
	}
}

func (p *pgDb) createTablesIfNotExist() error {
	create_sql := "CREATE TABLE IF NOT EXISTS users (id SERIAL NOT NULL PRIMARY KEY, firstname TEXT NOT NULL, lastname TEXT NOT NULL)"
	if rows, err := p.dbConn.Query(create_sql); err != nil {
		return err
	} else {
		rows.Close()
	}
	return nil
}

func (p *pgDb) prepareSqlStatements() (err error) {

	if p.sqlSelectUser, err = p.dbConn.Preparex(
		"SELECT id, firstname, lastname FROM users",
	); err != nil {
		return err
	}

	return nil
}

func (p *pgDb) SelectUser() ([]*model.User, error) {
	user := make([]*model.User, 0)
	if err := p.sqlSelectUser.Select(&user); err != nil {
		return nil, err
	}
	return user, nil
}
