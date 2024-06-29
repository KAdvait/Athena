package main

import (
	"fmt"
	"log"
	"github.com/KAdvait/Athena/athena"
	
)

func main() {
	user := map[string]string{
		"name": "advait",
		"age":  "21",
	}
	_ = user //all datatypes
	db, err := athena.New()
	if err != nil {
		log.Fatal(err)
	}
	coll, err := db.CreateCollection("users")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", coll)
}
