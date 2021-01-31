package migrate

import (
	"github.com/globalsign/mgo"
	"github.com/pshvedko/grpc-product/storage"
)

func init() {
	storage.Register(func(db *mgo.Database) error {
		return nil
	}, func(db *mgo.Database) error {
		return nil
	})
}
