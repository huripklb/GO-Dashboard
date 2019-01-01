package main

import (
	"auth"
	"fmt"
	"log"
	"net/http"

	tf "testfolder"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	//r.HandleFunc("/", HomePage)
	//r.HandleFunc("/another", AnotherPage)
	r.HandleFunc("/testauth", testAuth).Methods("POST")
	r.HandleFunc("/testfolder", tf.FolderPage).Methods("POST")
	r.HandleFunc("/testpost", tf.PostPage).Methods("POST")
	r.HandleFunc("/testget", tf.GetPage).Methods("GET")
	r.HandleFunc("/testput", tf.TestPutBeans).Methods("POST")    // PutBeans
	r.HandleFunc("/testput", tf.TestReserveBeans).Methods("GET") // ReserveBeans
	r.Use(auth.BasicAuth)

	log.Fatal(http.ListenAndServe(":8001", r))
}

func testAuth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Success")
}
