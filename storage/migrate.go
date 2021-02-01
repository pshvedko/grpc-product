package storage

import (
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/globalsign/mgo"
)

const defaultMigrationCollection = "migrations"

type Migration struct {
	Id      int `bson:"_id"`
	Notes   string
	Applied time.Time
	up      MigrationFunc
	down    MigrationFunc
}

func (m Migration) Up(db *mgo.Database) (err error) {
	if err = m.up(db); err != nil {
		return
	}
	return db.C(defaultMigrationCollection).Insert(m)
}

func (m Migration) Down(db *mgo.Database) (err error) {
	if err = m.down(db); err != nil {
		return
	}
	return db.C(defaultMigrationCollection).RemoveId(m.Id)
}

type MigrationFunc func(db *mgo.Database) error

var order int

var migrations []Migration

func Register(up, down MigrationFunc) {
	_, file, _, _ := runtime.Caller(1)
	name := strings.Split(strings.TrimSuffix(filepath.Base(file), ".go"), "_")
	version, err := strconv.Atoi(name[0])
	if err != nil {
		panic(err)
	}
	if version != order {
		panic("invalid migration order")
	}
	migrations = append(migrations, Migration{
		Id:      version,
		Notes:   strings.Join(name[1:], " "),
		Applied: time.Now(),
		up:      up,
		down:    down,
	})
	order++
}

func migrate(db *mgo.Database, id uint32) (patch int, err error) {
	if id > 0 {
		return
	}
	return up(db, order)
}

func up(db *mgo.Database, to int) (patch int, err error) {
	patch, err = db.C(defaultMigrationCollection).Count()
	if err != nil {
		return
	}
	for _, m := range migrations[patch:to] {
		if err = m.Up(db); err != nil {
			return
		}
	}
	return
}
