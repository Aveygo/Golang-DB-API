package endpoints

import (
	"log"
	"net/http"

	"remote_db/global"

	"go.etcd.io/bbolt"
)

func Get(w http.ResponseWriter, r *http.Request) {
	// Get the key from the url
	key := r.URL.Query().Get("key")

	err := global.DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("myBucket"))
		value := b.Get([]byte(key))
		w.Write(value)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
