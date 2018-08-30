package main

import (
"net/http"
"log"
 "fmt"
 "io/ioutil" )
func main() {
	 res, err := http.Get("127.0.0.1/cobregram")
	if err != nil {
	    log.Fatal(err)
	}
	contenido, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Printf("%s", contenido)
	 
}