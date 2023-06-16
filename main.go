package main

import (
	"fmt"
	"keyspeed/util"
	"log"
	"net/http"
)

const PORT = ":8071"

func main() {

	fmt.Println("Hello World")

	setup()
	util.BLogln("serving on ", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))

}
