package cronng

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
	uuid "github.com/satori/go.uuid"
	gorp "gopkg.in/gorp.v1"
)

func NewLocalStrage() *gorp.DbMap {
	db, err := sql.Open("sqlite3", "/tmp/cronbgdb.bin")
	if err != nil {
		panic(err)
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.AddTableWithName(Job{}, "job").SetKeys(false, "Id")
	dbmap.AddTableWithName(Execution{}, "execution").SetKeys(false, "Id")
	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		panic(err)
	}
	return dbmap
}

func (self *Job) PreInsert(s gorp.SqlExecutor) error {
	self.Id = uuid.NewV1().String()
	self.Created = time.Now()
	return nil
}

func (self *Execution) PreInsert(s gorp.SqlExecutor) error {
	self.Id = uuid.NewV1().String()
	self.Started = time.Now()
	return nil
}
