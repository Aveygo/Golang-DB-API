package endpoints

import (
	"log"
	"net/http"

	"remote_db/global"

	"go.etcd.io/bbolt"
)

func Set(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")

	err := global.DB.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("myBucket"))
		err := b.Put([]byte(key), []byte(value))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
}
