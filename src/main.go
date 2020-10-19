package main

import (
	"gitlab.com/ludw1gj/mathfever-api/src/router"
	"log"
	"net/http"
)

func main() {
	log.Println("mathfever listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router.Load()))
}
