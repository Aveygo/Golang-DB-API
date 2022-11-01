package main

import (
	"log"
	"net/http"

	"remote_db/endpoints"
	"remote_db/global"

	"go.etcd.io/bbolt"
)

func main() {

	// Initialize the database
	var err error
	global.DB, err = bbolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer global.DB.Close()

	// Create the main bucket
	global.DB.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("myBucket"))
		if err != nil {
			return err
		}
		return nil
	})

	// Create the endpoints
	http.HandleFunc("/set", endpoints.Set)
	http.HandleFunc("/get", endpoints.Get)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
