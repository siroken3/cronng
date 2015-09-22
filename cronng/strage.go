package cronng

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	gorp "gopkg.in/gorp.v1"
)

func NewLocalStrage() (dbmap *gorp.DbMap) {
	db, err := sql.Open("sqlite3", "/tmp/cronbgdb.bin")
	if _err != nil {
		panic(_err)
	}
	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.AddTableWithName(Job{}, "job").SetKeys(true, "Id")
	_err = dbmap.CreateTalesIfNotExists()
	if _err != nil {
		panic(_err)
	}
	return
}
