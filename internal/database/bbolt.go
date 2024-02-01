package database

import (
	"sync"
	"time"

	"github.com/zaigie/palworld-server-tool/internal/logger"
	"go.etcd.io/bbolt"
)

var db *bbolt.DB
var dbMutex sync.Mutex

func InitDB() *bbolt.DB {
	db_, err := bbolt.Open("pst.db", 0600, &bbolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		logger.Panic(err)
	}
	// players
	err = db_.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("players"))
		return err
	})
	if err != nil {
		logger.Panic(err)
	}
	// guilds
	err = db_.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("guilds"))
		return err
	})
	if err != nil {
		logger.Panic(err)
	}
	return db_
}

func GetDB() *bbolt.DB {
	dbMutex.Lock()
	defer dbMutex.Unlock()
	if db == nil {
		db = InitDB()
	}
	return db
}
