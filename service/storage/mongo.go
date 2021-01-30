package storage

import (
	"github.com/globalsign/mgo"
)

func z() {
	iter := mgo.Iter{}
	iter.Done()
}
