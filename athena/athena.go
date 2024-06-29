package athena

import (
	"fmt"

	"go.etcd.io/bbolt"
)

const (
	defaultDBName = "default"
)
type  Collection struct {
	bucket *bbolt.Bucket
}
type Athena struct {
	db *bbolt.DB
}

func New()(*Athena, error) {
	dbname := fmt.Sprintf("%s.ath", defaultDBName)
	db, err := bbolt.Open(dbname, 0666, nil)
	if err != nil {
		return nil, err
	}
	return &Athena{
		db : db,
	}, nil
}

func (a *Athena) CreateCollection(name string)(*Collection, error){
	coll := Collection{}
	err := a.db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucket([]byte(name))
		if err != nil {
			return err
		}
		coll.bucket = bucket
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &coll, nil
}

func (a *Athena) Put (coll *Collection, k, v string){
	//47:28
}
// db.Update(func(tx *bbolt.Tx) error {
		
// 	id := uuid.New()
// 	for k, v := range user {
// 		if err := bucket.Put([]byte(k), []byte(v)); err != nil {
// 			return err
// 		}
// 	}
// 	if err := bucket.Put([]byte("id"), []byte(id.String())); err != nil {
// 		return err
// 	}
// 	return nil
// })