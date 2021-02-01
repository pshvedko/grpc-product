package migrate

import (
	"github.com/globalsign/mgo"
	"github.com/pshvedko/grpc-product/storage"
)

func init() {
	storage.Register(func(db *mgo.Database) error {
		return db.C("products").Create(&mgo.CollectionInfo{})
	}, func(db *mgo.Database) error {
		return db.C("products").DropCollection()
	})
}
