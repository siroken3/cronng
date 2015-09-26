package cronng

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	uuid "github.com/satori/go.uuid"
)

func generate_uuid(scope *gorm.Scope) {
	if !scope.HasError() {
		scope.SetColumn("ID", uuid.NewV1().String())
	}
}

func NewLocalStrage() gorm.DB {
	gorm.DefaultCallback.Create().Before("gorm:save_before_associations").Register("cronng:generate_uuidmodel_id", generate_uuid)
	db, err := gorm.Open("sqlite3", "/tmp/cronbgdb.bin")
	if err != nil {
		panic(err)
	}
	//	db.LogMode(true)
	db.AutoMigrate(&Vm{})
	db.AutoMigrate(&VmRss{})
	db.AutoMigrate(&VmSwap{})
	db.AutoMigrate(&Monitoring{})
	db.AutoMigrate(&Job{})
	db.AutoMigrate(&Execution{})
	return db
}

/*
func (self *Job) PreInsert(s gorp.SqlExecutor) error {
	self.Id =
	self.Created = time.Now()
	return nil
}

func (self *Execution) PreInsert(s gorp.SqlExecutor) error {
	self.Id = uuid.NewV1().String()
	self.Started = time.Now()
	return nil
}
*/
