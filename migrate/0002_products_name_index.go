package migrate

import (
	"github.com/globalsign/mgo"
	"github.com/pshvedko/grpc-product/storage"
)

func init() {
	storage.Register(func(db *mgo.Database) error {
		return db.C("products").EnsureIndex(mgo.Index{
			Key:    []string{"name"},
			Unique: true,
		})
	}, func(db *mgo.Database) error {
		return db.C("products").DropIndexName("name_1")
	})
}
